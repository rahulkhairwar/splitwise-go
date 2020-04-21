package splitwise

import (
	"context"
	"errors"
	"fmt"
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
	Email                    string             `json:"email"`
	RegistrationStatus       RegistrationStatus `json:"registration_status"`
	DefaultCurrency          string             `json:"default_currency"`
	Locale                   string             `json:"locale"`
	NotificationsRead        string             `json:"notifications_read"`
	NotificationsCount       int                `json:"notifications_count"`
	NotificationsPreferences struct {
		AddedAsFriend bool `json:"added_as_friend"`
	} `json:"notifications"`
}

func (c *Client) GetCurrentUser(ctx context.Context) (*User, error) {
	req, err := c.Get("/get_current_user", nil)

	var user *User

	_, err = c.do(ctx, req, &user)
	if err != nil {
		return nil, err
	}

	return user, err
}

func (c *Client) GetUser(ctx context.Context, id int) (*User, error) {
	req, err := c.Get(fmt.Sprintf("/get_user/:%d", id), nil)

	var user *User

	_, err = c.do(ctx, req, &user)
	if err != nil {
		return nil, err
	}

	return user, err
}
