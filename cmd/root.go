// $ go run cmd/app/main.go
package cmd

import (
	"context"

	firebase "firebase.google.com/go"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"

	"github.com/ch-random/random-launcher-backend/config"
	"github.com/ch-random/random-launcher-backend/delivery/httpserver"
)

func rootCommand() *cobra.Command {
	var cmd = &cobra.Command{
		Use:     appName,
		Short:   short,
		Long:    long,
		Version: version,
		PreRunE: preRunE,
		RunE:    rootRunE,
		// PostRunE,
		// PersistentPostRunE,
	}
	return cmd
}

func rootRunE(cmd *cobra.Command, args []string) (err error) {
	if len(args) > 0 {
		return errArgsProvided
	}

	_, err = firebase.NewApp(context.Background(), nil)
	if err != nil {
		log.Warn().Err(err).Msg("failed to connect to Firebase")
	}

	e := httpserver.NewHandler()

	port := getEnvOrDefault("PORT", config.Port)
	log.Info().Msgf("listening on %s", port)
	addr := ":" + port
	if err := e.Start(addr); err != nil {
		log.Warn().Err(err).Msg("failed to serve HTTP")
	}
	return
}
