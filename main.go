package main

import (
	"go-api-test/api"
	"go-api-test/util"
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal().Err(err).Msg("cannot load config")
	}

	if config.Environment == "development" {
		log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	}

	server, err := api.NewServer(config)
	if err != nil {
		log.Fatal().Err(err).Msg("cannot create server")
	}

	err = server.Start(":8080")
	if err != nil {
		log.Fatal().Err(err).Msg("cannot start server")
	}
}
