package main

import (
	"context"
	"fmt"
	"golang.org/x/oauth2"
)

const (
	BaseUrl = "https://www.splitwise.com"
	VERSION = "v3.0"

	RequestTokenUrl   = BaseUrl + "/api/" + VERSION + "/get_request_token"
	AccessTokenUrl    = BaseUrl + "/api/" + VERSION + "/get_access_token"
	AuthorizeUrl      = BaseUrl + "/authorize"
	GetCurrentUserUrl = BaseUrl + "/api/" + VERSION + "/get_current_user"
	GetUserUrl        = BaseUrl + "/api/" + VERSION + "/get_user"
	GetFriendsUrl     = BaseUrl + "/api/" + VERSION + "/get_friends"
	GetGroupsUrl      = BaseUrl + "/api/" + VERSION + "/get_groups"
	GetGroupUrl       = BaseUrl + "/api/" + VERSION + "/get_group"
	GetCurrencyUrl    = BaseUrl + "/api/" + VERSION + "/get_currencies"
	GetCategoryUrl    = BaseUrl + "/api/" + VERSION + "/get_categories"
	GetExpensesUrl    = BaseUrl + "/api/" + VERSION + "/get_expenses"
	GetExpenseUrl     = BaseUrl + "/api/" + VERSION + "/get_expense"
	CreateExpenseUrl  = BaseUrl + "/api/" + VERSION + "/create_expense"
	CreateGroupUrl    = BaseUrl + "/api/" + VERSION + "/create_group"
)

func main() {
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: ""})
	tc := oauth2.NewClient(ctx, ts)

	fmt.Printf("tc : %+v\n", tc)
}
