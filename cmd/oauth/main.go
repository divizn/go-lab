package main

import (
	"github.com/divizn/go-lab/pkg/env"
	"github.com/markbates/goth"
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
		discord.New(cfg.DISCORD_KEY, cfg.DISCORD_SECRET, "http://localhost:3000/auth/discord/callback", discord.ScopeIdentify, discord.ScopeEmail),
	)
}
