package ghstatus

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

var fakeResponses = []struct {
	method, endpoint, body string
}{
	{
		"GET",
		"/api/status.json",
		`{"status":"good","last_updated":"2013-07-31T12:09:46Z"}`,
	},
	{
		"GET",
		"/api/messages.json",
		`[{
			"body": "Everything operating normally.",
			"created_on": "2013-07-29T22:23:19Z",
			"status": "good"
		},
		{
			"body": "We are continuing to work on the increased exception rate on the GitHub API. We will update again as soon as have the source of these exceptions resolved. ",
			"created_on": "2013-07-29T21:09:54Z",
			"status": "minor"
		},
		{
			"body": "We've mitigated the DDoS attack and the site should responding normally. We're still investigating the cause of the small increase in exceptions when accessing the GitHub API.",
			"created_on": "2013-07-29T16:10:54Z",
			"status": "minor"
		},
		{
			"body": "We're currently experiencing a large DDoS attack. The site is experiencing major packet loss and is mostly unavailable. We're working to further mitigate the attack.",
			"created_on": "2013-07-29T15:05:38Z",
			"status": "major"
		},
		{
			"body": "We're investigating a small increase in exceptions affecting the GitHub API.",
			"created_on": "2013-07-29T13:29:24Z",
			"status": "minor"
		}]`,
	},
	{
		"GET",
		"/api/last-message.json",
		`{"status":"good","body":"Everything operating normally.","created_on":"2013-07-29T22:23:19Z"}`,
	},
}

// Start internal webserver returning fake responses (unless REALHTTP is set in
// environment).
func init() {
	if os.Getenv("REALHTTP") != "" {
		return
	}
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		for _, f := range fakeResponses {
			if r.Method == f.method && r.URL.Path == f.endpoint {
				fmt.Fprint(w, f.body)
			}
		}
	}))
	StatusApiUrl = ts.URL
}

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
