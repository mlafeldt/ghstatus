// These tests for the ghstatus package start an internal web server that
// returns fake responses. To talk to the real service, set the environment
// variable REALHTTP.
package ghstatus

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/stretchr/testify/assert"
	"net/http/httptest"
	"testing"
)

func init() {
	if os.Getenv("REALHTTP") == "" {
		ts := httptest.NewServer(http.HandlerFunc(serveTestResponses))
		SetServiceURL(ts.URL)
	}
}

func serveTestResponses(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "method must be GET", http.StatusMethodNotAllowed)
		return
	}
	content, err := ioutil.ReadFile("testdata" + r.URL.Path)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Fprint(w, string(content))
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
