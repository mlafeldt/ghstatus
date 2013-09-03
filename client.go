package ghstatus

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const (
	serviceURL = "https://status.github.com"
)

type Client struct {
	httpClient *http.Client
	ServiceURL string
}

var DefaultClient = &Client{
	httpClient: http.DefaultClient,
	ServiceURL: serviceURL,
}

func NewClient() (*Client, error) {
	c := DefaultClient
	return c, nil
}

func (c *Client) sendRequest(endpoint string, v interface{}) error {
	resp, err := c.httpClient.Get(c.ServiceURL + endpoint)
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
