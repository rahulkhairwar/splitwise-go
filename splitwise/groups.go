package splitwise

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"time"
)

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

// GetGroups fetches all the groups of the current user.
func (c *Client) GetGroups(ctx context.Context) ([]*Group, error) {
	url := c.addAccessTokenToUrl(GetGroupsUrl)
	r, err := c.Get(ctx, url, nil)
	if err != nil {
		return nil, err
	}

	var resp struct {
		Groups []*Group `json:"groups"`
	}
	bts, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}
	if err = json.Unmarshal(bts, &resp); err != nil {
		return nil, err
	}

	return resp.Groups, nil
}

// GetGroupById allows the current user to fetch one of his/her groups, by its id.
func (c * Client) GetGroupById(ctx context.Context, id int64) (*Group, error) {
	url := c.addAccessTokenToUrl(fmt.Sprintf(GetGroupByIdUrl, id))
	r, err := c.Get(ctx, url, nil)
	if err != nil {
		return nil, err
	}

	var resp struct {
		Group *Group `json:"group"`
	}
	bts, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}
	if err = json.Unmarshal(bts, &resp); err != nil {
		return nil, err
	}
	return resp.Group, nil
}

// CreateGroup allows the current user to create a new group.
func (c* Client) CreateGroup(ctx context.Context, g *Group) (*Group, error) {
	type r struct {
		*Group
		Errors []string `json:"errors"`
	}
	var resp struct {
		Group *r `json:"group"`
	}
	fmt.Println("resp : ", resp)
	return nil, nil
}

func (c* Client) DeleteGroupById(ctx context.Context, id int64) (*DeleteResponse, error) {
	return nil, nil
}

func (c* Client) UndeleteGroupById(ctx context.Context, id int64) (*DeleteResponse, error) {
	return nil, nil
}

func (c* Client) AddUserToGroup(ctx context.Context, groupId, userId int64, firstName, lastName, email string) (*DeleteResponse, error) {
	return nil, nil
}

func (c* Client) RemoveUserFromGroup(ctx context.Context, groupId, userId int64) (*DeleteResponse, error) {
	return nil, nil
}
