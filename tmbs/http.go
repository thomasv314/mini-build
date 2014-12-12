package tmbs

import (
	"fmt"
	"io"
	"net/http"
)

func StartPushListener(config Configuration) {
	go startListener(config.ListenPort)
}

func startListener(listenPort string) {
	fmt.Println("Started listening for push notifications on port ", listenPort)
	http.HandleFunc("/push", recievedPushNotification)
	http.ListenAndServe(listenPort, nil)
}

func recievedPushNotification(res http.ResponseWriter, req *http.Request) {

	res.Header().Set(
		"Content-Type",
		"text/html",
	)

	res.Header().Set(
		"Access-Control-Allow-Origin",
		"*",
	)

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
