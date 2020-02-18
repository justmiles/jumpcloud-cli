package cmd

import (
	"fmt"
	"log"
	"os"

	jc "github.com/justmiles/jumpcloud-cli/lib"
	"github.com/spf13/cobra"
)

var (
	userName, attributeName, attributeValue string
)
var userCmd = &cobra.Command{
	Use:   "user",
	Short: "interact with JumpCloud users",
	RunE:  nil,
}

// user list subcommand
var userListCmd = &cobra.Command{
	Use:   "list",
	Short: "list jumpcloud users",
	Run: func(cmd *cobra.Command, args []string) {
		jc.ListUsers()
	},
}

var userAttributesCmd = &cobra.Command{
	Use:   "attributes",
	Short: "show attributes for a user",
	Run: func(cmd *cobra.Command, args []string) {
		err := jc.UserAttributes(userName)
		if err != nil {
			log.Fatal(err)
		}
	},
}

var attributeMatchesCmd = &cobra.Command{
	Use:   "attribute-matches",
	Short: "exits successfully if a user's attribute key/value pair matches",
	Run: func(cmd *cobra.Command, args []string) {
		matches, err := jc.UserAttributeMatches(userName, attributeName, attributeValue)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(matches)
		if !matches {
			os.Exit(1)
		}
	},
}

func init() {

	// user attributes flags
	userAttributesCmd.PersistentFlags().StringVarP(&userName, "username", "u", "", "username")
	cobra.MarkFlagRequired(userAttributesCmd.PersistentFlags(), "username")

	// user attribute matches flags
	attributeMatchesCmd.PersistentFlags().StringVarP(&userName, "username", "u", "", "username")
	attributeMatchesCmd.PersistentFlags().StringVarP(&attributeName, "key", "k", "", "attribute name")
	attributeMatchesCmd.PersistentFlags().StringVarP(&attributeValue, "value", "v", "", "attribute value")
	cobra.MarkFlagRequired(userAttributesCmd.PersistentFlags(), "username")
	cobra.MarkFlagRequired(userAttributesCmd.PersistentFlags(), "key")
	cobra.MarkFlagRequired(userAttributesCmd.PersistentFlags(), "value")

	userCmd.AddCommand(userListCmd, userAttributesCmd, attributeMatchesCmd)
	rootCmd.AddCommand(userCmd)

}
