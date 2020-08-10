package jc

import (
	"log"

	"github.com/TheJumpCloud/jcapi"
)

const (
	contentType = "application/json"
)

// ListGroups will return all JumpCLoud Groups
func ListGroups(query, output string) {

	var options = make(map[string]interface{})
	options["limit"] = int32(100)
	res, _, err := apiClientV2.GroupsApi.GroupsList(apiClientV2Auth, contentType, contentType, options)
	if err != nil {
		log.Fatalf("Could not read system users, err='%s'\n", err)
	}

	outputData(output, query, res)
}

// GraphUserGroupMembersList will return all users in group
func GraphUserGroupMembersList(groupID, query, output string) {

	var options = make(map[string]interface{})
	options["limit"] = int32(100)
	res, _, err := apiClientV2.UserGroupsApi.GraphUserGroupMembersList(apiClientV2Auth, groupID, contentType, contentType, options)
	if err != nil {
		log.Fatalf("Could not read system users, err='%s'\n", err)
	}

	var users []jcapi.JCUser
	for _, connection := range res {
		user, err := apiClientV1.GetSystemUserById(connection.To.Id, true)
		if err != nil {
			log.Fatalf("Could not get user details, err='%s'\n", err)
		}
		users = append(users, user)
	}

	outputData(output, query, users)
}
