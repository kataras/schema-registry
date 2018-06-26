package cmd

import (
	"log"
	"sort"

	"github.com/spf13/cobra"
)

var subjectsCmd = &cobra.Command{
	Use:   "subjects",
	Short: "lists all registered subjects",
	RunE: func(*cobra.Command, []string) error {
		subs, err := client.Subjects()
		if err != nil {
			return err
		}

		log.Printf("there are %d subjects\n", len(subs))
		sort.Strings(subs)
		for _, s := range subs {
			app.Print(s)
		}
		return nil
	},
}

func init() {
	app.AddCommand(subjectsCmd)
}
