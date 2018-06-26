package cmd

import (
	"fmt"

	"github.com/kataras/bite"
	"github.com/spf13/cobra"
)

var existsCmd = &cobra.Command{
	Use:     "exists <subject>",
	Short:   "checks if the schema provided through stdin exists for the subject",
	PreRunE: bite.ArgsRange(1, 1),
	RunE: func(cmd *cobra.Command, args []string) error {
		subject := args[0]
		has, in, err := bite.ReadInPipe()
		if err != nil {
			return err
		}

		if !has {
			return fmt.Errorf("expected schema from input pipe")
		}

		isreg, sch, err := client.IsRegistered(subject, string(in))
		if err != nil {
			return err
		}
		app.Print("exists: %v\n", isreg)
		if isreg {
			app.Print("id: %d\n", sch.ID)
			app.Print("version: %d\n", sch.Version)
		}
		return nil
	},
}

func init() {
	app.AddCommand(existsCmd)
}
