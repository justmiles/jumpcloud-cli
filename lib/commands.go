package jc

import (
	"fmt"
	"log"

	"github.com/TheJumpCloud/jcapi"
)

// ListGroups will return all JumpCLoud Groups
func ExecuteCommandAgainstSystem(system, command string) {

	var options = make(map[string]interface{})
	options["limit"] = int32(100)

	// Create the command

	var jcCommand = jcapi.JCCommand{
		Name:           "jcapi-command",
		Command:        command,
		CommandType:    "linux",
		CommandRunners: []string{"5a5d2f9827ed807b0b705e0c"}, // TODO get current user
		User:           jcapi.COMMAND_ROOT_USER,
		LaunchType:     "manual",
		Schedule:       "immediate",
		Timeout:        "0", // No timeout
		ListensTo:      "",
		Trigger:        "",
		Sudo:           false,
		// Shell:       "shell",
		Systems: []string{"60118e1ea6a53c191da3452a"},
		Skip:    0,
		Limit:   10,
	}

	jcCommand, err := apiClientV1.AddUpdateCommand(jcapi.Insert, jcCommand)
	if err != nil {
		log.Fatalf("Could not add/update command, err='%s'\n", err)
	}

	fmt.Println(jcCommand)

	// Run the command
	err = apiClientV1.RunCommand(jcCommand)
	if err != nil {
		log.Fatalf("Could not run command, err='%s'\n", err)
	}

	// get the command results

	// Delete the command

}
