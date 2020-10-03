package splitwise

import "fmt"

var (
	version3_0   = "v3.0"
	splitwiseUrl = "https://secure.splitwise.com"

	BaseApiUrl      = fmt.Sprintf("%s/api/%s", splitwiseUrl, version3_0)
	RequestTokenUrl = BaseApiUrl + "/get_request_token"
	AccessTokenUrl  = BaseApiUrl + "/get_access_token"
	AuthorizeUrl    = splitwiseUrl + "/oauth/authorize"
	TokenUrl        = splitwiseUrl + "/oauth/token"

	GetCurrentUserUrl = BaseApiUrl + "/get_current_user"
	GetUserUrl        = BaseApiUrl + "/get_user/%d"
	GetFriendsUrl     = BaseApiUrl + "/get_friends"
	GetGroupsUrl      = BaseApiUrl + "/get_groups"
	GetGroupUrl       = BaseApiUrl + "/get_group"
	GetCurrencyUrl    = BaseApiUrl + "/get_currencies"
	GetCategoryUrl    = BaseApiUrl + "/get_categories"
	GetExpensesUrl    = BaseApiUrl + "/get_expenses"
	GetExpenseUrl     = BaseApiUrl + "/get_expense"
	CreateExpenseUrl  = BaseApiUrl + "/create_expense"
	CreateGroupUrl    = BaseApiUrl + "/create_group"
)
