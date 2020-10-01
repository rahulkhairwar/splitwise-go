package splitwise

import (
	"context"
	"golang.org/x/oauth2"
	"log"
	"net/http"
	"reflect"
	"testing"
)

func TestClient_GetCurrentUser(t *testing.T) {
	type fields struct {
		RestClient *RestClient
		conf       *oauth2.Config
		state      string
		logger     log.Logger
		token      *oauth2.Token
	}
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *User
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Client{
				RestClient: tt.fields.RestClient,
				conf:       tt.fields.conf,
				state:      tt.fields.state,
				logger:     tt.fields.logger,
				token:      tt.fields.token,
			}
			got, err := c.GetCurrentUser(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetCurrentUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetCurrentUser() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_accessTokenToUrl(t *testing.T) {
	type fields struct {
		RestClient *RestClient
		conf       *oauth2.Config
		state      string
		logger     log.Logger
		token      *oauth2.Token
	}
	type args struct {
		url string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Client{
				RestClient: tt.fields.RestClient,
				conf:       tt.fields.conf,
				state:      tt.fields.state,
				logger:     tt.fields.logger,
				token:      tt.fields.token,
			}
			if got := c.accessTokenToUrl(tt.args.url); got != tt.want {
				t.Errorf("accessTokenToUrl() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_handleCallback(t *testing.T) {
	type fields struct {
		RestClient *RestClient
		conf       *oauth2.Config
		state      string
		logger     log.Logger
		token      *oauth2.Token
	}
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Client{
				RestClient: tt.fields.RestClient,
				conf:       tt.fields.conf,
				state:      tt.fields.state,
				logger:     tt.fields.logger,
				token:      tt.fields.token,
			}

			_ = c.conf
		})
	}
}

func TestClient_handleLogin(t *testing.T) {
	type fields struct {
		RestClient *RestClient
		conf       *oauth2.Config
		state      string
		logger     log.Logger
		token      *oauth2.Token
	}
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Client{
				RestClient: tt.fields.RestClient,
				conf:       tt.fields.conf,
				state:      tt.fields.state,
				logger:     tt.fields.logger,
				token:      tt.fields.token,
			}

			_ = c.conf
		})
	}
}

func TestNew(t *testing.T) {
	type args struct {
		consumerKey string
		secret      string
		redirectUrl string
		httpClient  http.Client
	}
	tests := []struct {
		name string
		args args
		want *Client
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := New(tt.args.consumerKey, tt.args.secret, tt.args.redirectUrl, tt.args.httpClient); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_generateRandomState(t *testing.T) {
	// tests := []struct {
	// 	name string
	// 	want string
	// }{{}, {}}
	for i := 0; i < 5; i++ {
		t.Run(string(i), func(t *testing.T) {
			if got := generateRandomState(); len(got) != 32 {
				t.Errorf("generateRandomState() = %v", got)
			}
		})
	}
}