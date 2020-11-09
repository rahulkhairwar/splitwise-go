package splitwise

import (
	"encoding/json"
	"fmt"
	"github.com/kr/pretty"
)

type User struct {
	Id        int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Picture   *struct {
		Small  string `json:"small,omitempty"`
		Medium string `json:"medium,omitempty"`
		Large  string `json:"large,omitempty"`
	} `json:"picture,omitempty"`
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

type UpdateUserResponse struct {
	User *User `json:"user"`
	Errors []interface{} `json:"errors"`
}

func indentString(v interface{}) (string, error) {
	b, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		return "", err
	}
	return string(b), nil
}
