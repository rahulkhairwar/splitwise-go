package splitwise

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/kr/pretty"
	"io/ioutil"
)

type userCommon struct {
	Id        int64  `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Picture   *struct {
		Small  string `json:"small,omitempty"`
		Medium string `json:"medium,omitempty"`
		Large  string `json:"large,omitempty"`
	} `json:"picture,omitempty"`
}

type User struct {
	userCommon
	Email                string                `json:"email,omitempty"`
	Password             string                `json:"password,omitempty"`
	RegistrationStatus   string                `json:"registration_status,omitempty"`
	DateFormat           string                `json:"date_format,omitempty"`
	DefaultCurrency      string                `json:"default_currency,omitempty"`
	DefaultGroupId       int64                 `json:"default_group_id,omitempty"`
	Locale               string                `json:"locale,omitempty"`
	NotificationsRead    string                `json:"notifications_read,omitempty"`
	NotificationsCount   int                   `json:"notifications_count,omitempty"`
	NotificationSettings *NotificationSettings `json:"notifications,omitempty"`
	Balance              []*balance            `json:"balance,omitempty"`
	// NotificationSettings     *NotificationSettings `json:"notification_settings,omitempty"`
}

func (u *User) String() string {
	return fmt.Sprintf("%# v", pretty.Formatter(u))
}

/*func (u *User) MarshalJSON() ([]byte, error) {
	// n := u.NotificationSettings
	return nil, nil
}*/

/*func (u *User) UnmarshalJSON(data []byte) error {
	// if err := json.Unmarshal(data, u); err != nil {
	// 	return err
	// }

	var m map[string]interface{}
	if err := json.Unmarshal(data, &m); err != nil {
		return err
	}

	if m["notification_settings"] != nil {
		u.NotificationSettings = m["notification_settings"].(*NotificationSettings)
	} else if m["notifications"] != nil {
		u.NotificationSettings = m["notifications"].(*NotificationSettings)
	}

	fmt.Printf("m : %+v\nu : %+v\n", m, u)
	return nil
}*/

// type UpdateUser struct {
// 	FirstName            string                `json:"first_name"`
// 	LastName             string                `json:"last_name"`
// 	Email                string                `json:"email,omitempty"`
// 	Password             string                `json:"password,omitempty"`
// 	DateFormat           string                `json:"date_format,omitempty"`
// 	DefaultCurrency      string                `json:"default_currency,omitempty"`
// 	DefaultGroupId       string                `json:"default_group_id,omitempty"`
// 	Locale               string                `json:"locale,omitempty"`
// 	NotificationSettings *NotificationSettings `json:"notification_settings,omitempty"`
// }

type UpdateUserResponse struct {
	User   *User         `json:"user"`
	Errors []interface{} `json:"errors"`
}

type NotificationSettings struct {
	AddedAsFriend  bool `json:"added_as_friend,omitempty"`
	AddedToGroup   bool `json:"added_to_group,omitempty"`
	ExpenseAdded   bool `json:"expense_added,omitempty"`
	ExpenseUpdated bool `json:"expense_updated,omitempty"`
	Bills          bool `json:"bills,omitempty"`
	Payments       bool `json:"payments,omitempty"`
	MonthlySummary bool `json:"monthly_summary,omitempty"`
	Announcements  bool `json:"announcements,omitempty"`
}

func (n *NotificationSettings) String() string {
	return fmt.Sprintf("%# v", pretty.Formatter(n))
	// return indentString(n)
	// return fmt.Sprintf("{ AddedAsFriend : %+v, AddedToGroup : %+v, ExpenseAdded : %+v, ExpenseUpdated : %+v, Bills : %+v, Payments : %+v, MonthlySummary : %+v, Announcements : %+v }\n")
}

type RegistrationStatus int

const (
	_ = iota
	RegistrationStatus_DUMMY
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

// GetCurrentUser allows a user to fetch their account information. All fields of the user's account are returned.
func (c *Client) GetCurrentUser(ctx context.Context) (*User, error) {
	url := c.addAccessTokenToUrl(GetCurrentUserUrl)
	r, err := c.Get(ctx, url, nil)
	if err != nil {
		return nil, err
	}

	var resp struct {
		User *User `json:"user"`
	}
	bts, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}
	fmt.Println("resp.body : ", string(bts))
	if err = json.Unmarshal(bts, &resp); err != nil {
		return nil, err
	}
	return resp.User, nil
}

// GetUser allows a user to fetch another user's information. For the non-active user, all fields will not be
// returned.
func (c *Client) GetUser(ctx context.Context, id int64) (*User, error) {
	url := c.addAccessTokenToUrl(fmt.Sprintf(GetUserByIdUrl, id))
	r, err := c.Get(ctx, url, nil)
	if err != nil {
		return nil, err
	}

	var resp struct {
		User *User `json:"user"`
	}
	bts, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}
	if err = json.Unmarshal(bts, &resp); err != nil {
		return nil, err
	}
	return resp.User, nil
}

// UpdateUser allows one to update information about their own account, and allows editing of the first_name, last_name,
// and email for acquaintances who have not logged in yet.
func (c *Client) UpdateUser(ctx context.Context, user *User) (*UpdateUserResponse, error) {
	url := c.addAccessTokenToUrl(fmt.Sprintf(UpdateUserByIdUrl, user.Id))

	type alias *User
	req := (alias)(user)
	// this way, we can achieve different json tags, one for the unmarshalling, and another for the marshalling
	req.NotificationSettings = user.NotificationSettings

	resp, err := c.Post(ctx, url, req, nil)
	if err != nil {
		return nil, err
	}
	bts, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var r UpdateUserResponse
	if err = json.Unmarshal(bts, &r); err != nil {
		return nil, err
	}
	return &r, nil
}

/*func (c *Client) GetUser(ctx context.Context, id int) (*User, error) {
	req, err := c.Get(fmt.Sprintf("/get_user/:%d", id), nil)

	var user *User

	_, err = c.do(ctx, req, &user)
	if err != nil {
		return nil, err
	}

	return user, err
}*/
