// Unless explicitly stated otherwise all files in this repository are licensed
// under the Apache License Version 2.0.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2016-present Datadog, Inc.

//go:build windows && npm
// +build windows,npm

package tracer

import (
	"context"
	"fmt"
	"runtime"
	"sync"
	"time"

	"golang.org/x/sys/windows"
	"golang.org/x/sys/windows/registry"
	"golang.org/x/sys/windows/svc"
	"golang.org/x/sys/windows/svc/mgr"

	"github.com/DataDog/datadog-agent/pkg/ebpf"
	"github.com/DataDog/datadog-agent/pkg/network"
	"github.com/DataDog/datadog-agent/pkg/network/config"
	"github.com/DataDog/datadog-agent/pkg/network/dns"
	"github.com/DataDog/datadog-agent/pkg/network/driver"
	"github.com/DataDog/datadog-agent/pkg/network/protocols/http"
	"github.com/DataDog/datadog-agent/pkg/process/util"
	"github.com/DataDog/datadog-agent/pkg/util/log"
	"github.com/DataDog/datadog-agent/pkg/util/winutil"
)

const (
	defaultPollInterval = int(15)
	defaultBufferSize   = 512
	minBufferSize       = 256
)

// Tracer struct for tracking network state and connections
type Tracer struct {
	config          *config.Config
	driverInterface *network.DriverInterface
	stopChan        chan struct{}
	state           network.State
	reverseDNS      dns.ReverseDNS
	httpMonitor     http.Monitor

	activeBuffer *network.ConnectionBuffer
	closedBuffer *network.ConnectionBuffer
	connLock     sync.Mutex

	timerInterval int

	// Connections for the tracer to exclude
	sourceExcludes []*network.ConnectionFilter
	destExcludes   []*network.ConnectionFilter

	// polling loop for connection event
	closedEventLoop sync.WaitGroup
}

func isClosedSourceAllowed() bool {
	ddRegKey := `SOFTWARE\DataDog\Datadog Agent`
	regKey, err := registry.OpenKey(registry.LOCAL_MACHINE, `SOFTWARE\DataDog\Datadog Agent`, registry.QUERY_VALUE)
	if err != nil {
		log.Warnf("unable to open registry key %s: %v", ddRegKey, err)
		return false
	}
	defer regKey.Close()

	ddRegName := "AllowClosedSource"
	val, _, err := regKey.GetIntegerValue("AllowClosedSource")
	if err != nil {
		log.Warnf("unable to get value for %s: %v", ddRegName, err)
		return false
	}

	if val == 1 {
		log.Infof("%s is set to 1; closed-source software allowed", ddRegName)
		return true
	} else if val == 0 {
		log.Infof("%s is set to 0; closed-source software not allowed", ddRegName)
		return false
	} else {
		log.Infof("Unexpected value set for %s: %d", ddRegName, ddRegName)
		return false
	}
}

func ddnpmIsEnabled() (enabled bool, err error) {
	enabled = false

	// connect to SCM
	manager, err := winutil.OpenSCManager(windows.SC_MANAGER_CONNECT)
	if err != nil {
		return
	}
	defer manager.Disconnect()

	// get a handle to the ddnpm service
	ddnpmServiceAccess := windows.SERVICE_QUERY_CONFIG
	ddnpmServiceHandle, err := winutil.OpenService(manager, "ddnpm", uint32(ddnpmServiceAccess))
	if err != nil {
		return
	}
	defer ddnpmServiceHandle.Close()

	ddnpmConfig, err := ddnpmServiceHandle.Config()
	if err != nil {
		return
	}

	return (ddnpmConfig.StartType != windows.SERVICE_DISABLED), nil

}

