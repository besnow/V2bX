package node

import (
	"testing"

	"github.com/InazumaV/V2bX/api/panel"
)

func TestCompareUserList(t *testing.T) {
	tests := []struct {
		name         string
		old          []panel.UserInfo
		new          []panel.UserInfo
		wantDeleted  int
		wantAdded    int
		wantModified int
	}{
		{
			name: "no changes",
			old:  []panel.UserInfo{{Uuid: "u1", Id: 1, SpeedLimit: 10, DeviceLimit: 2}},
			new:  []panel.UserInfo{{Uuid: "u1", Id: 1, SpeedLimit: 10, DeviceLimit: 2}},
		},
		{
			name:      "added user",
			old:       []panel.UserInfo{{Uuid: "u1", Id: 1}},
			new:       []panel.UserInfo{{Uuid: "u1", Id: 1}, {Uuid: "u2", Id: 2}},
			wantAdded: 1,
		},
		{
			name:        "deleted user",
			old:         []panel.UserInfo{{Uuid: "u1", Id: 1}, {Uuid: "u2", Id: 2}},
			new:         []panel.UserInfo{{Uuid: "u1", Id: 1}},
			wantDeleted: 1,
		},
		{
			name:         "speedlimit modified",
			old:          []panel.UserInfo{{Uuid: "u1", Id: 1, SpeedLimit: 10}},
			new:          []panel.UserInfo{{Uuid: "u1", Id: 1, SpeedLimit: 20}},
			wantModified: 1,
		},
		{
			name:         "devicelimit modified",
			old:          []panel.UserInfo{{Uuid: "u1", Id: 1, DeviceLimit: 1}},
			new:          []panel.UserInfo{{Uuid: "u1", Id: 1, DeviceLimit: 2}},
			wantModified: 1,
		},
		{
			name: "uuid same and limits unchanged not modified",
			old:  []panel.UserInfo{{Uuid: "u1", Id: 1, SpeedLimit: 10, DeviceLimit: 2}},
			new:  []panel.UserInfo{{Uuid: "u1", Id: 1, SpeedLimit: 10, DeviceLimit: 2}},
		},

		{
			name: "uuid same id changed only not modified",
			old:  []panel.UserInfo{{Uuid: "u1", Id: 1, SpeedLimit: 10, DeviceLimit: 2}},
			new:  []panel.UserInfo{{Uuid: "u1", Id: 2, SpeedLimit: 10, DeviceLimit: 2}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			deleted, added, modified := compareUserList(tt.old, tt.new)
			if len(deleted) != tt.wantDeleted {
				t.Fatalf("deleted count mismatch, got %d want %d", len(deleted), tt.wantDeleted)
			}
			if len(added) != tt.wantAdded {
				t.Fatalf("added count mismatch, got %d want %d", len(added), tt.wantAdded)
			}
			if len(modified) != tt.wantModified {
				t.Fatalf("modified count mismatch, got %d want %d", len(modified), tt.wantModified)
			}
		})
	}
}
