package xray

import (
	"bytes"
	"encoding/json"
	"testing"
)

func TestMigrateTUICInboundConfigIDsUsers(t *testing.T) {
	input := []byte(`[{"protocol":"tuic","settings":{"users":[{"uuid":"user-1","password":"pw"}],"tag":"keep"}}]`)

	output := migrateTUICInboundConfigIDs(input)
	configs := decodeConfigList(t, output)
	settings := objectValue(t, configs[0], "settings")
	users := listValue(t, settings, "users")
	user := users[0].(map[string]interface{})

	assertNoKey(t, user, "uuid")
	assertValue(t, user, "id", "user-1")
	assertValue(t, user, "password", "pw")
	assertValue(t, settings, "tag", "keep")
}

func TestMigrateTUICOutboundConfigIDsSettings(t *testing.T) {
	input := []byte(`[{"protocol":"tuic","settings":{"uuid":"outbound-id","password":"pw","congestion_control":"bbr"}}]`)

	output := migrateTUICOutboundConfigIDs(input)
	configs := decodeConfigList(t, output)
	settings := objectValue(t, configs[0], "settings")

	assertNoKey(t, settings, "uuid")
	assertValue(t, settings, "id", "outbound-id")
	assertValue(t, settings, "password", "pw")
	assertValue(t, settings, "congestion_control", "bbr")
}

func TestMigrateTUICOutboundConfigIDsServers(t *testing.T) {
	input := []byte(`[{"protocol":"tuic","settings":{"servers":[{"address":"example.com","uuid":"server-id","password":"pw"}]}}]`)

	output := migrateTUICOutboundConfigIDs(input)
	configs := decodeConfigList(t, output)
	settings := objectValue(t, configs[0], "settings")
	servers := listValue(t, settings, "servers")
	server := servers[0].(map[string]interface{})

	assertNoKey(t, server, "uuid")
	assertValue(t, server, "id", "server-id")
	assertValue(t, server, "address", "example.com")
	assertValue(t, server, "password", "pw")
}

func TestMigrateTUICConfigIDsKeepsExistingID(t *testing.T) {
	input := []byte(`[{"protocol":"tuic","settings":{"uuid":"legacy-id","id":"current-id","servers":[{"uuid":"legacy-server","id":"current-server","password":"pw"}]}}]`)

	output := migrateTUICOutboundConfigIDs(input)
	configs := decodeConfigList(t, output)
	settings := objectValue(t, configs[0], "settings")
	servers := listValue(t, settings, "servers")
	server := servers[0].(map[string]interface{})

	assertValue(t, settings, "id", "current-id")
	assertNoKey(t, settings, "uuid")
	assertValue(t, server, "id", "current-server")
	assertNoKey(t, server, "uuid")
	assertValue(t, server, "password", "pw")
}

func TestMigrateTUICConfigIDsLeavesNonTUICUnchanged(t *testing.T) {
	input := []byte(`[{"protocol":"freedom","settings":{"uuid":"unchanged","servers":[{"uuid":"server"}]}}]`)

	output := migrateTUICOutboundConfigIDs(input)
	if !bytes.Equal(output, input) {
		t.Fatalf("expected non-TUIC config to remain unchanged, got %s", output)
	}
	configs := decodeConfigList(t, output)
	settings := objectValue(t, configs[0], "settings")
	assertValue(t, settings, "uuid", "unchanged")
}

func TestMigrateTUICConfigIDsLeavesInvalidJSONUnchanged(t *testing.T) {
	input := []byte(`[{"protocol":"tuic","settings":{"uuid":"broken"}`)

	output := migrateTUICOutboundConfigIDs(input)
	if !bytes.Equal(output, input) {
		t.Fatalf("expected invalid JSON to remain unchanged, got %s", output)
	}
}

func decodeConfigList(t *testing.T, data []byte) []map[string]interface{} {
	t.Helper()
	var configs []map[string]interface{}
	if err := json.Unmarshal(data, &configs); err != nil {
		t.Fatalf("migrated JSON should unmarshal: %v", err)
	}
	return configs
}

func objectValue(t *testing.T, object map[string]interface{}, key string) map[string]interface{} {
	t.Helper()
	value, ok := object[key].(map[string]interface{})
	if !ok {
		t.Fatalf("expected %q to be an object, got %#v", key, object[key])
	}
	return value
}

func listValue(t *testing.T, object map[string]interface{}, key string) []interface{} {
	t.Helper()
	value, ok := object[key].([]interface{})
	if !ok {
		t.Fatalf("expected %q to be a list, got %#v", key, object[key])
	}
	return value
}

func assertNoKey(t *testing.T, object map[string]interface{}, key string) {
	t.Helper()
	if _, ok := object[key]; ok {
		t.Fatalf("expected %q to be deleted from %#v", key, object)
	}
}

func assertValue(t *testing.T, object map[string]interface{}, key string, expected interface{}) {
	t.Helper()
	if object[key] != expected {
		t.Fatalf("expected %q to be %#v, got %#v", key, expected, object[key])
	}
}
