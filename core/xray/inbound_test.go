package xray

import (
	"bytes"
	"testing"

	"github.com/InazumaV/V2bX/api/panel"
	"github.com/InazumaV/V2bX/conf"
)

func buildTestOptions() *conf.Options {
	return &conf.Options{
		ListenIP:    "0.0.0.0",
		XrayOptions: conf.NewXrayOptions(),
		CertConfig: &conf.CertConfig{
			CertMode: "file",
			CertFile: "test_data/1.pem",
			KeyFile:  "test_data/1.key",
		},
	}
}

func TestBuildHysteria2Inbound(t *testing.T) {
	opts := buildTestOptions()
	info := &panel.NodeInfo{
		Type:   "hysteria2",
		Common: &panel.CommonNode{ServerPort: 443},
		Hysteria2: &panel.Hysteria2Node{
			UpMbps:                  100,
			DownMbps:                200,
			Ignore_Client_Bandwidth: true,
			ObfsType:                "salamander",
			ObfsPassword:            "secret",
		},
	}
	inbound, err := buildInboundDetourConfig(opts, info, "test-hy2")
	if err != nil {
		t.Fatalf("buildInboundDetourConfig returned error: %v", err)
	}
	if inbound.Protocol != "hysteria2" {
		t.Fatalf("expected protocol hysteria2, got %s", inbound.Protocol)
	}
	if inbound.StreamSetting == nil || inbound.StreamSetting.Network == nil || *inbound.StreamSetting.Network != "hysteria2" {
		t.Fatalf("expected stream network hysteria2, got %+v", inbound.StreamSetting)
	}
	if inbound.StreamSetting.Security != "tls" {
		t.Fatalf("expected stream security tls, got %s", inbound.StreamSetting.Security)
	}
	if inbound.StreamSetting.TLSSettings == nil {
		t.Fatalf("expected TLS settings to be present")
	}
	if inbound.Settings == nil || !bytes.Contains(*inbound.Settings, []byte("up_mbps")) {
		t.Fatalf("expected hysteria2 settings to contain up_mbps")
	}
}

func TestBuildTuicInbound(t *testing.T) {
	opts := buildTestOptions()
	info := &panel.NodeInfo{
		Type:   "tuic",
		Common: &panel.CommonNode{ServerPort: 444},
		Tuic: &panel.TuicNode{
			CongestionControl: "bbr",
			ZeroRTTHandshake:  true,
		},
	}
	inbound, err := buildInboundDetourConfig(opts, info, "test-tuic")
	if err != nil {
		t.Fatalf("buildInboundDetourConfig returned error: %v", err)
	}
	if inbound.Protocol != "tuic" {
		t.Fatalf("expected protocol tuic, got %s", inbound.Protocol)
	}
	if inbound.StreamSetting == nil || inbound.StreamSetting.Network == nil || *inbound.StreamSetting.Network != "tuic" {
		t.Fatalf("expected stream network tuic, got %+v", inbound.StreamSetting)
	}
	if inbound.StreamSetting.Security != "tls" {
		t.Fatalf("expected stream security tls, got %s", inbound.StreamSetting.Security)
	}
	if inbound.StreamSetting.TLSSettings == nil {
		t.Fatalf("expected TLS settings to be present")
	}
	if inbound.Settings == nil || !bytes.Contains(*inbound.Settings, []byte("congestionControl")) {
		t.Fatalf("expected tuic settings to contain congestionControl")
	}
}

func TestBuildAnyTLSInbound(t *testing.T) {
	opts := buildTestOptions()
	info := &panel.NodeInfo{
		Type:   "anytls",
		Common: &panel.CommonNode{ServerPort: 445},
		AnyTls: &panel.AnyTlsNode{
			PaddingScheme: []string{"random"},
		},
	}
	inbound, err := buildInboundDetourConfig(opts, info, "test-anytls")
	if err != nil {
		t.Fatalf("buildInboundDetourConfig returned error: %v", err)
	}
	if inbound.Protocol != "anytls" {
		t.Fatalf("expected protocol anytls, got %s", inbound.Protocol)
	}
	if inbound.StreamSetting == nil || inbound.StreamSetting.Network == nil || *inbound.StreamSetting.Network != "tcp" {
		t.Fatalf("expected stream network tcp, got %+v", inbound.StreamSetting)
	}
	if inbound.StreamSetting.Security != "tls" {
		t.Fatalf("expected stream security tls, got %s", inbound.StreamSetting.Security)
	}
	if inbound.StreamSetting.TLSSettings == nil {
		t.Fatalf("expected TLS settings to be present")
	}
	if inbound.Settings == nil || !bytes.Contains(*inbound.Settings, []byte("paddingScheme")) {
		t.Fatalf("expected anytls settings to contain paddingScheme")
	}
}
