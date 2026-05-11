package panel

import (
	"strings"
	"testing"
)

func TestDecodeUserListJSONTopLevel(t *testing.T) {
	body := `{"users":[{"id":1,"uuid":"u1","speed_limit":2,"device_limit":3}]}`
	userList, err := decodeUserListJSON(strings.NewReader(body))
	if err != nil {
		t.Fatal(err)
	}
	if len(userList.Users) != 1 || userList.Users[0].Id != 1 || userList.Users[0].Uuid != "u1" {
		t.Fatalf("unexpected users: %#v", userList.Users)
	}
}

func TestDecodeUserListJSONNestedFallback(t *testing.T) {
	body := `{"data":{"meta":{"page":1},"users":[{"id":7,"uuid":"nested"}]}}`
	userList, err := decodeUserListJSON(strings.NewReader(body))
	if err != nil {
		t.Fatal(err)
	}
	if len(userList.Users) != 1 || userList.Users[0].Id != 7 || userList.Users[0].Uuid != "nested" {
		t.Fatalf("unexpected users: %#v", userList.Users)
	}
}

func TestDecodeUserListJSONMissingUsers(t *testing.T) {
	_, err := decodeUserListJSON(strings.NewReader(`{"data":[]}`))
	if err == nil {
		t.Fatal("expected error for missing users array")
	}
}
