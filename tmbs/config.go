package tmbs

import (
	"encoding/json"
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
	file, _ := os.Open(config_file)
	encoder := json.NewEncoder(file)
	err := encoder.Encode(&config)
	exitIfError(err, "Could not save configuration file.")
}
