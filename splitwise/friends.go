package splitwise

import (
	"context"
	"time"
)

type Friend struct {
	*userCommon
	Balance []*balance `json:"balance"`
	Groups  []*struct {
		GroupId int64      `json:"group_id"`
		Balance []*balance `json:"balance"`
	} `json:"groups"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (c* Client) GetFriends(ctx context.Context) ([]*Friend, error) {
	return nil, nil
}

func (c* Client) GetFriendById(ctx context.Context, id int64) (*Friend, error) {
	return nil, nil
}

func (c* Client) CreateFriend(ctx context.Context, userEmail, userFirstName, userLastName string) (*Friend, error) {
	return nil, nil
}

func (c* Client) CreateFriends(ctx context.Context, userEmails, userFirstNames, userLastNames []string) ([]*Friend, error) {
	return nil, nil
}

func (c* Client) DeleteFriendById(ctx context.Context, id int64) (*DeleteResponse, error) {
	return nil, nil
}
