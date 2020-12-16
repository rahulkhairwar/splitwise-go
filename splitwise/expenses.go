package splitwise

import (
	"context"
	"time"
)

type Expense struct {
	Id                     int64                  `json:"id"`
	GroupId                int64                  `json:"group_id"`
	Description            string                 `json:"description"`
	Repeats                bool                   `json:"repeats"`
	RepeatInterval         RepeatInterval         `json:"repeat_interval"`
	EmailReminder          bool                   `json:"email_reminder"`
	EmailReminderInAdvance EmailReminderInAdvance `json:"email_reminder_in_advance"`
	Details                string                 `json:"details"`
	CommentsCount          int                    `json:"comments_count"`
	Payment                bool                   `json:"payment"`
	TransactionConfirmed   bool                   `json:"transaction_confirmed"`
	Cost                   string                 `json:"cost"`
	CurrencyCode           string                 `json:"currency_code"`
	Repayments             []*struct {
		From   int64  `json:"from"`
		To     int64  `json:"to"`
		Amount string `json:"amount"`
	} `json:"repayments"`
	Date      time.Time `json:"date"`
	CreatedAt time.Time `json:"created_at"`
	// CreatedBy
	UpdatedAt time.Time `json:"updated_at"`
	// 	UpdatedBy
	DeletedAt time.Time `json:"deleted_at"`
	// DeletedBy
	Category *struct {
		Id   int64  `json:"id"`
		Name string `json:"name"`
	} `json:"category"`
	Receipt *struct {
		Large    string `json:"large"`
		Original string `json:"original"`
	} `json:"receipt"`
	Users []*struct {
		User       User   `json:"user"`
		UserId     int64  `json:"user_id"`
		PaidShare  string `json:"paid_share"`
		OwedShare  string `json:"owed_share"`
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

func (c *Client) GetExpenses(ctx context.Context) ([]*Expense, error) {
	return nil, nil
}
