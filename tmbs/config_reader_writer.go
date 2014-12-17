package tmbs

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

func LoadJSONFile(path string, inp interface{}) error {
	file, err := os.Open(path)
	defer file.Close()

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&inp)

	return err
}

func WriteJSONFile(path string, js interface{}) error {
	jsonbytes, err := json.MarshalIndent(js, " ", " ")
	if err != nil {
		return err
	} else {
		err = ioutil.WriteFile(path, jsonbytes, 0644)
		if err != nil {
			return err
		}
		return nil
	}
}
