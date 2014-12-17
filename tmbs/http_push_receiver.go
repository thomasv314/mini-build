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
	)

	var pushNotification GitPushNotification

	if strings.Contains(userAgent, "GitHub") {

		jsonI, _ = parseRequestFormJSON(req)
		pushNotification, err := parseGitHubInterface(jsonI)
		alertIfError(err, "Could not parse github interface.")

	} else {

		jsonI, _ = parseRequestFormJSON(req)
		pushNotification, err := parseBitbucketInterface(jsonI)
		alertIfError(err, "Could not parse bitbucket interface.")

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

// Parses a Go interface based off the payload and creates a push notification
func parseGitHubInterface(js map[string]interface{}) (GitPushNotification, error) {

	notification := GitPushNotification{PushType: "GitHub"}

	var (
		author     string
		message    string
		id         string
		commitTime time.Time
		err        error
	)

	comsArr := js["commits"].([]interface{})

	for i := range comsArr {
		c := comsArr[i].(map[string]interface{})

		author = c["author"].(map[string]interface{})["username"].(string)
		message = c["message"].(string)
		id = c["id"].(string)

		commitTime, err = time.Parse(time.RFC3339, c["timestamp"].(string))

		commit := GitCommit{id, author, message, commitTime}

		notification.Commits = append(notification.Commits, commit)
		fmt.Println(commit)
	}

	return notification, err
}
