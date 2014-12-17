package tmbs

import (
	"fmt"
	"net/http"
	"strings"
)

// Start a new go routine for the HTTP listener
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

	http.HandleFunc("/", handler)
	err := http.ListenAndServe(config.ListenPort, nil)
	exitIfError(err, "Could not start HTTP listener. Check to see it's not already running.")

}

// Slices up the parameters and decides where to route the request
func handler(res http.ResponseWriter, req *http.Request) {
	params := strings.Split(req.URL.Path[1:], "/")
	switch params[0] {
	case "push":
		{
			logRequest(200, params)
			RenderPushNotification(res, req)
		}
	case "":
		{
			logRequest(200, params)
			RenderHomepage(res, req)
		}
	default:
		{
			logRequest(404, params)
		}
	}
}

// Writes a simple log to std
func logRequest(status int, str []string) {
	fmt.Println("   [", status, "]   ", strings.Join(str, "/"))
}

// Sets the headers for CORS
func setHeader(res http.ResponseWriter) http.ResponseWriter {
	res.Header().Set(
		"Content-Type",
		"text/html",
	)

	res.Header().Set(
		"Access-Control-Allow-Origin",
		"*",
	)

	return res
}
