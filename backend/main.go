package main

import (
	"chess/server"
	"fmt"

	"github.com/caarlos0/env"
	"github.com/carlosarismendi/dddhelper/shared/infrastructure/dotenv"
)

type Config struct {
	Port               string `env:"PORT" envDefault:"8081"`
	ServeSinglePageApp bool   `env:"SERVE_SINGLE_PAGE_APP" envDefault:"false"`
}

func newConfig() *Config {
	cfg := Config{}
	if err := env.Parse(&cfg); err != nil {
		panic(err)
	}
	return &cfg
}

func main() {
	dotenv.Load()

	cfg := newConfig()
	s := server.NewServer("", fmt.Sprintf(":%s", cfg.Port))
	s.Static("/app", "./dist", cfg.ServeSinglePageApp)
	s.Run()
}
