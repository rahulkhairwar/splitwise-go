package splitwise

type User struct {
	Id        int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Picture   *struct {
		Small  string `json:"small,omitempty"`
		Medium string `json:"medium,omitempty"`
		Large  string `json:"large,omitempty"`
	} `json:"picture,omitempty"`
	Email                    string `json:"email,omitempty"`
	RegistrationStatus       string `json:"registration_status,omitempty"`
	DefaultCurrency          string `json:"default_currency,omitempty"`
	Locale                   string `json:"locale,omitempty"`
	NotificationsRead        string `json:"notifications_read,omitempty"`
	NotificationsCount       int    `json:"notifications_count,omitempty"`
	NotificationsPreferences *struct {
		AddedAsFriend bool `json:"added_as_friend,omitempty"`
	} `json:"notifications,omitempty"`
}
