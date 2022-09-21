package cmd

import (
	"github.com/spf13/cobra"

	"github.com/ch-random/random-launcher-backend/repository/pscale"
	"github.com/ch-random/random-launcher-backend/migration"
)

func migrateCommand() *cobra.Command {
	var dropsDB bool

	cmd := cobra.Command{
		Use:   "migrate",
		Short: "Execute database schema migration",
		RunE: func(cmd *cobra.Command, args []string) error {
			engine, err := pscale.GetDB()
			if err != nil {
				return err
			}
			db, err := engine.DB()
			if err != nil {
				return err
			}
			defer db.Close()
			if dropsDB {
				if err := migration.DropAll(engine); err != nil {
					return err
				}
			}
			_, err = migration.Migrate(engine)
			return err
		},
	}

	flags := cmd.Flags()
	flags.BoolVar(&dropsDB, "reset", false, "whether to truncate database (drop all tables)")

	return &cmd
}
