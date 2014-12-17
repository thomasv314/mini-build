package tmbs

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

func RenderPushNotification(res http.ResponseWriter, req *http.Request) {
	res = setHeader(res)

	err := req.ParseForm()

	if err != nil {
		fmt.Println("Error parsing form")
	} else {
		jsonStr, err := url.QueryUnescape(req.PostForm["payload"][0])
		exitIfError(err, "Could not unescape")

		//		var jsonInt interface{}

		//		json.Unmarshal([]byte(jsonStr), jsonInt)

		jsonbytes, err := json.MarshalIndent([]byte(jsonStr), " ", " ")
		exitIfError(err, " Dang wtf Could not indent")
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
