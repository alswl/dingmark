package message

import "testing"

var m = &Message{"text"}

func TestGetType(t *testing.T) {
	if m.GetType() != "text" {
		t.Errorf("get message type not match")
	}
}

func TestSetType(t *testing.T) {
	m.SetType("link")
	if m.Type != "link" {
		t.Errorf("can't set message type")
	}
}
