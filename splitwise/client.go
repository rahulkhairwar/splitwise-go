package splitwise

import (
	"context"
	"fmt"
	"golang.org/x/oauth2"
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
