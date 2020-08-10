package cmd

import (
	jc "github.com/justmiles/jumpcloud-cli/lib"
	"github.com/spf13/cobra"
)

var groupCmd = &cobra.Command{
	Use:   "group",
	Short: "interact with JumpCloud groups",
	RunE:  nil,
}

// user list subcommand
var groupListCmd = &cobra.Command{
	Use:   "list",
	Short: "list jumpcloud groups",
	Run: func(cmd *cobra.Command, args []string) {
		jc.ListGroups(query, output)
	},
}

// user list subcommand
var (
	groupID                      string
	graphUserGroupMembersListCMD = &cobra.Command{
		Use:   "list-members",
		Short: "list mebers of group",
		Run: func(cmd *cobra.Command, args []string) {
			jc.GraphUserGroupMembersList(groupID, query, output)
		},
	}
)

func init() {

	// list-members
	graphUserGroupMembersListCMD.PersistentFlags().StringVarP(&groupID, "id", "i", groupID, "group ID to list")

	groupCmd.AddCommand(groupListCmd, graphUserGroupMembersListCMD)
	rootCmd.AddCommand(groupCmd)

}
