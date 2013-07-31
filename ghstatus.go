// The ghstatus package allows you to check the system status of GitHub from
// your own applications. The status information is retrieved from GitHub's
// system status API (https://status.github.com/api).

package ghstatus

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

// The root URL of GitHub's system status API.
const StatusApiUrl = "https://status.github.com/api"

// Possible status values.
const (
	StatusGood  = "good"
	StatusMinor = "minor"
	StatusMajor = "major"
)

// Current system status as returned by the /status endpoint.
type Status struct {
	Status      string `json:"status"`
	LastUpdated string `json:"last_updated"`
}

// A status message as returned by the /messages and /last-message endpoints.
type Message struct {
	Status    string `json:"status"`
	Body      string `json:"body"`
	CreatedOn string `json:"created_on"`
}

func sendRequest(endpoint string, v interface{}) error {
	resp, err := http.Get(StatusApiUrl + endpoint + ".json")
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	return json.Unmarshal(body, v)
}

// Get current system status and timestamp.
func GetStatus() (*Status, error) {
	var status *Status
	if err := sendRequest("/status", &status); err != nil {
		return nil, err
	}
	return status, nil
}

// Get most recent human communications with status and timestamp.
func GetMessages() ([]Message, error) {
	var messages []Message
	if err := sendRequest("/messages", &messages); err != nil {
		return nil, err
	}
	return messages, nil
}

// Get last human communication, status, and timestamp.
func GetLastMessage() (*Message, error) {
	var message *Message
	if err := sendRequest("/last-message", &message); err != nil {
		return nil, err
	}
	return message, nil
}
