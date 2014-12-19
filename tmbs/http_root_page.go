package tmbs

import (
	"encoding/json"
	"io"
	"net/http"
)

func RenderHomepage(res http.ResponseWriter, req *http.Request) {

	res = setHeader(res)

	var config Configuration = Configuration{}

	var response string

	err := LoadJSONFile(GetTmbsDirectory()+"/config.json", &config)
	if err != nil {
		response = "Could not load config.json"
	} else {
		bytes, err := json.MarshalIndent(config, " ", " ")
		if err != nil {
			response = "Could not parse config.json"
		} else {
			response = string(bytes)
		}
	}

	io.WriteString(res, response)

}
