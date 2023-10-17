package http

import (
	"bytes"
	"github.com/pkg/errors"
	"io"
	"net/http"
	"strings"
)

var (
	longConnection *Client
)

type Client struct {
	client http.Client
}

// NewClient create http client
// When keepAlive is true, the connection can be reused, vice versa.
func NewClient(keepAlive bool) *Client {
	if keepAlive {
		if longConnection != nil {
			return longConnection
		} else {
			c := &Client{
				client: newRawClient(!keepAlive),
			}
			longConnection = c
			return longConnection
		}
	} else {
		return &Client{
			client: newRawClient(!keepAlive),
		}
	}
}

// Do
func (c *Client) Do(method, url string, header map[string]string, body []byte) ([]byte, error) {
	var result []byte
	buffer := bytes.NewBuffer(body)
	req, err := http.NewRequest(strings.ToUpper(method), url, buffer)
	if err != nil {
		return nil, err
	}
	// add header
	for label, value := range header {
		req.Header.Add(label, value)
	}
	// request
	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	// get response
	if resp.StatusCode != http.StatusOK {
		return nil, errors.New(resp.Status)
	}
	if result, err = io.ReadAll(resp.Body); err != nil {
		return nil, err
	}
	return result, nil
}

// Close release http client
func (c *Client) Close() {
	c.client.CloseIdleConnections()
}

func newRawClient(disableKeepAlive bool) http.Client {
	return http.Client{
		Transport: &http.Transport{
			DisableKeepAlives: disableKeepAlive,
		},
	}
}
