package tmbs

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// Starts an HTTP server that listens on a specified port and
// launches new builds when post requests come in to /post/
func StartHttpListener() {
	go startListener()
}

func startListener() {

	fmt.Println(
		" - listening for commits to",
		len(config.Repositories),
		"repositories on", "0.0.0.0",
		config.ListenPort,
	)

	http.HandleFunc("/push", recievedPushNotification)
	err := http.ListenAndServe(config.ListenPort, nil)
	exitIfError(err, "Could not start HTTP listener. Check to see it's not already running.")

}

func setHeader(res http.ResponseWriter) {

	res.Header().Set(
		"Content-Type",
		"text/html",
	)

	res.Header().Set(
		"Access-Control-Allow-Origin",
		"*",
	)

}

func recievedPushNotification(res http.ResponseWriter, req *http.Request) {

	setHeader(res)

	var js interface{}

	decoder := json.NewDecoder(req.Body)
	err := decoder.Decode(js)

	if err != nil {
		fmt.Println("ERROR", err)
	} else {
		jsbytes, err := json.MarshalIndent(js, " ", "  ")
		exitIfError(err, "Could not indent received request.")

		err = ioutil.WriteFile(GetTmbsDirectory()+"/myreq.json", jsbytes, 0644)
		exitIfError(err, "Could not write the requ.")
		fmt.Println("Recieved and wrote request")
	}

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
