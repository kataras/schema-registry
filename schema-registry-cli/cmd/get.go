package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"strconv"

	"github.com/landoop/schema-registry"

	"github.com/kataras/bite"
	"github.com/spf13/cobra"
)

// get can handle three argument styles: <id>, <subj ver> or <subj>
var getCmd = &cobra.Command{
	Use:   "get <id> | (<subject> [<version>])",
	Short: "retrieves a schema specified by id or subject",
	Long: `The schema can be requested by id or subject.
When a subject is given, optionally one can provide a specific version. If no
version is specified, the latest version is returned.
`,
	PreRunE: bite.ArgsRange(1, 2),
	RunE: func(cmd *cobra.Command, args []string) error {
		subject := args[0]
		id, idParseErr := strconv.Atoi(subject)

		if len(args) == 1 { // it's either the ID or the subject name.

			if idParseErr != nil { // it's string, the subject.
				sch, err := client.GetLatestSchema(subject)
				if err != nil {
					return err
				}
				return printSchema(sch)
			}

			schRaw, err := client.GetSchemaByID(id)
			if err != nil {
				return err
			}

			return app.Print(schRaw)
		}

		// args contain the subject name and the versio number.
		ver, err := strconv.Atoi(args[1])
		if err != nil {
			return fmt.Errorf("2nd argument must be a version number")
		}

		sch, err := client.GetSchemaBySubject(subject, ver)
		if err != nil {
			return err
		}

		return printSchema(sch)
	},
}

func init() {
	app.AddCommand(getCmd)
}

func printSchema(sch schemaregistry.Schema) error {
	log.Printf("version: %d\n", sch.Version)
	log.Printf("id: %d\n", sch.ID)
	var indented bytes.Buffer
	if err := json.Indent(&indented, []byte(sch.Schema), "", "  "); err != nil {
		fmt.Println(sch.Schema) //isn't a json object, which is legal
		return err
	}
	return app.Print(indented.String())
}
