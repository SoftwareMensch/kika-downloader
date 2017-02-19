package http

import (
	"encoding/json"
	"fmt"
	"testing"
	"strings"
)

type ipInfoDTO struct {
	Ip string
}

func TestSocksProxy(t *testing.T) {
	nonProxyClient, err := NewClient()
	if err != nil {
		t.Error(err)
	}

	resp, err := nonProxyClient.Get("https://ipinfo.io/json")
	if err != nil {
		t.Error(err)
	}
	defer resp.Body.Close()

	ipInfo := &ipInfoDTO{}
	json.NewDecoder(resp.Body).Decode(ipInfo)

	nonProxyIP := ipInfo.Ip

	proxyClient, err := NewClient()
	if err != nil {
		t.Error(err)
	}

	resp, err = proxyClient.ProxyGet("https://ipinfo.io/json")
	if err != nil {
		t.Error(err)
	}
	defer resp.Body.Close()

	ipInfo = &ipInfoDTO{}
	json.NewDecoder(resp.Body).Decode(ipInfo)

	proxyIP := ipInfo.Ip

	if strings.Compare(nonProxyIP, proxyIP) == 0 {
		t.Error(fmt.Errorf("proxy not working (%s = %s)", nonProxyIP, proxyIP))
	}
}
