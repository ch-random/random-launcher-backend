package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/ch-random/random-launcher-backend/migration"
	"github.com/ch-random/random-launcher-backend/utils"
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

	for _, v := range migration.Versions {
		fmt.Println(utils.GetFuncName(v))
	}
	return
}
