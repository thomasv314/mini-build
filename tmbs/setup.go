package tmbs

import (
	"fmt"
	"os"
	"os/user"
)

func Setup() {
	usr, err := user.Current()
	if err != nil {
		panic(err)
	}

	fmt.Println("Setting up", usr.HomeDir)

	appDir := usr.HomeDir + "/.tmbs"
	repoDir := appDir + "/repos"
	testDir := appDir + "/tests"

	checkOrCreateDir(appDir, "application", 0777)
	checkOrCreateDir(repoDir, "repository", 0777)
	checkOrCreateDir(testDir, "tests", 0777)

	fmt.Println("Good to go!")
}

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
