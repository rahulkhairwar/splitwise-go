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

type Client struct {
	*RestClient
	conf           *oauth2.Config
	state          string
	errRedirectUrl string
	logger         log.Logger
	accessToken    string
}

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

func (c *Client) HandleLogin(w http.ResponseWriter, r *http.Request) {
	c.state = generateRandomState()
	url := c.conf.AuthCodeURL(c.state)
	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}

func (c *Client) HandleCallback(w http.ResponseWriter, r *http.Request) {
	state := r.FormValue("state")
	if state == "" || state != c.state {
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
	if err = json.Unmarshal(bts, &resp); err != nil {
		return nil, err
	}
	return resp.User, err
}

func (c* Client) GetUser(ctx context.Context, id int64) (*User, error) {
	url := c.addAccessTokenToUrl(fmt.Sprintf(GetUserUrl, id))
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
