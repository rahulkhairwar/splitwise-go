package splitwise

import (
	"context"
	"encoding/json"
	"fmt"
	"golang.org/x/oauth2"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"time"
)

var alphabet = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ+-0123456789")

// Client struct stores configuration for Splitwise oAuth2.0 login, and also the access token, once it's been fetched
// after the user has granted access to their account.
type Client struct {
	*RestClient
	conf           *oauth2.Config
	state          string
	errRedirectUrl string
	logger         log.Logger
	accessToken    string
}

// New creates and returns a new Splitwise Client.
func New(consumerKey, secret, redirectUrl, errRedirectUrl string, httpClient http.Client) *Client {
	return &Client{
		RestClient: &RestClient{
			HttpClient: &httpClient,
		},
		conf: &oauth2.Config{
			RedirectURL:  redirectUrl,
			ClientID:     consumerKey,
			ClientSecret: secret,
			Endpoint: oauth2.Endpoint{
				AuthURL:  AuthorizeUrl,
				TokenURL: TokenUrl,
			},
		},
		errRedirectUrl: errRedirectUrl,
		logger:         log.Logger{},
	}
}

func (c *Client) GetState() string {
	return c.state
}

func (c* Client) SetAccessToken(a string) {
	c.accessToken = a
}

// HandleLogin is an HTTP handler to connect to Splitwise via oAuth2 flow, redirecting to the Splitwise authorization
// page, for the user to log in, and grant API access to their data. This handler also generates a random state,
// as required by the oAuth2.0 flow, and sends it along with the request.
func (c *Client) HandleLogin(w http.ResponseWriter, r *http.Request) {
	c.state = generateRandomState()
	url := c.conf.AuthCodeURL(c.state)
	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}

// HandleCallback is an HTTP handler to handle the callback from Splitwise, after authentication by the user. This
// handler can be replaced by your own implementation of a handler easily. Use the client.GetState() function to check
// the stored state, and store the returned access token using client.SetToken().
func (c *Client) HandleCallback(w http.ResponseWriter, r *http.Request) {
	state := r.FormValue("state")
	if state == "" || state != c.state {
		// todo : should (level-)log this, to the logger provided by the user.
		fmt.Println("Invalid state!")
		http.Redirect(w, r, c.errRedirectUrl, http.StatusTemporaryRedirect)
		return
	}

	token, err := c.conf.Exchange(context.Background(), r.FormValue("code"))
	if err != nil {
		fmt.Println("failed to get token due to : ", err)
		http.Redirect(w, r, c.errRedirectUrl, http.StatusTemporaryRedirect)
		return
	}
	c.accessToken = token.AccessToken
}

// GetCurrentUser API allows a user to fetch their account information. All fields of the user's account are returned.
func (c *Client) GetCurrentUser(ctx context.Context) (*User, error) {
	url := c.addAccessTokenToUrl(GetCurrentUserUrl)
	r, err := c.Get(ctx, url, nil)
	if err != nil {
		return nil, err
	}

	var resp struct {
		User *User `json:"user"`
	}
	bts, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}
	fmt.Println("resp.body : ", string(bts))
	if err = json.Unmarshal(bts, &resp); err != nil {
		return nil, err
	}
	return resp.User, err
}

// GetUser API allows a user to fetch another user's information. For the non-active user, all fields will not be
// returned.
func (c *Client) GetUser(ctx context.Context, id int64) (*User, error) {
	url := c.addAccessTokenToUrl(fmt.Sprintf(GetUserByIdUrl, id))
	r, err := c.Get(ctx, url, nil)
	if err != nil {
		return nil, err
	}

	var resp struct {
		User *User `json:"user"`
	}
	bts, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}
	if err = json.Unmarshal(bts, &resp); err != nil {
		return nil, err
	}
	return resp.User, err
}

// UpdateUser allows one to update information about their own account, and allows editing of the first_name, last_name,
// and email for acquaintances who have not logged in yet.
func (c *Client) UpdateUser(ctx context.Context, user *User) (*UpdateUserResponse, error) {
	url := c.addAccessTokenToUrl(fmt.Sprintf(UpdateUserByIdUrl, user.Id))

	type alias *User
	req := (alias)(user)
	req.NotificationSettings = user.NotificationSettings

	resp, err := c.Post(ctx, url, req, nil)
	if err != nil {
		return nil, err
	}
	bts, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var r UpdateUserResponse
	if err = json.Unmarshal(bts, &r); err != nil {
		return nil, err
	}
	return &r, nil
}

func (c *Client) GetGroups(ctx context.Context) ([]*Group, error) {
	return nil, nil
}

func (c * Client) GetGroupById(ctx context.Context, id int64) (*Group, error) {
	return nil, nil
}

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

func (c* Client) GetExpenses(ctx context.Context) ([]*Expense, error) {
	return nil, nil
}

func (c *Client) addAccessTokenToUrl(url string) string {
	return url + "?access_token=" + c.accessToken
}

func generateRandomState() string {
	rand.Seed(time.Now().UnixNano())
	var s string
	for i := 0; i < 32; i++ {
		s += string(alphabet[rand.Intn(len(alphabet))])
	}
	return s
}
