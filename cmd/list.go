package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/ch-random/random-launcher-backend/migration"
)

func listCommand() *cobra.Command {
	var cmd = &cobra.Command{
		Use:     "list",
		Short:   "List all versions available",
		Version: version,
		PreRunE: preRunE,
		RunE:    listRunE,
	}
	return cmd
}

func listRunE(cmd *cobra.Command, args []string) (err error) {
	if len(args) > 0 {
		return errArgsProvided
	}

	for _, v := range migration.GetVersions() {
		fmt.Println(v)
	}
	return
}
