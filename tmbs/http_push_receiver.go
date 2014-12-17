package tmbs

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func RenderPushNotification(res http.ResponseWriter, req *http.Request) {

	// Set the headers for CORS
	res = setHeader(res)

	var (
		jsonInterface interface{}
		userAgent     string = req.Header.Get("User-Agent")
		source        string
		err           error
	)

	if userAgent == "Bitbucket.org" {

		source = "bitbucket"
		jsonInterface, err = parseRequestFormJSON(req, true)

	} else if strings.Contains(userAgent, "GitHub") {

		source = "github"
		jsonInterface, err = parseRequestFormJSON(req, false)

	} else {

		fmt.Println("Received a request from unknown source")
		fmt.Println("\tUser agent", req.Header.Get("User-Agent"))
		req.ParseForm()
		fmt.Println("\tbody", req.Body)
		fmt.Println("\tform", req.PostForm)

	}

	if source != "" {
		jsonbytes, err := json.MarshalIndent(jsonInterface, " ", " ")

		if err != nil {

			fmt.Println("Error indenting JSON")
		}

		ioutil.WriteFile("latest.json", jsonbytes, 0644)
	}

	if err != nil {
		fmt.Println("Some error", err)
	}

	//repo := "http://bitbucket.org/tommyvyo/arthouse.git"
	//	commit := "aeb8430c"
	//	link <- BuildCommand{"bitbucket", commit, repo}
	// Format Test One
	// Format Test Two

}

func parseRequestFormJSON(req *http.Request, needsQueryUnescape bool) (js interface{}, err error) {
	err = req.ParseForm()

	jsonstr := req.PostForm["payload"][0]
	jsonstr, err = url.QueryUnescape(jsonstr)
	exitIfError(err, "Could not unescape string")
	jsonBytes := []byte(jsonstr)

	var jsonInterface interface{}
	err = json.Unmarshal(jsonBytes, &jsonInterface)

	fmt.Println(jsonInterface)

	return jsonInterface, err
}
