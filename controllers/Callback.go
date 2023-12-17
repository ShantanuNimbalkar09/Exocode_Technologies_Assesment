package controllers

import (
	"context"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

const googleAPIEndpoint = "https://www.googleapis.com/oauth2/v2/userinfo?access_token="

type CallbackHandler struct{}

var Callback = CallbackHandler{}

func (h CallbackHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	state := r.FormValue("state")
	code := r.FormValue("code")
	data, err := getUserData(state, code)
	if err != nil {
		handleError(w, err, "Error getting user data", http.StatusInternalServerError)
		return
	}
	fmt.Fprintf(w, "Data : %s", data)
}

func handleError(w http.ResponseWriter, err error, message string, statusCode int) {
	log.Println("Error:", err)
	http.Error(w, message, statusCode)
}

func getUserData(state, code string) ([]byte, error) {
	if state != RandomString {
		return nil, errors.New("invalid user state")
	}

	token, err := ssogolang.Exchange(context.Background(), code)
	if err != nil {
		return nil, fmt.Errorf("error exchanging code for token: %w", err)

	}

	response, err := http.Get(googleAPIEndpoint + token.AccessToken)
	if err != nil {
		return nil, fmt.Errorf("error getting user data from Google API: %w", err)

	}

	defer response.Body.Close()

	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %w", err)

	}

	return data, nil
}
