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

	GetCurrentUserUrl          = BaseApiUrl + "/get_current_user"
	GetUserByIdUrl             = BaseApiUrl + "/get_user/%d"
	UpdateUserByIdUrl          = BaseApiUrl + "/update_user/%d"
	GetGroupsUrl               = BaseApiUrl + "/get_groups"
	GetGroupByIdUrl            = BaseApiUrl + "/get_group/%d"
	CreateGroupUrl             = BaseApiUrl + "/create_group"
	DeleteGroupByIdUrl         = BaseApiUrl + "/delete_group/%d"
	UnDeleteGroupByIdUrl       = BaseApiUrl + "/undelete_group/%d"
	AddUserToGroupUrl          = BaseApiUrl + "/add_user_to_group"
	RemoveUserFromGroupUrl     = BaseApiUrl + "/remove_user_from_group"
	GetFriendsUrl              = BaseApiUrl + "/get_friends"
	GetFriendByIdUrl           = BaseApiUrl + "/get_friend/%d"
	CreateFriendUrl            = BaseApiUrl + "/create_friend"
	CreateFriendsUrl           = BaseApiUrl + "/create_friends"
	DeleteFriendByIdUrl        = BaseApiUrl + "/delete_friend/%d"
	GetExpenseByIdUrl          = BaseApiUrl + "/get_expense/%d"
	GetExpensesUrl             = BaseApiUrl + "/get_expenses"
	CreateExpenseUrl           = BaseApiUrl + "/create_expense"
	UpdateExpenseByIdUrl       = BaseApiUrl + "/update_expense/%d"
	DeleteExpenseByIdUrl       = BaseApiUrl + "/delete_expense/%d"
	UndeleteExpenseByIdUrl     = BaseApiUrl + "/undelete_expense/%d"
	GetCommentsForExpenseIdUrl = BaseApiUrl + "/get_comments?expense_id=%d"
	CreateCommentUrl = BaseApiUrl + "create_comment"

	GetCurrencyUrl = BaseApiUrl + "/get_currencies"
	GetCategoryUrl = BaseApiUrl + "/get_categories"
)
