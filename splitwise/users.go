package splitwise

import (
	"errors"
	"fmt"
)

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

/*func (c *Client) GetUser(ctx context.Context, id int) (*User, error) {
	req, err := c.Get(fmt.Sprintf("/get_user/:%d", id), nil)

	var user *User

	_, err = c.do(ctx, req, &user)
	if err != nil {
		return nil, err
	}

	return user, err
}*/
