package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"

	jc "github.com/justmiles/jumpcloud-cli/lib"
)

var (
	command, system string
	timeout         int
)

// user list subcommand
var executeCMD = &cobra.Command{
	Use:   "execute",
	Short: "execute a command against a system or group of systems",
	Run: func(cmd *cobra.Command, args []string) {

		var exitCode = 0

		commandResults, err := jc.ExecuteCommandAgainstSystem(system, command, timeout)

		for _, commandResult := range commandResults {

			if commandResult.Response.Error != "" {
				fmt.Println(commandResult.System, commandResult.Response.Error)
			} else {
				fmt.Println(commandResult.System, commandResult.Response.Data.Output)
			}

			if commandResult.Response.Data.ExitCode > exitCode {
				exitCode = commandResult.Response.Data.ExitCode
			}
		}

		if err != nil {
			log.Fatal(err)
		}

		os.Exit(exitCode)
	},
}

func init() {

	// execute
	executeCMD.PersistentFlags().StringVarP(&command, "command", "c", "", "command to execute against the remote machine(s)")
	executeCMD.PersistentFlags().StringVarP(&system, "system", "s", "", "system to target with this command")
	executeCMD.PersistentFlags().IntVarP(&timeout, "timeout", "t", 0, "timeout after seconds")

	rootCmd.AddCommand(executeCMD)

}

// TODO: support more than one system
