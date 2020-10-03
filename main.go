package main

import (
	"context"
	"fmt"
	"github.com/rahulkhairwar/splitwise-go/splitwise"
	"golang.org/x/oauth2"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"os"
	"time"

	_ "github.com/joho/godotenv/autoload"
)

var (
	consumerKey    = os.Getenv("CONSUMER_KEY")
	consumerSecret = os.Getenv("CONSUMER_SECRET")
	conf           = oauth2.Config{
		RedirectURL:  "http://localhost:8080/callback",
		ClientID:     consumerKey,
		ClientSecret: consumerSecret,
		Endpoint: oauth2.Endpoint{
			AuthURL:  "https://secure.splitwise.com/oauth/authorize",
			TokenURL: "https://secure.splitwise.com/oauth/token",
		},
	}
)

func handleHome(w http.ResponseWriter, r *http.Request) {
	var html = `<html><body><a href="/login">Splitwise Login</a></body></html>`
	_, err := fmt.Fprint(w, html)
	if err != nil {
		log.Fatalln("failed to write to response writer due to : ", err)
	}
}

var state = ""

func handleLogin(w http.ResponseWriter, r *http.Request) {
	state = generateRandomState()
	url := conf.AuthCodeURL(state)
	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}

var alphabet = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ+-0123456789")

func generateRandomState() string {
	rand.Seed(time.Now().UnixNano())
	s := ""
	for i := 0; i < 32; i++ {
		s += string(alphabet[rand.Intn(len(alphabet))])
	}
	fmt.Println("generated state, s : ", s)
	return s
}

func handleCallback(w http.ResponseWriter, r *http.Request) {
	if r.FormValue("state") != state {
		fmt.Println("Invalid state!")
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}

	token, err := conf.Exchange(context.Background(), r.FormValue("code"))
	if err != nil {
		fmt.Println("failed to get token due to : ", err)
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}

	resp, err := http.Get("https://secure.splitwise.com/api/v3.0/get_current_user?access_token=" + token.AccessToken)
	if err != nil {
		fmt.Println("failed to GET current user from Splitwise due to : ", err)
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}
	defer resp.Body.Close()
	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("failed to read all from response body due to : ", err)
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}
	fmt.Println("current_user : ", string(bytes))
}

func main() {
	a := splitwise.New(consumerKey, consumerSecret, "http://localhost:8080/callback", "/", http.Client{})
	http.HandleFunc("/", handleHome)
	http.HandleFunc("/login", a.HandleLogin)
	http.HandleFunc("/callback", a.HandleCallback)
	http.ListenAndServe(":8080", nil)

/*	http.HandleFunc("/", handleHome)
	http.HandleFunc("/login", handleLogin)
	http.HandleFunc("/callback", handleCallback)
	http.ListenAndServe(":8080", nil)*/
}

func main3() {
	ctx := context.Background()
	consumerKey := os.Getenv("CONSUMER_KEY")
	consumerSecret := os.Getenv("CONSUMER_SECRET")
	conf := oauth2.Config{
		ClientID:     consumerKey,
		ClientSecret: consumerSecret,
		Endpoint: oauth2.Endpoint{
			AuthURL:  "https://secure.splitwise.com/oauth/authorize",
			TokenURL: "https://secure.splitwise.com/oauth/token",
		},
		RedirectURL: "http://localhost:8080/callback",
	}

	conf.AuthCodeURL("testState")

	tok, err := conf.Exchange(ctx, "testCode")
	if err != nil {
		log.Fatalf("failed to exchange due to : %+v\n", err)
	}
	log.Printf("tok : %+v\n", tok)
}

func main2() {
	// ctx := context.Background()
	// consumerKey := os.Getenv("CONSUMER_KEY")
	// consumerSecret := os.Getenv("CONSUMER_SECRET")
	// client := splitwise.New(consumerKey, consumerSecret, http.Client{})
	// user, err := client.GetCurrentUser(ctx)
	// if err != nil {
	// 	log.Fatalf("failed to get current user due to : %+v\n", err)
	// }
	// log.Printf("current user : %+v\n", user)

	// splitwise.Do(consumerKey, consumerSecret)
	// fmt.Printf("httpClient : %+v\n", httpClient)
}
