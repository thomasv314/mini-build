package tmbs

import (
	"io"
	"net/http"
)

func RenderHomepage(res http.ResponseWriter, req *http.Request) {

	res = setHeader(res)
	io.WriteString(res, "<html><head><title>TMBS</title></head><body>Thomas' Mini Build Server</body></html>")

}
