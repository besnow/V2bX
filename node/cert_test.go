package node

import (
	"path/filepath"
	"testing"
)

func Test_generateSelfSslCertificate(t *testing.T) {
	dir := t.TempDir()
	certFile := filepath.Join(dir, "1.pem")
	keyFile := filepath.Join(dir, "1.key")
	if err := generateSelfSslCertificate("domain.com", certFile, keyFile); err != nil {
		t.Fatal(err)
	}
}