func ddnpmIsRunning() (running bool, err error) {
	running = false

	// connect to SCM
	manager, err := winutil.OpenSCManager(windows.SC_MANAGER_CONNECT)
	if err != nil {
		return
	}
	defer manager.Disconnect()

	// get a handle to the ddnpm service
	ddnpmServiceAccess := windows.SERVICE_CHANGE_CONFIG | windows.SERVICE_START | windows.SERVICE_QUERY_STATUS
	ddnpmServiceHandle, err := winutil.OpenService(manager, "ddnpm", uint32(ddnpmServiceAccess))
	if err != nil {
		return
	}
	defer ddnpmServiceHandle.Close()

	ddnpmStatus, err := ddnpmServiceHandle.Query()
	if err != nil {
		return
	}

	running = (ddnpmStatus.State == svc.Running)
	return
}

func enableDdnpmService() (enabled bool, err error) {
	enabled = false

	// connect to SCM
	manager, err := winutil.OpenSCManager(windows.SC_MANAGER_CONNECT)
	if err != nil {
		return
	}
	defer manager.Disconnect()

	// get a handle to the ddnpm service
	ddnpmServiceAccess := windows.SERVICE_CHANGE_CONFIG | windows.SERVICE_START | windows.SERVICE_QUERY_STATUS
	ddnpmServiceHandle, err := winutil.OpenService(manager, "ddnpm", uint32(ddnpmServiceAccess))
	if err != nil {
		return
	}
	defer ddnpmServiceHandle.Close()

	config := mgr.Config{
		StartType: mgr.StartManual,
	}

	// set it to manual start
	err = ddnpmServiceHandle.UpdateConfig(config)
	if err != nil {
		log.Warnf("could not enable %s: %v", err)
	}
	enabled = true
	return
}

func startDdnpm() (running bool, err error) {
	running = false

	// connect to SCM
	manager, err := winutil.OpenSCManager(windows.SC_MANAGER_CONNECT)
	if err != nil {
		return
	}
	defer manager.Disconnect()

	// get a handle to the ddnpm service
	ddnpmServiceAccess := windows.SERVICE_START | windows.SERVICE_QUERY_STATUS
	ddnpmServiceHandle, err := winutil.OpenService(manager, "ddnpm", uint32(ddnpmServiceAccess))
	if err != nil {
		return
	}
	defer ddnpmServiceHandle.Close()
	err = ddnpmServiceHandle.Start()
	return
}

