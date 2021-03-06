// These tests for the ghstatus package start an internal web server that
// returns fake responses. To talk to the real service, set the environment
// variable REALHTTP.
package ghstatus_test

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/mlafeldt/ghstatus"
	"github.com/stretchr/testify/assert"
)

var client = ghstatus.DefaultClient

func init() {
	if os.Getenv("REALHTTP") == "" {
		ts := httptest.NewServer(http.FileServer(http.Dir("testdata")))
		client = &ghstatus.Client{ServiceURL: ts.URL}
	}
}

func checkStatus(s string) bool {
	switch s {
	case ghstatus.Good, ghstatus.Minor, ghstatus.Major:
		return true
	}
	return false
}

func TestGetStatus(t *testing.T) {
	status, err := client.GetStatus()
	if assert.NoError(t, err) {
		assert.True(t, checkStatus(status.Status))
		assert.False(t, status.LastUpdated.IsZero())
	}
}

func TestGetMessages(t *testing.T) {
	messages, err := client.GetMessages()
	if assert.NoError(t, err) {
		for _, m := range messages {
			assert.True(t, checkStatus(m.Status))
			assert.NotEmpty(t, m.Body)
			assert.False(t, m.CreatedOn.IsZero())
		}
	}
}

func TestGetLastMessage(t *testing.T) {
	m, err := client.GetLastMessage()
	if assert.NoError(t, err) {
		assert.True(t, checkStatus(m.Status))
		assert.NotEmpty(t, m.Body)
		assert.False(t, m.CreatedOn.IsZero())
	}
}
