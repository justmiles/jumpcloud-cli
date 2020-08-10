package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
	attributes                                   []string
	query, attributeName, attributeValue, output string
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "jc",
	Short: "cli to interact with JumpCloud",
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute(version string) {
	rootCmd.Version = version
	rootCmd.SetVersionTemplate(`{{printf "%s" .Version}}
`)
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&query, "query", "q", query, "list of attributes to be added to the report")
	rootCmd.PersistentFlags().StringVarP(&output, "output", "o", "json", "The formatting style for command output: table, csv, or json")
}