// NewTracer returns an initialized tracer struct
func NewTracer(config *config.Config) (*Tracer, error) {

	var di *network.DriverInterface
	var err error

	// check driver service is enabled
	enabled, _ := ddnpmIsEnabled()
	if enabled {
		log.Info("Driver is enabled, checking if running")
		// check driver is running
		running, _ := ddnpmIsRunning()
		if running {
			log.Info("Driver is running, check if can connect")
			// try to connect to the NPM driver
			di, err = network.NewDriverInterface(config, driver.NewHandle)
			if err != nil {
				log.Warnf("Driver running but could not connect: %v", err)
				return nil, err
			}
			log.Info("Driver connection success")
		}
	} else {
		log.Info("Driver service not enabled, if closed source is allowed will enable")
		if isClosedSourceAllowed() {
			enabled, err := enableDdnpmService()
			if enabled == false {
				log.Warnf("Could not enable ddnpm: %v", err)
				return nil, err
			}
			running, err := startDdnpm()
			if running == false {
				log.Warnf("Could not run ddnpm: %v", err)
				return nil, err
			}
		} else {
			return nil, fmt.Errorf("The Windows driver is not enabled, ensure closed source is allowed")
		}
	}

	di, err = network.NewDriverInterface(config, driver.NewHandle)
	if err != nil {
		log.Warnf("Driver running but could not connect: %v", err)
		return nil, err
	}

	state := network.NewState(
		config.ClientStateExpiry,
		config.MaxClosedConnectionsBuffered,
		config.MaxConnectionsStateBuffered,
		config.MaxDNSStatsBuffered,
		config.MaxHTTPStatsBuffered,
	)

	reverseDNS := dns.NewNullReverseDNS()
	if config.DNSInspection {
		reverseDNS, err = dns.NewReverseDNS(config)
		if err != nil {
			return nil, err
		}
	}

	tr := &Tracer{
		config:          config,
		driverInterface: di,
		stopChan:        make(chan struct{}),
		timerInterval:   defaultPollInterval,
		state:           state,
		activeBuffer:    network.NewConnectionBuffer(defaultBufferSize, minBufferSize),
		closedBuffer:    network.NewConnectionBuffer(defaultBufferSize, minBufferSize),
		reverseDNS:      reverseDNS,
		httpMonitor:     newHttpMonitor(config, di.GetHandle()),
		sourceExcludes:  network.ParseConnectionFilters(config.ExcludedSourceConnections),
		destExcludes:    network.ParseConnectionFilters(config.ExcludedDestinationConnections),
	}
	tr.closedEventLoop.Add(1)
	go func() {
		defer tr.closedEventLoop.Done()

		runtime.LockOSThread()
		defer runtime.UnlockOSThread()
	waitloop:
		for {
			evt, _ := windows.WaitForSingleObject(di.GetClosedFlowsEvent(), windows.INFINITE)
			switch evt {
			case windows.WAIT_OBJECT_0:
				_, err = tr.driverInterface.GetClosedConnectionStats(tr.closedBuffer, func(c *network.ConnectionStats) bool {
					return !tr.shouldSkipConnection(c)
				})
				closedConnStats := tr.closedBuffer.Connections()

				tr.state.StoreClosedConnections(closedConnStats)

			case windows.WAIT_FAILED:
				break waitloop

			default:
				log.Infof("got other wait value %v", evt)
			}
		}

	}()
	return tr, nil
}

// Stop function stops running tracer
func (t *Tracer) Stop() {
	close(t.stopChan)
	if t.httpMonitor != nil { //nolint
		_ = t.httpMonitor.Stop()
	}
	t.reverseDNS.Close()
	err := t.driverInterface.Close()
	if err != nil {
		log.Errorf("error closing driver interface: %s", err)
	}
	t.closedEventLoop.Wait()
}

// GetActiveConnections returns all active connections
func (t *Tracer) GetActiveConnections(clientID string) (*network.Connections, error) {
	t.connLock.Lock()
	defer t.connLock.Unlock()

	_, err := t.driverInterface.GetOpenConnectionStats(t.activeBuffer, func(c *network.ConnectionStats) bool {
		return !t.shouldSkipConnection(c)
	})
	if err != nil {
		return nil, fmt.Errorf("error retrieving open connections from driver: %w", err)
	}
	_, err = t.driverInterface.GetClosedConnectionStats(t.closedBuffer, func(c *network.ConnectionStats) bool {
		return !t.shouldSkipConnection(c)
	})
	if err != nil {
		return nil, fmt.Errorf("error retrieving closed connections from driver: %w", err)
	}
	activeConnStats := t.activeBuffer.Connections()
	closedConnStats := t.closedBuffer.Connections()

	// check for expired clients in the state
	t.state.RemoveExpiredClients(time.Now())

	t.state.StoreClosedConnections(closedConnStats)

	var delta network.Delta
	if t.httpMonitor != nil { //nolint
		delta = t.state.GetDelta(clientID, uint64(time.Now().Nanosecond()), activeConnStats, t.reverseDNS.GetDNSStats(), t.httpMonitor.GetHTTPStats())
	} else {
		delta = t.state.GetDelta(clientID, uint64(time.Now().Nanosecond()), activeConnStats, t.reverseDNS.GetDNSStats(), nil)
	}

	t.activeBuffer.Reset()
	t.closedBuffer.Reset()

	ips := make([]util.Address, 0, len(delta.Conns)*2)
	for _, conn := range delta.Conns {
		ips = append(ips, conn.Source, conn.Dest)
	}
	names := t.reverseDNS.Resolve(ips)
	telemetryDelta := t.state.GetTelemetryDelta(clientID, t.getConnTelemetry())
	return &network.Connections{
		BufferedData:  delta.BufferedData,
		HTTP:          delta.HTTP,
		DNS:           names,
		DNSStats:      delta.DNSStats,
		ConnTelemetry: telemetryDelta,
	}, nil
}

