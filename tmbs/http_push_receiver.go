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

	fmt.Println("Trying to figure out URl struct")
	fmt.Println("req.Host", req.Host)
	fmt.Println(req.Header)
	err := req.ParseForm()

	if err != nil {
		fmt.Println("Error parsing form")
	} else {
		jsonStr, err := url.QueryUnescape(req.PostForm["payload"][0])
		exitIfError(err, "Could not unescape")

		fmt.Println("JSON STR\n", jsonStr)

		jsonStrByte := []byte(jsonStr)

		var tempInterface interface{}
		err = json.Unmarshal(jsonStrByte, &tempInterface)

		jsonbytes, err := json.MarshalIndent(tempInterface, " ", " ")

		ioutil.WriteFile("bitbucket.json", jsonbytes, 0644)
	}

	fmt.Println(req.PostForm)

	//	jsonbytes, err := json.MarshalIndent(js, " ", " ")

	//	if err != nil {
	//		fmt.Println("Received invalid json?")
	//	} else {
	//		err = ioutil.WriteFile("example.json", jsonbytes, 0644)
	//		exitIfError(err, "Couldn't save file.")
	//	}

	//	io.WriteString(res, string(jsonbytes))

	// Print what we need
	//	fmt.Println(string(body))
	//	io.WriteString(res, string(body))

	//	req.ParseForm()
	//	fmt.Println("Form:", req.Form)

	//repo := "http://bitbucket.org/tommyvyo/arthouse.git"
	//	commit := "aeb8430c"

	//	link <- BuildCommand{"bitbucket", commit, repo}
	/*
		io.WriteString(
			res,
			`<doctype html>
			<html>
			<head>
			<title>Hello World</title>
			</head>
			<body>
			Hello World!
			</body>
			</html>`,
		)
	*/

}
