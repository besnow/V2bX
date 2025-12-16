package conf

import (
	"os"
	"testing"
	"time"
)

func TestConf_LoadFromPath(t *testing.T) {
	c := New()
	t.Log(c.LoadFromPath("../example/config.json"), c.NodeConfig)
}

func TestConf_Watch(t *testing.T) {
	tmp, err := os.CreateTemp("", "conf_watch_*.json")
	if err != nil {
		t.Fatalf("create temp config: %v", err)
	}
	tmp.Close()
	defer os.Remove(tmp.Name())

	c := New()
	if err := c.Watch(tmp.Name(), "", "", func() {}); err != nil {
		t.Fatalf("watch temp config: %v", err)
	}

	time.Sleep(50 * time.Millisecond)
}
