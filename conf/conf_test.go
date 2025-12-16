package conf

import (
	"testing"
)

func TestConf_LoadFromPath(t *testing.T) {
	c := New()
	t.Log(c.LoadFromPath("../example/config.json"), c.NodeConfig)
}

func TestConf_Watch(t *testing.T) {
	// TODO: the watch helper blocks forever in tests; skip until it is
	// refactored to support context cancellation.
	t.Skip("watcher blocks without external events")
	c := New()
	t.Log(c.Watch("./1.json", "", "", func() {}))
	select {}
}
