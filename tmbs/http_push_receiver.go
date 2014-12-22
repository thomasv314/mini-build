package tmbs

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"
)

// Receives a POST request
func RenderPushNotification(repoName string, res http.ResponseWriter, req *http.Request) {

	// Set the headers for CORS
	res = setHeader(res)

	var (
		jsonI     map[string]interface{}
		userAgent string = req.Header.Get("User-Agent")
		source    string = "none"
		commits   []GitCommit
		err       error
	)

	if strings.Contains(userAgent, "GitHub") {
		source = "github"
		jsonI, _ = parseRequestFormJSON(req)
		commits, err = parseGitHubInterface(repoName, jsonI)
	} else {
		source = "bitbucket"
		jsonI, _ = parseRequestFormJSON(req)
		commits, err = parseBitbucketInterface(repoName, jsonI)
	}

	if source != "" {
		if err != nil {
			fmt.Println(" - Received a push notification from", source, "but we could not decode.")
		} else {
			// We've got a push notification to add
			config, err := LoadConfiguration()

			if err != nil {
				fmt.Println(" - Could not add this push notification to the configuration.")
			} else {
				// Find the Repo we're pushing to in the configuration
				for i := range config.Repositories {
					if repoName == config.Repositories[i].Name {
						config.Repositories[i].Commits = append(config.Repositories[i].Commits, commits...)
						err = WriteJSONFile(GetTmbsDirectory()+"/config.json", &config)
						alertIfError(err, "Can't write json file..")
						break
					}
				}
			}
		}
	} else {
		fmt.Println(" - Received a push notification from an unknown source for", repoName)
	}
}

// Returns a Go interface mapped to the incoming JSON payload
func parseRequestFormJSON(req *http.Request) (js map[string]interface{}, err error) {
	err = req.ParseForm()

	jsonstr := req.PostForm["payload"][0]
	jsonstr, err = url.QueryUnescape(jsonstr)
	exitIfError(err, "Could not unescape string")
	jsonBytes := []byte(jsonstr)

	var jsonInterface map[string]interface{}
	err = json.Unmarshal(jsonBytes, &jsonInterface)

	return jsonInterface, err
}

// Parses a Go interface based off the payload from bitbucket
func parseBitbucketInterface(repoName string, js map[string]interface{}) ([]GitCommit, error) {
	var (
		author     string
		message    string
		id         string
		commitTime time.Time
		err        error
	)

	bitbucketTimeParser := "2006-01-02 15:04:05"

	comsArr := js["commits"].([]interface{})

	commits := make([]GitCommit, len(comsArr))
	for i := range comsArr {
		c := comsArr[i].(map[string]interface{})
		author = c["author"].(string)
		message = c["message"].(string)
		id = c["raw_node"].(string)
		commitTime, err = time.Parse(bitbucketTimeParser, c["timestamp"].(string))
		commits[i] = GitCommit{id, repoName, author, message, commitTime, "received", "bitbucket"}
	}

	return commits, err
}

// Parses a Go interface based off the payload from github
func parseGitHubInterface(repoName string, js map[string]interface{}) ([]GitCommit, error) {

	var (
		author     string
		message    string
		id         string
		commitTime time.Time
		err        error
	)

	comsArr := js["commits"].([]interface{})
	commits := make([]GitCommit, len(comsArr))

	for i := range comsArr {
		c := comsArr[i].(map[string]interface{})

		author = c["author"].(map[string]interface{})["username"].(string)
		message = c["message"].(string)
		id = c["id"].(string)
		commitTime, err = time.Parse(time.RFC3339, c["timestamp"].(string))

		commits[i] = GitCommit{id, repoName, author, message, commitTime, "received", "github"}
	}

	return commits, err
}
