// These tests for the ghstatus package start an internal web server that
// returns fake responses. To talk to the real service, set the environment
// variable REALHTTP.

package ghstatus

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

var testResponses = map[string]string{
	"GET /api/status.json":       `{"status":"good","last_updated":"2013-07-31T12:09:46Z"}`,
	"GET /api/last-message.json": `{"status":"good","body":"Everything operating normally.","created_on":"2013-07-29T22:23:19Z"}`,
	"GET /api/messages.json": `[
		{
			"body": "Everything operating normally.",
			"created_on": "2013-07-29T22:23:19Z",
			"status": "good"
		},
		{
			"body": "We are continuing to work on the increased exception rate on the GitHub API.",
			"created_on": "2013-07-29T21:09:54Z",
			"status": "minor"
		},
		{
			"body": "We've mitigated the DDoS attack and the site should responding normally.",
			"created_on": "2013-07-29T16:10:54Z",
			"status": "minor"
		},
		{
			"body": "We're currently experiencing a large DDoS attack.",
			"created_on": "2013-07-29T15:05:38Z",
			"status": "major"
		},
		{
			"body": "We're investigating a small increase in exceptions affecting the GitHub API.",
			"created_on": "2013-07-29T13:29:24Z",
			"status": "minor"
		}]`,
}

func init() {
	if os.Getenv("REALHTTP") != "" {
		return
	}
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if body := testResponses[r.Method+" "+r.URL.Path]; body != "" {
			fmt.Fprint(w, body)
		} else {
			http.Error(w, "", http.StatusNotFound)
		}
	}))
	SetServiceURL(ts.URL)
}

func checkStatus(s string) bool {
	switch s {
	case StatusGood, StatusMinor, StatusMajor:
		return true
	}
	return false
}

func TestGetStatus(t *testing.T) {
	status, err := GetStatus()
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%+v", status)
	if !checkStatus(status.Status) {
		t.Errorf("Invalid Status: %s", status.Status)
	}
	if status.LastUpdated.IsZero() {
		t.Error("LastUpdated is zero")
	}
}

func TestGetMessages(t *testing.T) {
	messages, err := GetMessages()
	if err != nil {
		t.Fatal(err)
	}
	if len(messages) == 0 {
		t.Error("No messages returned")
	}
	for _, m := range messages {
		t.Logf("%+v", m)
		if !checkStatus(m.Status) {
			t.Errorf("Invalid Status: %s", m.Status)
		}
		if m.Body == "" {
			t.Error("Body empty")
		}
		if m.CreatedOn.IsZero() {
			t.Error("CreatedOn is zero")
		}
	}
}

func TestGetLastMessage(t *testing.T) {
	message, err := GetLastMessage()
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%+v", message)
	if !checkStatus(message.Status) {
		t.Errorf("Invalid Status: %s", message.Status)
	}
	if message.Body == "" {
		t.Error("Body empty")
	}
	if message.CreatedOn.IsZero() {
		t.Error("CreatedOn is zero")
	}
}
