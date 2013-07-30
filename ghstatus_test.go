package ghstatus

import "testing"

func TestGetStatus(t *testing.T) {
	status, err := GetStatus()
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%+v", status)
}

func TestGetMessages(t *testing.T) {
	messages, err := GetMessages()
	if err != nil {
		t.Fatal(err)
	}
	for _, m := range messages {
		t.Logf("%+v", m)
	}
}

func TestGetLastMessage(t *testing.T) {
	message, err := GetLastMessage()
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%+v", message)
}
