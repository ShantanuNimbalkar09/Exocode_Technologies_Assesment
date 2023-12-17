package controllers

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

var GithubSignin = GithubSignInHandler{}

type GithubSignInHandler struct{}

const githubRedirectURI = "http://localhost:3000/login/github/callback"

func (h GithubSignInHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	githubClientID := getGithubClientID()

	redirectURL := fmt.Sprintf("https://github.com/login/oauth/authorize?client_id=%s&redirect_uri=%s", githubClientID, githubRedirectURI)

	http.Redirect(w, r, redirectURL, 301)

}

func getGithubClientID() string {

	githubClientID, exists := os.LookupEnv("GITHUB_CLIENT_ID")
	if !exists {
		log.Fatalf("Github Client ID not defined in .env file")
	}

	return githubClientID
}
