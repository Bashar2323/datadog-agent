using System;
using System.Linq;
using System.Runtime.InteropServices;
using System.Security.Principal;
using System.ServiceProcess;
using Datadog.CustomActions.Extensions;
using Datadog.CustomActions.Interfaces;
using Datadog.CustomActions.Native;
using Microsoft.Deployment.WindowsInstaller;
using Microsoft.Win32;

namespace Datadog.CustomActions
{
    public class ServiceCustomAction
    {
        private readonly ISession _session;
        private readonly INativeMethods _nativeMethods;
        private readonly IRegistryServices _registryServices;
        private readonly IDirectoryServices _directoryServices;
        private readonly IFileServices _fileServices;
        private readonly IServiceController _serviceController;

        public ServiceCustomAction(
            ISession session,
            INativeMethods nativeMethods,
            IRegistryServices registryServices,
            IDirectoryServices directoryServices,
            IFileServices fileServices,
            IServiceController serviceController)
        {
            _session = session;
            _nativeMethods = nativeMethods;
            _registryServices = registryServices;
            _directoryServices = directoryServices;
            _fileServices = fileServices;
            _serviceController = serviceController;
        }

        public ServiceCustomAction(ISession session)
        : this(
            session,
            new Win32NativeMethods(),
            new RegistryServices(),
            new DirectoryServices(),
            new FileServices(),
            new Native.ServiceController()
        )
        {
        }

        private static ActionResult EnsureNpmServiceDepdendency(ISession session)
        {
            try
            {
                var addLocal = session.Property("ADDLOCAL").ToUpper();
                session.Log($"ADDLOCAL={addLocal}");
                using var systemProbeDef = Registry.LocalMachine.OpenSubKey(@"SYSTEM\CurrentControlSet\Services\datadog-system-probe", true);
                if (systemProbeDef != null)
                {
                    if (string.IsNullOrEmpty(addLocal))
                    {
                        systemProbeDef.SetValue("DependOnService", new[]
                        {
                            "datadogagent"
                        }, RegistryValueKind.MultiString);

                    }
                    else if (addLocal.Contains("NPM") ||
                             addLocal.Contains("ALL"))
                    {
                        systemProbeDef.SetValue("DependOnService", new[]
                        {
                            "datadogagent",
                            "ddnpm"
                        }, RegistryValueKind.MultiString);
                    }
                }
                else
                {
                    session.Log("Registry key does not exist");
                }
            }
            catch (Exception e)
            {
                session.Log($"Could not update system probe dependent service: {e.Message}");
                session.Log(e.ToString());
            }
            return ActionResult.Success;
        }

        [CustomAction]
        public static ActionResult EnsureNpmServiceDependency(Session session)
        {
            return EnsureNpmServiceDepdendency(new SessionWrapper(session));
        }

