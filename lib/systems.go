package jc

import (
	"log"
)

// ListSystems will return all JumpCLoud systems
func ListSystems(query, outputFormat string) {
	systemList, err := apiClientV1.GetSystems(false)
	if err != nil {
		log.Fatalf("Could not read system users, err='%s'\n", err)
	}

	outputData(outputFormat, query, systemList)

}
