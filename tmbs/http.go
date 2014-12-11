package tmbs

import (
	"fmt"
	"io"
	"net/http"
)

func StartPushListener(listenPort string) {
	go startListener(listenPort)
}

func startListener(listenPort string) {
	fmt.Println("Now listening for push notifications on port ", listenPort)
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

	fmt.Println("Bitbucket push posted. Form:")

	req.ParseForm()
	fmt.Println("Form:", req.Form)

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
