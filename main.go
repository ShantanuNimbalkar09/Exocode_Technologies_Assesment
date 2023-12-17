package main

import (
	"app/controllers"
	"html/template"
	"log"
	"net/http"
)

const (
	port         = ":3000"
	staticFolder = "public"
)

func main() {
	fs := http.FileServer(http.Dir(staticFolder))

	http.Handle("/", fs)
	http.HandleFunc("/login", loginHandler)
	http.Handle("/signin", controllers.Signin)
	http.Handle("/callback", controllers.Callback)
	http.Handle("/login/github", controllers.GithubSignin)
	http.Handle("/login/github/callback", controllers.GithubCallback)

	log.Fatal(http.ListenAndServe(port, nil))

}

func loginHandler(w http.ResponseWriter, r *http.Request) {

	tmpl, err := template.ParseFiles("index.html")
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, nil)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}
