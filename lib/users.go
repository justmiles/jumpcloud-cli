package jc

import (
	"fmt"
	"log"
	"strings"

	"github.com/TheJumpCloud/jcapi"
)

// ListUsers will return all JumpCLoud users
func ListUsers() {
	// Grab all system users (with their tags if this is a Tags org):
	userList, err := apiClientV1.GetSystemUsers(!isGroups)
	if err != nil {
		log.Fatalf("Could not read system users, err='%s'\n", err)
	}

	header := []string{
		"Id",
		"UserName",
		"FirstName",
		"LastName",
		"Email",
		"Uid",
		"Gid",
		"Activated",
		"PasswordExpired",
	}

	renderTable(header, userListToString(userList))

}

// UserAttributeMatches will return attributes for a selected user
func UserAttributeMatches(userName, attributeName, attributeValue string) (bool, error) {
	user, err := getUserByName(userName)
	if err != nil {
		return false, err
	}
	for _, attribute := range user.Attributes {
		if attribute.Name == attributeName {
			return strings.Contains(strings.ToLower(attribute.Value), strings.ToLower(attributeValue)), nil
		}
	}

	return false, fmt.Errorf("attribute %s not found", attributeName)

}

// UserAttributes will return attributes for a selected user
func UserAttributes(userName string) error {
	user, err := getUserByName(userName)
	if err != nil {
		return err
	}

	var header, data []string
	for _, attribute := range user.Attributes {
		header = append(header, attribute.Name)
		data = append(data, attribute.Value)
	}

	renderTable(header, [][]string{data})

	return nil
}

func getUserByName(userName string) (jcapi.JCUser, error) {
	userList, err := apiClientV1.GetSystemUsers(!isGroups)
	if err != nil {
		return jcapi.JCUser{}, err
	}
	for _, user := range userList {
		if user.UserName == userName {
			return user, nil
		}
	}
	return jcapi.JCUser{}, fmt.Errorf("unable to find jumpcloud user with name %s", userName)
}

func userListToString(userList []jcapi.JCUser) (userListString [][]string) {
	for _, user := range userList {
		// for _, userAttributes := range user.Attributes {
		// 	fmt.Printf("%s: %s: %s\n", user.UserName, userAttributes.Name, userAttributes.Value)
		// }
		userListString = append(userListString,
			[]string{user.Id,
				user.UserName,
				user.FirstName,
				user.LastName,
				user.Email,
				user.Uid,
				user.Gid,
				fmt.Sprintf("%t", user.Activated),
				fmt.Sprintf("%t", user.PasswordExpired),
			})
	}
	return userListString
}
