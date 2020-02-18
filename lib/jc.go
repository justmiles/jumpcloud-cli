package jc

import (
	"context"
	"log"
	"os"

	"github.com/TheJumpCloud/jcapi"
	jcapiv2 "github.com/TheJumpCloud/jcapi-go/v2"
	"github.com/olekukonko/tablewriter"
)

var (
	apiURL      = config("JUMPCLOUD_API_URL", "https://console.jumpcloud.com/api")
	apiKey      = config("JUMPCLOUD_API_KEY", "")
	apiClientV1 = jcapi.NewJCAPI(apiKey, apiURL)
	apiClientV2 *jcapiv2.APIClient
	isGroups    bool
)

func init() {
	var err error
	// check if the org is on tags or groups:

	if apiKey == "" {
		log.Fatal("Environment variable JUMPCLOUD_API_KEY not set")
	}
	isGroups, err = isGroupsOrg(apiURL, apiKey)
	if err != nil {
		log.Fatalf("Could not determine your org type, err='%s'\n", err)
	}

	// if we're on a groups org, instantiate API client v2:
	if isGroups {
		apiClientV2 = jcapiv2.NewAPIClient(jcapiv2.NewConfiguration())
		apiClientV2.ChangeBasePath(apiURL + "/v2")
	}
}

func isGroupsOrg(urlBase string, apiKey string) (bool, error) {
	// instantiate a new API client object:
	client := jcapiv2.NewAPIClient(jcapiv2.NewConfiguration())
	client.ChangeBasePath(urlBase + "/v2")

	// set up the API key via context:
	auth := context.WithValue(context.TODO(), jcapiv2.ContextAPIKey, jcapiv2.APIKey{
		Key: apiKey,
	})

	// set up optional parameters:
	optionals := map[string]interface{}{
		"limit": int32(1), // limit the query to return 1 item
	}
	// in order to check for groups support, we just query for the list of User groups
	// (we just ask to retrieve 1) and check the response status code:
	_, res, err := client.UserGroupsApi.GroupsUserList(auth, "application/json", "application/json", optionals)

	// check if we're using the API v1:
	// we need to explicitly check for 404, since GroupsUserList will also return a json
	// unmarshalling error (err will not be nil) if we're running this endpoint against
	// a Tags org and we don't want to treat this case as an error:
	if res != nil && res.StatusCode == 404 {
		return false, nil
	}

	// if there was any kind of other error, return that:
	if err != nil {
		return false, err
	}

	// if we're using API v2, we're expecting a 200:
	if res.StatusCode == 200 {
		return true, nil
	}

	return false, nil
}

func renderTable(header []string, data [][]string) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader(header)
	table.SetBorders(tablewriter.Border{Left: true, Top: false, Right: true, Bottom: false})
	table.SetCenterSeparator("|")
	table.AppendBulk(data) // Add Bulk Data
	table.Render()
}

func config(s, e string) string {
	envVar := os.Getenv(s)
	if envVar != "" {
		return envVar
	}
	return e
}
