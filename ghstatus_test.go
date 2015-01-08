// These tests for the ghstatus package start an internal web server that
// returns fake responses. To talk to the real service, set the environment
// variable REALHTTP.
package ghstatus_test

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	ghstatus "github.com/mlafeldt/go-ghstatus"
	"github.com/stretchr/testify/assert"
)

func init() {
	if os.Getenv("REALHTTP") == "" {
		ts := httptest.NewServer(http.FileServer(http.Dir("testdata")))
		ghstatus.SetServiceURL(ts.URL)
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
	status, err := ghstatus.GetStatus()
	if assert.NoError(t, err) {
		assert.True(t, checkStatus(status.Status))
		assert.False(t, status.LastUpdated.IsZero())
	}
}

func TestGetMessages(t *testing.T) {
	messages, err := ghstatus.GetMessages()
	if assert.NoError(t, err) {
		for _, m := range messages {
			assert.True(t, checkStatus(m.Status))
			assert.NotEmpty(t, m.Body)
			assert.False(t, m.CreatedOn.IsZero())
		}
	}
}

func TestGetLastMessage(t *testing.T) {
	m, err := ghstatus.GetLastMessage()
	if assert.NoError(t, err) {
		assert.True(t, checkStatus(m.Status))
		assert.NotEmpty(t, m.Body)
		assert.False(t, m.CreatedOn.IsZero())
	}
}

func TestGetDailySummaries(t *testing.T) {
	summaries, err := GetDailySummaries()
	if assert.NoError(t, err) {
		assert.NotEmpty(t, summaries)
		s := summaries["2010-02-09"]
		assert.Equal(t, 25200.0, s.Good)
		assert.Equal(t, 61200.0, s.Minor)
		assert.Equal(t, 0, s.Major)
	}
}

func TestGetDailySummary(t *testing.T) {
	summary, err := GetDailySummary("2014-03-22")
	if assert.NoError(t, err) {
		assert.Equal(t, 77373.0, summary.Good)
		assert.Equal(t, 9027.0, summary.Minor)
		assert.Equal(t, 0, summary.Major)
	}
}
