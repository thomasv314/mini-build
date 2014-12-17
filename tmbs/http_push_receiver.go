package tmbs

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

func RenderPushNotification(res http.ResponseWriter, req *http.Request) {
	res = setHeader(res)

	var js interface{}

	decoder := json.NewDecoder(req.Body)
	err := decoder.Decode(&js)
	if err != nil {
		fmt.Println("ERROR DECODING", err)

		var body []byte
		body, err := ioutil.ReadAll(req.Body)

		sample := body

		fmt.Println("Printf with %x:")
		fmt.Printf("%x\n", sample)

		fmt.Println("Printf with % x:")
		fmt.Printf("% x\n", sample)

		fmt.Println("Printf with %q:")
		fmt.Printf("%q\n", sample)

		fmt.Println("Printf with %+q:")
		fmt.Printf("%+q\n", sample)

		// This line prints out the output above
		fmt.Println(string(body))
		exitIfError(err, "What the heck")
	}
	jsonbytes, err := json.MarshalIndent(js, " ", " ")

	if err != nil {
		fmt.Println("Received invalid json?")
	} else {
		err = ioutil.WriteFile("example.json", jsonbytes, 0644)
		exitIfError(err, "Couldn't save file.")
	}

	io.WriteString(res, string(jsonbytes))

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
