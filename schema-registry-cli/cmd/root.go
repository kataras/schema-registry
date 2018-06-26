package cmd

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/landoop/schema-registry"

	"github.com/kataras/bite"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var client *schemaregistry.Client

var (
	verbose bool
	url     string
)

// App represents the base command when called without any subcommands
var app = &bite.Application{
	Name:                          "schema-registry-cli",
	Description:                   "A command line interface for the Confluent schema registry",
	Version:                       "0.0.2",
	ShowSpinner:                   false,
	DisableOutputFormatController: true,
	PersistentFlags: func(set *bite.Flags) {
		set.BoolVarP(&verbose, "verbose", "v", false, "be verbose")
		set.StringVarP(&url, "url", "e", schemaregistry.DefaultURL, "schema registry url, overrides SCHEMA_REGISTRY_URL")

		viper.SetEnvPrefix("schema_registry")
		viper.BindPFlag("url", set.Lookup("url"))
		viper.BindEnv("url")
	},
	Setup: func(cmd *cobra.Command, args []string) (err error) {
		if !verbose {
			log.SetOutput(ioutil.Discard)
		}

		url := viper.GetString("url")
		log.Printf("schema registry url: %s\n", url)

		client, err = schemaregistry.NewClient(url)
		return
	},
}

// Execute adds all child commands to the root command sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the app.
func Execute() {
	if err := app.Run(os.Stdout, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(-1)
	}
}
