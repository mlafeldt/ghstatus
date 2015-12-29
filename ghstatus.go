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
const ServiceURL = "https://status.github.com"

// A client to talk to GitHub's system status API.
type Client struct {
	HTTPClient *http.Client
	ServiceURL string
}

// The default client to talk to GitHub's system status API.
var DefaultClient = &Client{
	HTTPClient: http.DefaultClient,
	ServiceURL: ServiceURL,
}

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

// Get current system status and timestamp.
func (c *Client) GetStatus() (*Status, error) {
	var status *Status
	if err := c.sendRequest("/api/status.json", &status); err != nil {
		return nil, err
	}
	return status, nil
}

// Get most recent human communications with status and timestamp.
func (c *Client) GetMessages() ([]Message, error) {
	var messages []Message
	if err := c.sendRequest("/api/messages.json", &messages); err != nil {
		return nil, err
	}
	return messages, nil
}

// Get last human communication, status, and timestamp.
func (c *Client) GetLastMessage() (*Message, error) {
	var message *Message
	if err := c.sendRequest("/api/last-message.json", &message); err != nil {
		return nil, err
	}
	return message, nil
}

func (c *Client) sendRequest(endpoint string, v interface{}) error {
	if c.HTTPClient == nil {
		c.HTTPClient = http.DefaultClient
	}
	if c.ServiceURL == "" {
		c.ServiceURL = ServiceURL
	}

	resp, err := c.HTTPClient.Get(c.ServiceURL + endpoint)
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
