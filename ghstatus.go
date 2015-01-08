// The ghstatus package allows you to check the system status of GitHub from
// your own applications. The status information is retrieved from GitHub's
// system status API (https://status.github.com/api).
package ghstatus

import (
	"time"
)

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

// Get current URL for system status API.
func ServiceURL() string {
	return DefaultClient.ServiceURL
}

// Set new URL for system status API.
func SetServiceURL(url string) {
	DefaultClient.ServiceURL = url
}

// Get current system status and timestamp.
func (c *Client) GetStatus() (*Status, error) {
	var status *Status
	if err := c.sendRequest("/api/status.json", &status); err != nil {
		return nil, err
	}
	return status, nil
}

// GetStatus is a wrapper around DefaultClient.GetStatus
func GetStatus() (*Status, error) {
	return DefaultClient.GetStatus()
}

// Get most recent human communications with status and timestamp.
func (c *Client) GetMessages() ([]Message, error) {
	var messages []Message
	if err := c.sendRequest("/api/messages.json", &messages); err != nil {
		return nil, err
	}
	return messages, nil
}

// GetMessages is a wrapper around DefaultClient.GetMessages
func GetMessages() ([]Message, error) {
	return DefaultClient.GetMessages()
}

// Get last human communication, status, and timestamp.
func (c *Client) GetLastMessage() (*Message, error) {
	var message *Message
	if err := c.sendRequest("/api/last-message.json", &message); err != nil {
		return nil, err
	}
	return message, nil
}

// GetLastMessage is a wrapper around DefaultClient.GetLastMessage
func GetLastMessage() (*Message, error) {
	return DefaultClient.GetLastMessage()
}
