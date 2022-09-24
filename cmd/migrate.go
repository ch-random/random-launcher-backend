package cmd

import (
	"github.com/spf13/cobra"
	"github.com/rs/zerolog/log"

	"github.com/ch-random/random-launcher-backend/migration"
	"github.com/ch-random/random-launcher-backend/repository/pscale"
)

var drops bool

func migrateCommand() *cobra.Command {
	cmd := cobra.Command{
		Use:     "migrate",
		Short:   "Execute database schema migration",
		PreRunE: preRunE,
		RunE:    migrateRunE,
	}
	flags := cmd.Flags()
	flags.BoolVar(&drops, "drop", false, "whether to truncate database (drop all tables)")
	return &cmd
}

func migrateRunE(cmd *cobra.Command, args []string) (err error) {
	if len(args) > 0 {
		return errArgsProvided
	}

	db, err := pscale.GetDB()
	if err != nil {
		return
	}

	// drop all tables
	if drops {
		if err = migration.DropAllTables(db); err != nil {
			log.Warn().Err(err)
			return
		}
	}

	// migrate
	_, err = migration.Migrate(db)
	return
}
