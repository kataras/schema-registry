package cmd

import (
	"sort"

	"github.com/kataras/bite"
	"github.com/spf13/cobra"
)

var versionsCmd = &cobra.Command{
	Use:     "versions",
	Short:   "lists all available versions",
	PreRunE: bite.ArgsRange(1, 1),
	RunE: func(_ *cobra.Command, args []string) error {
		subject := args[0]
		vers, err := client.Versions(subject)
		if err != nil {
			return err
		}

		sort.Ints(vers)

		return app.Print("%v\n", vers)
	},
}

func init() {
	app.AddCommand(versionsCmd)
}
