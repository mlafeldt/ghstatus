// The ghstatus package allows you to check the system status of GitHub from
// your own applications. The status information is retrieved from GitHub's
// system status API (https://status.github.com/api).
package ghstatus

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

// The URL of GitHub's system status API.
var serviceURL = "https://status.github.com"

// Possible status values set in Status and Message.
const (
	Good  = "good"
	Minor = "minor"
	Major = "major"
)

// Current system status as returned by GetStatus.
type Status struct {
	Status      string    `json:"status"`
	LastUpdated time.Time `json:"last_updated"`
}

// A status message as returned by GetMessages and GetLastMessage.
type Message struct {
	Status    string    `json:"status"`
	Body      string    `json:"body"`
	CreatedOn time.Time `json:"created_on"`
}

type DailySummary struct {
	Good  float64 `json:"good"`
	Minor float64 `json:"minor"`
	Major float64 `json:"major"`
}

// Get current URL for system status API.
func ServiceURL() string {
	return serviceURL
}

// Set new URL for system status API.
func SetServiceURL(url string) {
	serviceURL = url
}

// Get current system status and timestamp.
func GetStatus() (*Status, error) {
	var status *Status
	if err := sendRequest("/api/status.json", &status); err != nil {
		return nil, err
	}
	return status, nil
}

// Get most recent human communications with status and timestamp.
func GetMessages() ([]Message, error) {
	var messages []Message
	if err := sendRequest("/api/messages.json", &messages); err != nil {
		return nil, err
	}
	return messages, nil
}

// Get last human communication, status, and timestamp.
func GetLastMessage() (*Message, error) {
	var message *Message
	if err := sendRequest("/api/last-message.json", &message); err != nil {
		return nil, err
	}
	return message, nil
}

func GetDailySummary(date string) (*DailySummary, error) {
	var a map[string]DailySummary
	if err := sendRequest("/api/daily-summary.json", &a); err != nil {
		return nil, err
	}
	summary := a[date]
	return &summary, nil
}

func sendRequest(endpoint string, v interface{}) error {
	resp, err := http.Get(serviceURL + endpoint)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("HTTP error: %s", resp.Status)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	return json.Unmarshal(body, v)
}
