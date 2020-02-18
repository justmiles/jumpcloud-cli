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
	attributes                              []string
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
		jc.ListUsers(attributes)
	},
}

var userAttributesCmd = &cobra.Command{
	Use:   "attribute-list",
	Short: "show attributes for a user",
	Run: func(cmd *cobra.Command, args []string) {
		err := jc.UserAttributes(userName)
		if err != nil {
			log.Fatal(err)
		}
	},
}

var attributeMatchesCmd = &cobra.Command{
	Use:   "attribute-match",
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

var getAttribute = &cobra.Command{
	Use:   "attribute-get",
	Short: "get an attribute for a user",
	Run: func(cmd *cobra.Command, args []string) {
		attr, err := jc.GetUserAttribute(userName, attributeName)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(attr)
	},
}

var setAttribute = &cobra.Command{
	Use:   "attribute-set",
	Short: "set an attribute for a user",
	Run: func(cmd *cobra.Command, args []string) {
		err := jc.SetUserAttribute(userName, attributeName, attributeValue)
		if err != nil {
			log.Fatal(err)
		}
	},
}

var deleteAttribute = &cobra.Command{
	Use:   "attribute-delete",
	Short: "delte a user attribute",
	Run: func(cmd *cobra.Command, args []string) {
		err := jc.DeleteUserAttribute(userName, attributeName)
		if err != nil {
			log.Fatal(err)
		}
	},
}

func init() {

	// user-list
	userListCmd.PersistentFlags().StringArrayVarP(&attributes, "attribute", "a", attributes, "list of attributes to be added to the report")

	// attribute-list
	userAttributesCmd.PersistentFlags().StringVarP(&userName, "username", "u", "", "username")
	cobra.MarkFlagRequired(userAttributesCmd.PersistentFlags(), "username")

	// attribute-match
	attributeMatchesCmd.PersistentFlags().StringVarP(&userName, "username", "u", "", "username")
	attributeMatchesCmd.PersistentFlags().StringVarP(&attributeName, "key", "k", "", "attribute name")
	attributeMatchesCmd.PersistentFlags().StringVarP(&attributeValue, "value", "v", "", "attribute value")
	cobra.MarkFlagRequired(attributeMatchesCmd.PersistentFlags(), "username")
	cobra.MarkFlagRequired(attributeMatchesCmd.PersistentFlags(), "key")
	cobra.MarkFlagRequired(attributeMatchesCmd.PersistentFlags(), "value")

	// attribute-get
	getAttribute.PersistentFlags().StringVarP(&userName, "username", "u", "", "username")
	getAttribute.PersistentFlags().StringVarP(&attributeName, "key", "k", "", "attribute name")
	cobra.MarkFlagRequired(getAttribute.PersistentFlags(), "username")
	cobra.MarkFlagRequired(getAttribute.PersistentFlags(), "key")

	// attribute-set
	setAttribute.PersistentFlags().StringVarP(&userName, "username", "u", "", "username")
	setAttribute.PersistentFlags().StringVarP(&attributeName, "key", "k", "", "attribute name")
	setAttribute.PersistentFlags().StringVarP(&attributeValue, "value", "v", "", "attribute value")
	cobra.MarkFlagRequired(setAttribute.PersistentFlags(), "username")
	cobra.MarkFlagRequired(setAttribute.PersistentFlags(), "key")
	cobra.MarkFlagRequired(setAttribute.PersistentFlags(), "value")

	// attribute-delete
	deleteAttribute.PersistentFlags().StringVarP(&userName, "username", "u", "", "username")
	deleteAttribute.PersistentFlags().StringVarP(&attributeName, "key", "k", "", "attribute name")
	cobra.MarkFlagRequired(deleteAttribute.PersistentFlags(), "username")
	cobra.MarkFlagRequired(deleteAttribute.PersistentFlags(), "key")

	userCmd.AddCommand(userListCmd, userAttributesCmd, attributeMatchesCmd, getAttribute, setAttribute, deleteAttribute)
	rootCmd.AddCommand(userCmd)

}
