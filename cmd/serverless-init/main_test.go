// Unless explicitly stated otherwise all files in this repository are licensed
// under the Apache License Version 2.0.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2016-present Datadog, Inc.

//go:build !windows
// +build !windows

package main

import (
	"github.com/DataDog/datadog-agent/cmd/serverless-init/tag"
	"github.com/DataDog/datadog-agent/pkg/config"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestProxyNotLoaded(t *testing.T) {
	proxyHttp := "abc:1234"
	proxyHttps := "abc:5678"
	t.Setenv("DD_PROXY_HTTP", proxyHttp)
	t.Setenv("DD_PROXY_HTTPS", proxyHttps)
	proxyHttpConfig := config.Datadog.GetString("proxy.http")
	proxyHttpsConfig := config.Datadog.GetString("proxy.https")
	assert.Equal(t, 0, len(proxyHttpConfig))
	assert.Equal(t, 0, len(proxyHttpsConfig))
}

func TestProxyLoaded(t *testing.T) {
	proxyHttp := "abc:1234"
	proxyHttps := "abc:5678"
	t.Setenv("DD_PROXY_HTTP", proxyHttp)
	t.Setenv("DD_PROXY_HTTPS", proxyHttps)
	setupProxy()
	proxyHttpConfig := config.Datadog.GetString("proxy.http")
	proxyHttpsConfig := config.Datadog.GetString("proxy.https")
	assert.Equal(t, proxyHttp, proxyHttpConfig)
	assert.Equal(t, proxyHttps, proxyHttpsConfig)
}

func TestTagsSetup(t *testing.T) {
	t.Setenv("DD_TAGS", "key1:value1 key2:value2 key3:value3")
	tags := map[string]string{
		"key1": "value1",
		"key2": "value2",
		"key3": "value3",
	}
	_, _, _, metricAgent := setup()
	assert.True(t, mapContains(tag.MergeTags(map[string]string{}, metricAgent.GetExtraTags()), tags))
}

func mapContains(subject map[string]string, valuesToContain map[string]string) bool {
	for k, v := range valuesToContain {
		if subject[k] != v {
			return false
		}
	}
	return true
}
