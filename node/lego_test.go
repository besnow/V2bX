package node

import (
	"os"
	"testing"

	"github.com/InazumaV/V2bX/conf"
)

func newTestLego(t *testing.T) *Lego {
	t.Helper()
	if os.Getenv("V2BX_RUN_ACME_INTEGRATION") != "1" {
		t.Skip("skipping live ACME integration test; set V2BX_RUN_ACME_INTEGRATION=1 to run")
	}
	l, err := NewLego(&conf.CertConfig{
		CertMode:   "dns",
		Email:      "test@test.com",
		CertDomain: "test.test.com",
		Provider:   "cloudflare",
		DNSEnv: map[string]string{
			"CF_DNS_API_TOKEN": "123",
		},
		CertFile: "./cert/1.pem",
		KeyFile:  "./cert/1.key",
	})
	if err != nil {
		t.Fatalf("create lego client: %v", err)
	}
	return l
}

func TestLego_CreateCertByDns(t *testing.T) {
	l := newTestLego(t)
	if err := l.CreateCert(); err != nil {
		t.Error(err)
	}
}

func TestLego_RenewCert(t *testing.T) {
	l := newTestLego(t)
	t.Log(l.RenewCert())
}
