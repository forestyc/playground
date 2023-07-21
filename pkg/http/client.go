package http

import (
	"bytes"
	"errors"
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

// NewClient 创建http client
// keepAlive为true是创建长连接client, 可重复使用
// keepAlive为false是创建短连接client
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

// Do 发起请求
func (c *Client) Do(method, url string, header map[string]string, body []byte) ([]byte, error) {
	var result []byte
	buffer := bytes.NewBuffer(body)
	req, err := http.NewRequest(convertMethod(method), url, buffer)
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

// Close 销毁Client
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

func convertMethod(method string) string {
	m := strings.ToUpper(method)
	switch m {
	case http.MethodGet:
		return m
	case http.MethodHead:
		return m
	case http.MethodPost:
		return m
	case http.MethodPut:
		return m
	case http.MethodPatch:
		return m
	case http.MethodDelete:
		return m
	case http.MethodConnect:
		return m
	case http.MethodOptions:
		return m
	case http.MethodTrace:
		return m
	default:
		return ""
	}
}
