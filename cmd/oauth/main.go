package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/divizn/go-lab/pkg/env"
	"github.com/gorilla/mux"
	"github.com/markbates/goth"
	"github.com/markbates/goth/gothic"
	"github.com/markbates/goth/providers/discord"
	"github.com/markbates/goth/providers/github"
	"github.com/markbates/goth/providers/google"
)

func main() {
	cfg := env.Cfg

	goth.UseProviders(
		google.New(cfg.GOOGLE_KEY, cfg.GOOGLE_SECRET, "http://localhost:3000/auth/google/callback"),
		github.New(cfg.GITHUB_KEY, cfg.GITHUB_SECRET, "http://localhost:3000/auth/github/callback"),
		discord.New(cfg.DISCORD_KEY, cfg.DISCORD_SECRET, "http://localhost:3000/auth/discord/callback", discord.ScopeIdentify, discord.ScopeEmail),
	)
	r := mux.NewRouter()

	r.HandleFunc("/auth/{provider}", func(w http.ResponseWriter, r *http.Request) {
		gothic.BeginAuthHandler(w, r)
	})

	r.HandleFunc("/auth/{provider}/callback", func(w http.ResponseWriter, r *http.Request) {
		user, err := gothic.CompleteUserAuth(w, r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		tmpl, err := template.New("user").Parse(userTemplate)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "text/html")
		tmpl.Execute(w, user)
	})

	r.HandleFunc("/logout", func(w http.ResponseWriter, r *http.Request) {
		gothic.Logout(w, r)
		fmt.Fprintln(w, "Logged out successfully.")
	})

	fmt.Println("Server running at http://localhost:3000")
	log.Fatal(http.ListenAndServe(":3000", r))
}

var userTemplate = `
<h2>User Info</h2>
<p><strong>Provider:</strong> {{.Provider}}</p>
<p><strong>Name:</strong> {{.Name}}</p>
<p><strong>Email:</strong> {{.Email}}</p>
<p><strong>UserID:</strong> {{.UserID}}</p>
<p><strong>IDToken</strong> {{.IDToken}}</p>
<p><strong>Avatar:</strong><br><img src="{{.AvatarURL}}" alt="Avatar" width="100" /></p>
<a href="/logout">Logout</a>`
