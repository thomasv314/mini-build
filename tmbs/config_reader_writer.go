package tmbs

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

func LoadJSONFile(path string, inp interface{}) error {
	file, err := os.Open(path)
	defer file.Close()

	if err != nil {
		return err
	}

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&inp)
	if err != nil {
		return err
	} else {
		return nil
	}
}

func WriteJSONFile(path string, js interface{}) error {
	jsonbytes, err := json.MarshalIndent(js, " ", " ")
	if err != nil {
		return err
	} else {
		err = ioutil.WriteFile(path, jsonbytes, 0644)
		if err != nil {
			return err
		} else {
			return nil
		}
	}
}
