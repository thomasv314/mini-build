package tmbs

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func RenderHomepage(res http.ResponseWriter, req *http.Request) {

	res = setHeader(res)

	var config Configuration = Configuration{}

	err := LoadJSONFile(GetTmbsDirectory()+"/config.json", &config)
	alertIfError(err, "Could not load JSON.")

	fmt.Println("CONFIG", config)

	bytes, err := json.MarshalIndent(config, " ", " ")

	io.WriteString(res, string(bytes))

}
