package cmd

import (
	"github.com/spf13/cobra"

	jc "github.com/justmiles/jumpcloud-cli/lib"
)

var (
	command, system string
)

// user list subcommand
var executeCMD = &cobra.Command{
	Use:   "execute",
	Short: "execute a command against a system or group of systems",
	Run: func(cmd *cobra.Command, args []string) {

		// Do the work (call the lib)
		jc.ExecuteCommandAgainstSystem(system, command)
		// lib.ExecuteCommandAgainstGroup(group,command)
	},
}

func init() {

	// execute
	executeCMD.PersistentFlags().StringVarP(&command, "command", "c", "", "command to execute against the remote machine(s)")
	executeCMD.PersistentFlags().StringVarP(&system, "system", "s", "", "system to target with this command")

	rootCmd.AddCommand(executeCMD)

}
