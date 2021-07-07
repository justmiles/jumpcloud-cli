package jc

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/TheJumpCloud/jcapi"

	jcapiv2 "github.com/TheJumpCloud/jcapi-go/v2"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

// ListGroups will return all JumpCLoud Groups
func ExecuteCommandAgainstSystem(system, command string) (commandResults []jcapi.JCCommandResult, err error) {
	var options = make(map[string]interface{})
	options["limit"] = int32(100)

	// Create the command

	var jcCommand = jcapi.JCCommand{
		Name:       fmt.Sprintf("jc-cli-%s", randomId()),
		Command:    command,
		User:       jcapi.COMMAND_ROOT_USER,
		LaunchType: "manual",
		Schedule:   "immediate",
		Timeout:    "0", // TODO: support timeout from a flag
		ListensTo:  "",
		Trigger:    "",
		Sudo:       false,
		Skip:       0,
		Limit:      10,
	}

	jcCommand, err = apiClientV1.AddUpdateCommand(jcapi.Insert, jcCommand)
	if err != nil {
		return nil, fmt.Errorf("Could not add/update command, err='%s'\n", err)
	}

	localOpts := make(map[string]interface{})
	var graphType jcapiv2.GraphType
	graphType = "system"

	localOpts["body"] = jcapiv2.GraphManagementReq{
		Id:    system,
		Op:    "add",
		Type_: &graphType,
	}

	// associate the command with relevant instances
	_, err = apiClientV2.CommandsApi.GraphCommandAssociationsPost(apiClientV2Auth, jcCommand.Id, contentType, contentType, localOpts)
	if err != nil {
		return nil, fmt.Errorf("Could not graph command association, err='%s'\n", err)
	}

	// Run the command
	err = apiClientV1.RunCommand(jcCommand)
	if err != nil {
		return nil, fmt.Errorf("Could not run command, err='%s'\n", err)
	}

	// get the command results
	for {

		results, err := apiClientV1.GetCommandResultsByName(jcCommand.Name)
		if err != nil {
			return nil, fmt.Errorf("Could not find the command result for '%s', err='%s'", jcCommand.Name, err.Error())
		}

		// Walk the results and add their exit code to the map (maps system name to the result data)
		for _, result := range results {
			res, err := apiClientV1.GetCommandResultDetailsById(result.Id)
			if err != nil {
				return commandResults, fmt.Errorf("Could not get command result details by ID, err='%s'", err.Error())
			}
			commandResults = append(commandResults, res)
		}

		if len(results) > 0 {
			break
		}

		time.Sleep(10 * time.Second) // TODO: add some sort of exponential backoff to prevent rate limits

	}

	// TODO: Delete the command
	// err = apiClientV1.DeleteCommand(jcCommand)

	// return the results
	return commandResults, err
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyz")

func randomId() string {
	n := 6
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}
