package xray

import (
	"reflect"
	"testing"

	"github.com/InazumaV/V2bX/conf"
	coreConf "github.com/xtls/xray-core/infra/conf"
)

func TestIsHTTPTransportIncludesXHTTPAndSplitHTTP(t *testing.T) {
	for _, network := range []string{"xhttp", "splithttp", "ws", "grpc", "httpupgrade"} {
		t.Run(network, func(t *testing.T) {
			if !isHTTPTransport(network) {
				t.Fatalf("expected %q to be treated as an HTTP transport", network)
			}
		})
	}

	if isHTTPTransport("tcp") {
		t.Fatal("expected tcp to remain a non-HTTP transport")
	}
}

func TestXHTTPAndSplitHTTPSetTrustedXForwardedFor(t *testing.T) {
	for _, network := range []string{"xhttp", "splithttp"} {
		t.Run(network, func(t *testing.T) {
			stream := &coreConf.StreamConfig{}
			trusted := []string{"X-Real-IP"}

			applyInboundTransportOptions(stream, network, &conf.XrayOptions{}, trusted)

			if !reflect.DeepEqual(stream.SocketSettings.TrustedXForwardedFor, trusted) {
				t.Fatalf("expected %s Trusted XFF %v, got %v", network, trusted, stream.SocketSettings.TrustedXForwardedFor)
			}
		})
	}
}

func TestTrustedXForwardedForPanelValueTakesPrecedence(t *testing.T) {
	stream := &coreConf.StreamConfig{}
	trusted := []string{"CF-Connecting-IP", "X-Real-IP"}

	setTrustedXForwardedFor(stream, trusted)

	if !reflect.DeepEqual(stream.SocketSettings.TrustedXForwardedFor, trusted) {
		t.Fatalf("expected panel Trusted XFF %v, got %v", trusted, stream.SocketSettings.TrustedXForwardedFor)
	}
}

func TestTrustedXForwardedForDefaultsWhenEmpty(t *testing.T) {
	stream := &coreConf.StreamConfig{}

	setTrustedXForwardedFor(stream, nil)

	expected := []string{"X-Forwarded-For"}
	if !reflect.DeepEqual(stream.SocketSettings.TrustedXForwardedFor, expected) {
		t.Fatalf("expected default Trusted XFF %v, got %v", expected, stream.SocketSettings.TrustedXForwardedFor)
	}
}

func TestTCPDoesNotSetTrustedXForwardedFor(t *testing.T) {
	stream := &coreConf.StreamConfig{}

	applyInboundTransportOptions(stream, "tcp", &conf.XrayOptions{}, []string{"X-Real-IP"})

	if stream.SocketSettings != nil && len(stream.SocketSettings.TrustedXForwardedFor) > 0 {
		t.Fatalf("expected tcp to leave Trusted XFF unset, got %v", stream.SocketSettings.TrustedXForwardedFor)
	}
}

func TestApplyInboundTransportOptionsProxyProtocol(t *testing.T) {
	tests := []struct {
		name                string
		network             string
		enableProxyProtocol bool
	}{
		{name: "tcp disabled", network: "tcp", enableProxyProtocol: false},
		{name: "tcp enabled", network: "tcp", enableProxyProtocol: true},
		{name: "ws disabled", network: "ws", enableProxyProtocol: false},
		{name: "ws enabled", network: "ws", enableProxyProtocol: true},
		{name: "grpc disabled", network: "grpc", enableProxyProtocol: false},
		{name: "grpc enabled", network: "grpc", enableProxyProtocol: true},
		{name: "xhttp disabled", network: "xhttp", enableProxyProtocol: false},
		{name: "xhttp enabled", network: "xhttp", enableProxyProtocol: true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			stream := &coreConf.StreamConfig{}
			applyInboundTransportOptions(stream, tt.network, &conf.XrayOptions{EnableProxyProtocol: tt.enableProxyProtocol}, nil)

			switch tt.network {
			case "tcp":
				if stream.TCPSettings == nil {
					t.Fatal("expected TCP settings")
				}
				if stream.TCPSettings.AcceptProxyProtocol != tt.enableProxyProtocol {
					t.Fatalf("expected TCP proxy protocol %v, got %v", tt.enableProxyProtocol, stream.TCPSettings.AcceptProxyProtocol)
				}
			case "ws":
				if stream.WSSettings == nil {
					t.Fatal("expected WS settings")
				}
				if stream.WSSettings.AcceptProxyProtocol != tt.enableProxyProtocol {
					t.Fatalf("expected WS proxy protocol %v, got %v", tt.enableProxyProtocol, stream.WSSettings.AcceptProxyProtocol)
				}
			default:
				if stream.SocketSettings == nil {
					t.Fatal("expected socket settings")
				}
				if stream.SocketSettings.AcceptProxyProtocol != tt.enableProxyProtocol {
					t.Fatalf("expected socket proxy protocol %v, got %v", tt.enableProxyProtocol, stream.SocketSettings.AcceptProxyProtocol)
				}
			}
		})
	}
}

func TestTrustedXForwardedForDoesNotOverwriteSocketOptions(t *testing.T) {
	stream := &coreConf.StreamConfig{SocketSettings: &coreConf.SocketConfig{}}
	trusted := []string{"X-Real-IP"}

	applyInboundTransportOptions(stream, "xhttp", &conf.XrayOptions{EnableProxyProtocol: true, EnableTFO: true}, trusted)

	if stream.SocketSettings == nil {
		t.Fatal("expected socket settings")
	}
	if stream.SocketSettings.AcceptProxyProtocol != true {
		t.Fatalf("expected ProxyProtocol to remain enabled, got %v", stream.SocketSettings.AcceptProxyProtocol)
	}
	if stream.SocketSettings.TFO != true {
		t.Fatalf("expected TFO to remain enabled, got %#v", stream.SocketSettings.TFO)
	}
	if !reflect.DeepEqual(stream.SocketSettings.TrustedXForwardedFor, trusted) {
		t.Fatalf("expected Trusted XFF %v, got %v", trusted, stream.SocketSettings.TrustedXForwardedFor)
	}
}

func TestFallbackAndRealityXverArePreserved(t *testing.T) {
	fallbacks, err := buildVlessFallbacks([]conf.FallBackConfigForXray{{Dest: "127.0.0.1:8080", ProxyProtocolVer: 2}})
	if err != nil {
		t.Fatalf("buildVlessFallbacks returned error: %v", err)
	}
	if fallbacks[0].Xver != 2 {
		t.Fatalf("expected fallback Xver 2, got %d", fallbacks[0].Xver)
	}

	if got := realityXver(1, 2); got != 1 {
		t.Fatalf("expected TLS Reality Xver to take precedence, got %d", got)
	}
	if got := realityXver(0, 2); got != 2 {
		t.Fatalf("expected RealityConfig Xver fallback, got %d", got)
	}
}
