package internal

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

type RegistrationStatus int

const (
	RegistrationStatus_DUMMY = iota
	RegistrationStatus_INVITED
	RegistrationStatus_CONFIRMED
)

var registrationStatuses = [...]string{"dummy", "invited", "confirmed"}

func (r RegistrationStatus) String() string {
	return registrationStatuses[r]
}

func GetRegistrationStatus(s string) (RegistrationStatus, error) {
	for i, rs := range registrationStatuses {
		if rs == s {
			return RegistrationStatus(i), nil
		}
	}

	return -1, errors.New(fmt.Sprintf("invalid registration status : %+v", s))
}

type User struct {
	Id        int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Picture   struct {
		Small  string `json:"small"`
		Medium string `json:"medium"`
		Large  string `json:"large"`
	} `json:"picture"`
	Email string `json:"email"`
	RegistrationStatus RegistrationStatus `json:"registration_status"`
	DefaultCurrency string `json:"default_currency"`
	Locale string `json:"locale"`
	NotificationsRead string `json:"notifications_read"`
	NotificationsCount int `json:"notifications_count"`
	NotificationsPreferences struct {
		AddedAsFriend bool `json:"added_as_friend"`
	} `json:"notifications"`
}

func (c* Client) GetCurrentUser() (*User, error) {
	req, err := c.Get("/get_current_user", nil)

	var user *User
	_, err = c.do(req, &user)

	return user, err
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
