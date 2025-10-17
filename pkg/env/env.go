package env

import (
	"log"

	env "github.com/caarlos0/env/v11"
	_ "github.com/joho/godotenv/autoload"
)

type Config struct {
	GOOGLE_SECRET  string `env:"GOOGLE_SECRET,notEmpty"`
	DISCORD_SECRET string `env:"DISCORD_SECRET,notEmpty"`
	GITHUB_SECRET  string `env:"GITHUB_SECRET,notEmpty"`
	SPOTIFY_SECRET string `env:"SPOTIFY_SECRET,notEmpty"`

	GOOGLE_KEY  string `env:"GOOGLE_KEY,notEmpty"`
	DISCORD_KEY string `env:"DISCORD_KEY,notEmpty"`
	GITHUB_KEY  string `env:"GITHUB_KEY,notEmpty"`
	SPOTIFY_KEY string `env:"SPOTIFY_KEY,notEmpty"`
}

var Cfg Config

func init() {
	log.Println("Parsing env vars")
	if err := env.Parse(&Cfg); err != nil {
		panic(err)
	}
}
