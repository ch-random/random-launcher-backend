package cmd

import (
	"errors"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

const (
	appName = "random-launcher-backend"
	version = "v0.0.1"
	about   = "backend for random-launcher"
)

var (
	short = fmt.Sprintf("%s %s", appName, version)
	long  = fmt.Sprintf("%s\n%s", short, about)
	errArgsProvided = errors.New("you should not provide any arguments")
)

func preRunE(cmd *cobra.Command, _args []string) (err error) {
	log.Info().Msg(short)

	// Load environment variables from `.env`
	if err := godotenv.Load(); err != nil {
		log.Warn().Err(err).Msg("failed to load environment variables")
	}
	return
}

func Execute() (err error) {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	cobra.OnInitialize()
	rootCmd := rootCommand()
	rootCmd.AddCommand(
		listCommand(),
		migrateCommand(),
	)
	if err = rootCmd.Execute(); err != nil {
		log.Fatal().Msg(err.Error())
		return
	}
	return
}
