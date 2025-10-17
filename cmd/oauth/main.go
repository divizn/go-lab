package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/divizn/go-lab/pkg/env"
	"github.com/gorilla/mux"
	"github.com/markbates/goth"
	"github.com/markbates/goth/gothic"
	"github.com/markbates/goth/providers/discord"
	"github.com/markbates/goth/providers/github"
	"github.com/markbates/goth/providers/google"
	"github.com/markbates/goth/providers/spotify"
)

func main() {
	cfg := env.Cfg

	goth.UseProviders(
		google.New(cfg.GOOGLE_KEY, cfg.GOOGLE_SECRET, "http://localhost:3000/auth/google/callback"),
		spotify.New(cfg.SPOTIFY_KEY, cfg.SPOTIFY_SECRET, "http://localhost:3000/auth/spotify/callback"),
		github.New(cfg.GITHUB_KEY, cfg.GITHUB_SECRET, "http://localhost:3000/auth/github/callback"),
		discord.New(cfg.DISCORD_KEY, cfg.DISCORD_SECRET, "http://localhost:3000/auth/discord/callback", discord.ScopeIdentify, discord.ScopeEmail, "openid"),
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

		fmt.Fprintf(w, "Login successful!\n\n")
		fmt.Fprintf(w, "Provider: %s\n", user.Provider)
		fmt.Fprintf(w, "UserID: %s\n", user.UserID)
		fmt.Fprintf(w, "Name: %s\n", user.Name)
		fmt.Fprintf(w, "Email: %s\n", user.Email)
		fmt.Fprintf(w, "AccessToken: %s\n", user.AccessToken)

		// TODO: fix openid connect for google
		if user.IDToken != "" {
			fmt.Fprintf(w, "IDToken: %s\n", user.IDToken)
		}
	})

	r.HandleFunc("/logout", func(w http.ResponseWriter, r *http.Request) {
		gothic.Logout(w, r)
		fmt.Fprintln(w, "Logged out successfully.")
	})

	fmt.Println("Server running at http://localhost:3000")
	log.Fatal(http.ListenAndServe(":3000", r))
}
