package tmbs

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

func RenderPushNotification(res http.ResponseWriter, req *http.Request) {

	// Set the headers for CORS
	res = setHeader(res)

	if req.Header.Get("User-Agent") == "Bitbucket.org" {
		fmt.Println("Parsing push notification from bitbucket.")
		// Bitbucket sends it's JSON url encoded in the form
		err := req.ParseForm()
		jsonStr, err := url.QueryUnescape(req.PostForm["payload"][0])
		exitIfError(err, "Could not unescape")
		jsonBytes := []byte(jsonStr)

		// Unmarshal payload into temporary interface
		var tempInterface interface{}
		err = json.Unmarshal(jsonBytes, &tempInterface)
		alertIfError(err, "Can't unmarshal")

		// Marshal, indent, and write temp interface to a file
		jsonbytes, err := json.MarshalIndent(tempInterface, " ", " ")
		ioutil.WriteFile("bitbucket.json", jsonbytes, 0644)

	} else {

		fmt.Println("Received a request from unkown")
		fmt.Println("User agent", req.Header.Get("User-Agent"))
		fmt.Println("body", req.Body)
		fmt.Println("form", req.PostForm)

	}
	//repo := "http://bitbucket.org/tommyvyo/arthouse.git"
	//	commit := "aeb8430c"

	//	link <- BuildCommand{"bitbucket", commit, repo}

}