        private ActionResult ConfigureServiceUsers()
        {
            try
            {
                // Lookup account so we can determine how to set the password according to the ChangeServiceConfig rules.
                // https://learn.microsoft.com/en-us/windows/win32/api/winsvc/nf-winsvc-changeserviceconfigw
                var ddAgentUserName = $"{_session.Property("DDAGENTUSER_PROCESSED_FQ_NAME")}";
                var userFound = _nativeMethods.LookupAccountName(ddAgentUserName,
                    out _,
                    out _,
                    out var securityIdentifier,
                    out _);
                if (!userFound)
                {
                    throw new Exception($"Could not find user {ddAgentUserName}.");
                }

                var ddAgentUserPassword = _session.Property("DDAGENTUSER_PROCESSED_PASSWORD");
                var isServiceAccount = _nativeMethods.IsServiceAccount(securityIdentifier);
                if (!isServiceAccount && string.IsNullOrEmpty(ddAgentUserPassword))
                {
                    _session.Log("Password not provided, will not change service user password");
                    // set to null so we don't modify the service config
                    ddAgentUserPassword = null;
                }
                else if (isServiceAccount)
                {
                    _session.Log("Ignoring provided password because account is a service account");
                    // Follow rules for ChangeServiceConfig
                    if (securityIdentifier.IsWellKnown(WellKnownSidType.LocalSystemSid) ||
                        securityIdentifier.IsWellKnown(WellKnownSidType.LocalServiceSid) ||
                        securityIdentifier.IsWellKnown(WellKnownSidType.NetworkServiceSid))
                    {
                        // Specify an empty string if the account has no password or if the service runs in the LocalService, NetworkService, or LocalSystem account.
                        ddAgentUserPassword = "";
                    }
                    else
                    {
                        // If the account name specified by the lpServiceStartName parameter is the name of a managed service account or virtual account name, the lpPassword parameter must be NULL.
                        ddAgentUserPassword = null;
                    }
                }

                _session.Log($"Configuring services with account {ddAgentUserName}");
                // ddagentuser
                _serviceController.SetCredentials("datadogagent", ddAgentUserName, ddAgentUserPassword);
                _serviceController.SetCredentials("datadog-trace-agent", ddAgentUserName, ddAgentUserPassword);
                // SYSTEM
                // LocalSystem is a SCM specific shorthand that doesn't need to be localized
                _serviceController.SetCredentials("datadog-system-probe", "LocalSystem", "");
                _serviceController.SetCredentials( "datadog-process-agent", "LocalSystem", "");
            }
            catch (Exception e)
            {
                _session.Log($"Failed to configure service logon users: {e.Message}");
                _session.Log(e.ToString());
                return ActionResult.Failure;
            }
            return ActionResult.Success;
        }

        [CustomAction]
        public static ActionResult ConfigureServiceUsers(Session session)
        {
            return new ServiceCustomAction(new SessionWrapper(session)).ConfigureServiceUsers();
        }

        /// <summary>
        /// Stop any existing datadog services
        /// </summary>
        /// <returns></returns>
        private ActionResult StopDDServices()
        {
            try
            {
                // Stop each service individually in case the install is broken
                // e.g. datadogagent doesn't exist or the service dependencies are not correect.
                var ddservices = new string[]
                {
                    "datadog-system-probe",
                    "ddnpm",
                    "datadog-process-agent",
                    "datadog-trace-agent",
                    "datadogagent"
                };
                foreach (var service in ddservices)
                {
                    var svcNames = _serviceController.GetServiceNames().FirstOrDefault(svc => svc.Item1 == service);
                    if (svcNames != null)
                    {
                        using var actionRecord = new Record(
                            "Stop Datadog services",
                            $"Stopping {svcNames.Item2} service",
                            ""
                        );
                        _session.Message(InstallMessage.ActionStart, actionRecord);
                        _session.Log($"Stopping service {service}");
                        _serviceController.StopService(service, TimeSpan.FromMinutes(3));
                    }
                    else
                    {
                        _session.Log($"Service {service} not found");
                    }
                }
            }
            catch (Exception e)
            {
                _session.Log($"Failed to stop services: {e.Message}");
                _session.Log(e.ToString());
                return ActionResult.Failure;
            }
            return ActionResult.Success;
        }

        [CustomAction]
        public static ActionResult StopDDServices(Session session)
        {
            return new ServiceCustomAction(new SessionWrapper(session)).StopDDServices();
        }

        private ActionResult StartDDServices()
        {
            try
            {
                using var actionRecord = new Record(
                    "Start Datadog services",
                    "Starting Datadog Agent service",
                    ""
                );
                _session.Message(InstallMessage.ActionStart, actionRecord);
                // only start the main agent service. it should start any other services
                // that should be running.
                _serviceController.StartService("datadogagent", TimeSpan.FromMinutes(3));
            }
            catch (Exception e)
            {
                _session.Log($"Failed to stop services: {e.Message}");
                _session.Log(e.ToString());
                // Allow service start to fail and continue the install
            }
            return ActionResult.Success;
        }

        [CustomAction]
        public static ActionResult StartDDServices(Session session)
        {
            return new ServiceCustomAction(new SessionWrapper(session)).StartDDServices();
        }
    }
}
