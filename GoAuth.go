package main

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
	"log"

	"github.com/gorilla/mux"
	"github.com/markbates/goth"
	"github.com/markbates/goth/gothic"
	"github.com/markbates/goth/providers/facebook"
	"github.com/markbates/goth/providers/twitter"
	"github.com/markbates/goth/providers/github"
	"github.com/joho/godotenv"
)

func main() {
	// Here we are loading in our .env file which will contain our Auth0 Client Secret and Domain
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	goth.UseProviders(
		github.New(os.Getenv("GITHUB_KEY"), os.Getenv("GITHUB_SECRET"), "http://localhost:3000/auth/github/callback"),
		twitter.New(os.Getenv("TWITTER_KEY"), os.Getenv("TWITTER_SECRET"), "http://localhost:3000/auth/twitter/callback"),
		facebook.New(os.Getenv("FACEBOOK_KEY"), os.Getenv("FACEBOOK_SECRET"), "http://localhost:3000/auth/facebook/callback"),
	)

	gothic.GetState = func(r *http.Request) string {
		return r.URL.Query().Get("state")
	}

	p := mux.NewRouter()
	p.HandleFunc("/auth/{provider}/callback", func(res http.ResponseWriter, req *http.Request) {
		user, err := gothic.CompleteUserAuth(res, req)
		if err != nil {
			fmt.Fprintln(res, err)
			return
		}
		t, _ := template.New("foo").Parse(userTemplate)
		t.Execute(res, user)
	})

	p.HandleFunc("/auth/{provider}", gothic.BeginAuthHandler)
	p.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		t, _ := template.New("foo").Parse(indexTemplate)
		t.Execute(res, nil)
	})
	http.ListenAndServe(":3000", p)
}

var indexTemplate = `
<p><a href="/auth/twitter">Log in with Twitter</a></p>
<p><a href="/auth/facebook">Log in with Facebook</a></p>
<p><a href="/auth/github">Log in with Github</a></p>
`

var userTemplate = `
<p>Name: {{.Name}}</p>
<p>Email: {{.Email}}</p>
<p>NickName: {{.NickName}}</p>
<p>Location: {{.Location}}</p>
<p>AvatarURL: {{.AvatarURL}} <img src="{{.AvatarURL}}"></p>
<p>Description: {{.Description}}</p>
<p>UserID: {{.UserID}}</p>
<p>AccessToken: {{.AccessToken}}</p>
`