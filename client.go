package ghstatus

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"time"
)

type Client struct {
	httpClient *http.Client
}

func (c *Client) sendRequest(endpoint string, v interface{}) error {
	resp, err := c.httpClient.Get(serviceURL + endpoint)
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

func NewClient() *Client {
	timeout := 5 * time.Second
	dialFunc := func(network, address string) (net.Conn, error) {
		c, err := net.DialTimeout(network, address, timeout)
		if err != nil {
			return nil, err
		}
		c.SetDeadline(time.Now().Add(timeout))
		return c, nil
	}
	return &Client{
		httpClient: &http.Client{
			Transport: &http.Transport{Dial: dialFunc, DisableKeepAlives: true},
		},
	}
}
