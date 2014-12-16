package tmbs

import (
	"encoding/json"
	"fmt"
	"os"
	"os/user"
)

// Returns the default TMBS directory
func GetTmbsDirectory() string {
	usr, err := user.Current()
	if err != nil {
		panic(err)
	}
	return usr.HomeDir + "/.tmbs"
}

// Setup the default directories for TMBS. Create an empty configuration file with sane defaults.
func Setup() {
	appDir := GetTmbsDirectory()

	// Create the directories
	checkOrCreateDir(appDir, "application", 0777)
	checkOrCreateDir(appDir+"/repos", "repository", 0777)
	checkOrCreateDir(appDir+"/tests", "tests", 0777)

	// Create a blank configuration file
	file, err := os.Create(appDir + "/config.json")
	defer file.Close()
	exitIfError(err, "Can't create empty config.json")

	// Create a default configuration
	emptyConfig := Configuration{ListenPort: ":59999", Repositories: make([]WatchedRepository, 0)}

	// Encode the empty configuration and write it to the file
	encoder := json.NewEncoder(file)
	err = encoder.Encode(&emptyConfig)
	exitIfError(err, "Failed to encode empty configuration file.")

	fmt.Println("Created file:", appDir+"/config.json")
	fmt.Println("TMBS setup correctly.")
}

// Check if a directory exists, if not create it.
func checkOrCreateDir(path string, name string, perm os.FileMode) {
	if checkExists(path) {
		fmt.Println("Already exists:", path)
	} else {
		err := os.Mkdir(path, perm)
		if err != nil {
			fmt.Println("Error creating", name, "directory.")
			panic(err)
		}
		fmt.Println("Created directory:", path)
	}
}

// Check if a file exists
func checkExists(path string) bool {
	if _, err := os.Stat(path); err != nil {
		if os.IsNotExist(err) {
			return false
			// file does not exist
		} else {
			fmt.Println("Error checking file:", err)
			return true
		}
	} else {
		return true
	}
}
