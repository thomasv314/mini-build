package tmbs

import (
	"encoding/json"
	"fmt"
	"os"
)

// Simple JSON configuration struct
type Configuration struct {
	ListenPort   string   // json: "listenPort"
	Repositories []string // json: "repos"
}

// Loads a configuration file from config.json in the directory
func LoadConfiguration() Configuration {
	fmt.Println("Thomas' Mini Build Server - hit return to quit")

	filename := "config.json"
	file, _ := os.Open(filename)
	decoder := json.NewDecoder(file)
	configuration := Configuration{}
	err := decoder.Decode(&configuration)

	if err != nil {
		fmt.Println("There was an error loading your configuration file.", "Error:", err)
	}
	fmt.Println("Loaded", filename)

	return configuration
}
