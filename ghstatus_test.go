// These tests for the ghstatus package start an internal web server that
// returns fake responses. To talk to the real service, set the environment
// variable REALHTTP.
package ghstatus

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/stretchr/testify/assert"
	"net/http/httptest"
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

func serveTestResponses(w http.ResponseWriter, r *http.Request) {
	if body := testResponses[r.Method+" "+r.URL.Path]; body != "" {
		fmt.Fprint(w, body)
	} else {
		http.Error(w, "", http.StatusNotFound)
	}
}

func init() {
	if os.Getenv("REALHTTP") == "" {
		ts := httptest.NewServer(http.HandlerFunc(serveTestResponses))
		SetServiceURL(ts.URL)
	}
}

func checkStatus(s string) bool {
	switch s {
	case Good, Minor, Major:
		return true
	}
	return false
}

func TestGetStatus(t *testing.T) {
	status, err := GetStatus()
	if assert.NoError(t, err) {
		assert.True(t, checkStatus(status.Status))
		assert.False(t, status.LastUpdated.IsZero())
	}
}

func TestGetMessages(t *testing.T) {
	messages, err := GetMessages()
	if assert.NoError(t, err) {
		assert.NotEmpty(t, messages)
		for _, m := range messages {
			assert.True(t, checkStatus(m.Status))
			assert.NotEmpty(t, m.Body)
			assert.False(t, m.CreatedOn.IsZero())
		}
	}
}

func TestGetLastMessage(t *testing.T) {
	m, err := GetLastMessage()
	if assert.NoError(t, err) {
		assert.True(t, checkStatus(m.Status))
		assert.NotEmpty(t, m.Body)
		assert.False(t, m.CreatedOn.IsZero())
	}
}

func TestSetConnectionTimeout(t *testing.T) {
	timeout := 10 * time.Minute
	SetConnectionTimeout(timeout)
	assert.Equal(t, ConnectionTimeout(), timeout)
}
