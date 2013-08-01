// The ghstatus package allows you to check the system status of GitHub from
// your own applications. The status information is retrieved from GitHub's
// system status API (https://status.github.com/api).

package ghstatus

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// The root URL of GitHub's system status API.
var ServiceURL = "https://status.github.com"

// Possible status values.
const (
	StatusGood  = "good"
	StatusMinor = "minor"
	StatusMajor = "major"
)

// Current system status as returned by the /api/status.json endpoint.
type Status struct {
	Status      string `json:"status"`
	LastUpdated string `json:"last_updated"`
}

// A status message as returned by the /api/messages.json and
// /api/last-message.json endpoints.
type Message struct {
	Status    string `json:"status"`
	Body      string `json:"body"`
	CreatedOn string `json:"created_on"`
}

func sendRequest(endpoint string, v interface{}) error {
	resp, err := http.Get(ServiceURL + endpoint)
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
