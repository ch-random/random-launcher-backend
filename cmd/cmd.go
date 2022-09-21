package cmd

import (
	"log"
	"os"

	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
    Use:   "launcher",
    Short: "game launcher",
    Run: func(cmd *cobra.Command, args []string) {
        log.Println("root command")
    },
}

func Init() (err error) {
    cobra.OnInitialize()
    RootCmd.AddCommand(
        migrateCommand(),
    )
	if err = RootCmd.Execute(); err != nil {
		log.Fatalln(os.Args[0], err)
	}
	return
}
