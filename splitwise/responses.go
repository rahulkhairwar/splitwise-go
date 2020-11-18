package splitwise

import (
	"encoding/json"
	"fmt"
	"github.com/kr/pretty"
	"time"
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
	*userCommon
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
	Balance              []*Balance            `json:"balance,omitempty"`
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

type Balance struct {
	CurrencyCode string `json:"currency_code"`
	Amount       string `json:"amount"`
}

type UpdateUserResponse struct {
	User   *User         `json:"user"`
	Errors []interface{} `json:"errors"`
}

type Group struct {
	Id                int64     `json:"id"`
	Name              string    `json:"name"`
	UpdatedAt         time.Time `json:"updated_at"`
	Members           []*User   `json:"members"`
	SimplifyByDefault bool      `json:"simplify_by_default"`
	OriginalDebts     []*Debt   `json:"original_debts"`
	SimplifiedDebts   []*Debt   `json:"simplified_debts"`
	Whiteboard        string    `json:"whiteboard"`
	GroupType         string    `json:"group_type"`
	InviteLink        string    `json:"invite_link"`
}

type Debt struct {
	From         int    `json:"from"`
	To           int    `json:"to"`
	Amount       string `json:"amount"`
	CurrencyCode string `json:"currency_code"`
}

type DeleteResponse struct {
	Success bool     `json:"success"`
	Errors  []string `json:"errors"`
}

type Friend struct {
	*userCommon
	Balance []*Balance `json:"balance"`
	Groups  []*struct {
		GroupId int64      `json:"group_id"`
		Balance []*Balance `json:"balance"`
	} `json:"groups"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Expense struct {
	Id                     int64                  `json:"id"`
	GroupId                int64                  `json:"group_id"`
	Description            string                 `json:"description"`
	Repeats                bool                   `json:"repeats"`
	RepeatInterval         RepeatInterval         `json:"repeat_interval"`
	EmailReminder          bool                   `json:"email_reminder"`
	EmailReminderInAdvance EmailReminderInAdvance `json:"email_reminder_in_advance"`
	Details string `json:"details"`
	CommentsCount int `json:"comments_count"`
	Payment bool `json:"payment"`
	TransactionConfirmed bool `json:"transaction_confirmed"`
	Cost string `json:"cost"`
	CurrencyCode string `json:"currency_code"`
	Repayments []*struct {
		From int64 `json:"from"`
		To int64 `json:"to"`
		Amount string `json:"amount"`
	} `json:"repayments"`
	Date time.Time `json:"date"`
	CreatedAt time.Time `json:"created_at"`
	// CreatedBy
	UpdatedAt time.Time `json:"updated_at"`
// 	UpdatedBy
	DeletedAt time.Time `json:"deleted_at"`
	// DeletedBy
	Category *struct {
		Id int64 `json:"id"`
		Name string `json:"name"`
	} `json:"category"`
	Receipt *struct {
		Large string `json:"large"`
		Original string `json:"original"`
	} `json:"receipt"`
	Users []*struct {
		User User `json:"user"`
		UserId int64 `json:"user_id"`
		PaidShare string `json:"paid_share"`
		OwedShare string `json:"owed_share"`
		NetBalance string `json:"net_balance"`
	} `json:"users"`
}

type RepeatInterval string

const (
	Never       RepeatInterval = "never"
	Weekly      RepeatInterval = "weekly"
	Fortnightly RepeatInterval = "fortnightly"
	Monthly     RepeatInterval = "monthly"
	Yearly      RepeatInterval = "yearly"
)

type EmailReminderInAdvance int

const (
	A EmailReminderInAdvance = -1
	B EmailReminderInAdvance = 0
	C EmailReminderInAdvance = 1
	D EmailReminderInAdvance = 3
	E EmailReminderInAdvance = 5
	F EmailReminderInAdvance = 7
	G EmailReminderInAdvance = 14
)

func indentString(v interface{}) (string, error) {
	b, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		return "", err
	}
	return string(b), nil
}
