package internal

import (
	"net/http"
	"net/url"
)

type Client struct {
	BaseUrl    *url.URL
	UserAgent  string
	httpClient *http.Client
}
