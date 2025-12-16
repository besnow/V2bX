//go:build integration
// +build integration

package node

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/InazumaV/V2bX/conf"
)

func newIntegrationLego(t *testing.T) *Lego {
	t.Helper()

	token := os.Getenv("CF_DNS_API_TOKEN")
	domain := os.Getenv("TEST_CERT_DOMAIN")
	if token == "" || domain == "" {
		t.Skip("CF_DNS_API_TOKEN and TEST_CERT_DOMAIN must be set for integration tests")
	}

	certPath := filepath.Join("./cert", fmt.Sprintf("%s.pem", domain))
	keyPath := filepath.Join("./cert", fmt.Sprintf("%s.key", domain))

	lego, err := NewLego(&conf.CertConfig{
		CertMode:   "dns",
		Email:      "integration@example.com",
		CertDomain: domain,
		Provider:   "cloudflare",
		DNSEnv: map[string]string{
			"CF_DNS_API_TOKEN": token,
		},
		CertFile: certPath,
		KeyFile:  keyPath,
	})
	if err != nil {
		t.Fatalf("failed to create lego client: %v", err)
	}

	return lego
}

func TestLego_CreateCertByDns(t *testing.T) {
	l := newIntegrationLego(t)

	if err := l.CreateCert(); err != nil {
		t.Fatalf("CreateCert() error = %v", err)
	}
}

func TestLego_RenewCert(t *testing.T) {
	l := newIntegrationLego(t)

	if _, err := os.Stat(l.config.CertFile); err != nil {
		if os.IsNotExist(err) {
			t.Skip("certificate file missing; run TestLego_CreateCertByDns first")
		}
		t.Fatalf("unexpected error checking cert file: %v", err)
	}

	if err := l.RenewCert(); err != nil {
		t.Fatalf("RenewCert() error = %v", err)
	}
}
