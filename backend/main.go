package main

import (
	"chess/server"
	"fmt"

	"github.com/caarlos0/env"
	"github.com/carlosarismendi/dddhelper/shared/infrastructure/dotenv"
)

type Config struct {
	Port                string `env:"PORT" envDefault:"8081"`
	ServeSinglePageApp  bool   `env:"SERVE_SINGLE_PAGE_APP" envDefault:"false"`
	SinglePageAppFolder string `env:"SINGLE_PAGE_APP_FOLDER" envDefault:"./dist"`
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
	s.Static("/app", cfg.SinglePageAppFolder, cfg.ServeSinglePageApp)
	s.Run()
}
