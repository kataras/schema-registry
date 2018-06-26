package cmd

import (
	"fmt"
	"log"

	"github.com/kataras/bite"
	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:     "add <subject>",
	Short:   "registers the schema provided through stdin",
	PreRunE: bite.ArgsRange(1, 1),
	RunE: func(cmd *cobra.Command, args []string) error {
		subject := args[0]
		has, in, err := bite.ReadInPipe()
		if err != nil {
			return err
		}

		if !has {
			return fmt.Errorf("expected avro schema from input pipe")
		}

		id, err := client.RegisterNewSchema(subject, string(in))
		if err != nil {
			return err
		}

		log.Printf("registered schema with id %d\n", id)
		return nil
	},
}

func init() {
	app.AddCommand(addCmd)
}
