package tmbs

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

// Starts an HTTP server that listens on a specified port and
// launches new builds when post requests come in to /post/
func StartHttpListener(config Configuration) {
	go startListener(config.ListenPort)
}

func startListener(listenPort string) {
	fmt.Println("Listening on", listenPort)
	http.HandleFunc("/push", recievedPushNotification)
	err := http.ListenAndServe(listenPort, nil)
	if err != nil {
		fmt.Println("Error. Could not start the HTTP listener.")
		fmt.Println(err)
		os.Exit(1)
	}
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

	req.ParseForm()
	fmt.Println("Form:", req.Form)

	repo := "http://bitbucket.org/tommyvyo/arthouse.git"
	commit := "aeb8430c"

	link <- BuildCommand{"bitbucket", commit, repo}

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

}
