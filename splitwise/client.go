package splitwise

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/url"
)

type Client struct {
	clientId   string
	secret     string
	publicKey  string
	httpClient *http.Client

	BaseUrl   *url.URL
	UserAgent string
}

type ClientOptions struct {
	ClientId   string
	Secret     string
	PublicKey  string
	HttpClient *http.Client
	BaseUrl    *url.URL
	UserAgent  string
}

func NewClient(options ClientOptions) (*Client, error) {
	if options.HttpClient == nil {
		options.HttpClient = &http.Client{}
	}

	return &Client{
		clientId:   options.ClientId,
		secret:     options.Secret,
		publicKey:  options.PublicKey,
		httpClient: options.HttpClient,
		BaseUrl:    options.BaseUrl,
		UserAgent:  options.UserAgent,
	}, nil
}

func (c *Client) do(req *http.Request, v interface{}) (*http.Response, error) {
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(v)
	return resp, err
}

func (c *Client) Get(path string, body interface{}) (*http.Request, error) {
	return c.newRequest("GET", path, body)
}

func (c *Client) Post(path string, body interface{}) (*http.Request, error) {
	return c.newRequest("POST", path, body)
}

func (c *Client) Put(path string, body interface{}) (*http.Request, error) {
	return c.newRequest("PUT", path, body)
}

func (c *Client) Delete(path string, body interface{}) (*http.Request, error) {
	return c.newRequest("DELETE", path, body)
}

func (c *Client) newRequest(method, path string, body interface{}) (*http.Request, error) {
	rel := &url.URL{Path: path}
	u := c.BaseUrl.ResolveReference(rel)

	var buf io.ReadWriter
	if body != nil {
		buf = new(bytes.Buffer)
		err := json.NewEncoder(buf).Encode(body)
		if err != nil {
			return nil, err
		}
	}

	req, err := http.NewRequest(method, u.String(), buf)
	if err != nil {
		return nil, err
	}

	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}

	req.Header.Set("Accept", "application/json")
	req.Header.Set("User-Agent", c.UserAgent)

	return req, nil
}
