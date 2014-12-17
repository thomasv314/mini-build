package tmbs

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

// Simple JSON configuration struct
type Configuration struct {
	ListenPort   string ":59999" // json: "listenPort"		: ":55559"
	Repositories []WatchedRepository
}

type WatchedRepository struct {
	Directory string
	Name      string
	Pushes    []GitPushNotification
}

var (
	config_file string = "/config.json"
	AppConfig   Configuration
)

// Loads a configuration file from config.json in the directory
func LoadConfiguration() (Configuration, error) {

	AppConfig := Configuration{}

	dir := GetTmbsDirectory() + config_file

	file, err := os.Open(dir)
	defer file.Close()

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&AppConfig)

	return AppConfig, err
}

func SaveConfiguration(config *Configuration) {

	cfgdir := GetTmbsDirectory() + config_file

	jsonbytes, err := json.MarshalIndent(config, " ", "  ")

	if err != nil {
		fmt.Println("Hm, could not marshal config.")
	} else {
		err = ioutil.WriteFile(cfgdir, jsonbytes, 0644)
		if err != nil {
			fmt.Println("Error writing the fockinnngg file")
		}
		exitIfError(err, "Could not save configuration file.")

	}
}
