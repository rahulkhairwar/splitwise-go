package splitwise

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type RestClient struct {
	HttpClient *http.Client
	UserAgent  string
}

func (c *RestClient) Get(ctx context.Context, path string, params map[string]string) (*http.Response, error) {
	return c.do(ctx, http.MethodGet, path, nil, params)
}

func (c *RestClient) Head(ctx context.Context, path string, body interface{}, params map[string]string) (*http.Response, error) {
	return c.do(ctx, http.MethodHead, path, body, params)
}

func (c *RestClient) Post(ctx context.Context, path string, body interface{}, params map[string]string) (*http.Response, error) {
	return c.do(ctx, http.MethodPost, path, body, params)
}

func (c *RestClient) Put(ctx context.Context, path string, body interface{}, params map[string]string) (*http.Response, error) {
	return c.do(ctx, http.MethodPut, path, body, params)
}

func (c *RestClient) Patch(ctx context.Context, path string, body interface{}, params map[string]string) (*http.Response, error) {
	return c.do(ctx, http.MethodPatch, path, body, params)
}

func (c *RestClient) Delete(ctx context.Context, path string, body interface{}, params map[string]string) (*http.Response, error) {
	return c.do(ctx, http.MethodDelete, path, body, params)
}

func (c *RestClient) Connect(ctx context.Context, path string, body interface{}, params map[string]string) (*http.Response, error) {
	return c.do(ctx, http.MethodConnect, path, body, params)
}

func (c *RestClient) Options(ctx context.Context, path string, body interface{}, params map[string]string) (*http.Response, error) {
	return c.do(ctx, http.MethodOptions, path, body, params)
}

func (c *RestClient) Trace(ctx context.Context, path string, body interface{}, params map[string]string) (*http.Response, error) {
	return c.do(ctx, http.MethodTrace, path, body, params)
}

func (c *RestClient) do(ctx context.Context, method, url string, body interface{}, params map[string]string) (*http.Response, error) {
	req, err := c.newRequest(method, url, body)
	if err != nil {
		return nil, err
	}

	q := req.URL.Query()
	for k, v := range params {
		q.Add(k, v)
	}

	req.URL.RawQuery = q.Encode()
	req = req.WithContext(ctx)
	fmt.Printf("Sending {%s} request to {%s} with :\nbody = {%+v}\nparams = {%+v}\n", method, url, body, params)

	resp, err := c.HttpClient.Do(req)
	if err != nil {
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		default:
		}
		return nil, err
	}
	return resp, nil
}

func (c *RestClient) newRequest(method, url string, body interface{}) (*http.Request, error) {
	var buf io.ReadWriter
	if body != nil {
		buf = new(bytes.Buffer)
		err := json.NewEncoder(buf).Encode(body)
		if err != nil {
			return nil, err
		}
	}

	req, err := http.NewRequest(method, url, buf)
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
