package controllers

import (
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

var ssogolang *oauth2.Config

var RandomString = "random-string"

func init() {
	err := godotenv.Load("./.env")
	if err != nil {
		log.Fatalf("Error loading .env file: %s", err)
	}

	ssogolang = &oauth2.Config{
		RedirectURL:  os.Getenv("REDIRECT_URL"),
		ClientID:     os.Getenv("CLIENT_ID"),
		ClientSecret: os.Getenv("CLIENT_SECRET"),
		Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email"},
		Endpoint:     google.Endpoint,
	}

}

var Signin = SignInHandler{}

type SignInHandler struct{}

func (h SignInHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	url := ssogolang.AuthCodeURL(RandomString)
	log.Printf("Redirect URL: %s", url)
	http.Redirect(w, r, url, http.StatusSeeOther)
}
