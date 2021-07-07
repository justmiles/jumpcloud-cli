package cmd

import (
	"github.com/spf13/cobra"

	jc "github.com/justmiles/jumpcloud-cli/lib"
)

var systemCMD = &cobra.Command{
	Use:   "system",
	Short: "interact with JumpCloud systems",
	RunE:  nil,
}

// user list subcommand
var systemrListCmd = &cobra.Command{
	Use:   "list",
	Short: "list JumpCloud systems",
	Run: func(cmd *cobra.Command, args []string) {
		jc.ListSystems(query, output)
	},
}

func init() {

	systemCMD.AddCommand(systemrListCmd)
	rootCmd.AddCommand(systemCMD)

}