// RegisterClient registers the client
func (t *Tracer) RegisterClient(clientID string) error {
	t.state.RegisterClient(clientID)
	return nil
}

func (t *Tracer) getConnTelemetry() map[network.ConnTelemetryType]int64 {
	tm := map[network.ConnTelemetryType]int64{}

	// allStats is the expvar map.  it is actually a map of maps
	// top level keys are:
	//   state (we don't need for this call)
	//   dns   ( the dns handle stats)
	//   each of the strings in DriverExpvarNames.  We're interested
	//   in driver.flowHandleStats, which is "driver_flow_handle_stats"
	if allstats, err := t.driverInterface.GetStats(); err == nil {
		if flowStats, ok := allstats["driver_flow_handle_stats"].(map[string]int64); ok {
			if fme, ok := flowStats["num_flows_missed_max_exceeded"]; ok {
				tm[network.NPMDriverFlowsMissedMaxExceeded] = fme
			}
		}
	}
	dnsStats := t.reverseDNS.GetStats()
	if pp, ok := dnsStats["packets_processed_transport"]; ok {
		tm[network.MonotonicDNSPacketsProcessed] = pp
	}
	if pd, ok := dnsStats["read_packets_skipped"]; ok {
		tm[network.MonotonicDNSPacketsDropped] = pd
	}

	return tm
}

// GetStats returns a map of statistics about the current tracer's internal state
func (t *Tracer) GetStats() (map[string]interface{}, error) {
	driverStats, err := t.driverInterface.GetStats()
	if err != nil {
		log.Errorf("not printing driver stats: %v", err)
	}

	stats := map[string]interface{}{
		"state": t.state.GetStats(),
		"dns":   t.reverseDNS.GetStats(),
	}
	for _, name := range network.DriverExpvarNames {
		stats[string(name)] = driverStats[name]
	}
	return stats, nil
}

// DebugNetworkState returns a map with the current tracer's internal state, for debugging
func (t *Tracer) DebugNetworkState(_ string) (map[string]interface{}, error) {
	return nil, ebpf.ErrNotImplemented
}

// DebugNetworkMaps returns all connections stored in the maps without modifications from network state
func (t *Tracer) DebugNetworkMaps() (*network.Connections, error) {
	return nil, ebpf.ErrNotImplemented
}

// DebugEBPFMaps is not implemented on this OS for Tracer
func (t *Tracer) DebugEBPFMaps(maps ...string) (string, error) {
	return "", ebpf.ErrNotImplemented
}

// DebugCachedConntrack is not implemented on this OS for Tracer
func (t *Tracer) DebugCachedConntrack(ctx context.Context) (interface{}, error) {
	return nil, ebpf.ErrNotImplemented
}

// DebugHostConntrack is not implemented on this OS for Tracer
func (t *Tracer) DebugHostConntrack(ctx context.Context) (interface{}, error) {
	return nil, ebpf.ErrNotImplemented
}

func newHttpMonitor(c *config.Config, dh driver.Handle) http.Monitor {
	if !c.EnableHTTPMonitoring && !c.EnableHTTPSMonitoring {
		return nil
	}
	log.Infof("http monitoring has been enabled")

	var monitor http.Monitor
	var err error

	monitor, err = http.NewWindowsMonitor(c, dh)

	if err != nil {
		log.Errorf("could not instantiate http monitor: %s", err)
		return nil
	}
	monitor.Start()
	return monitor
}
