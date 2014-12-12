package main

import (
	"fmt"
	"github.com/thomasv314/mini-build/tmbs"
	"os"
	"os/user"
)

var configuration tmbs.Configuration

// Main method for TMBS
func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		startTmbs()
	} else {
		if args[0] == "setup" {
			setupTmbs()
		} else {
			fmt.Println("What?")
		}
	}
}

func startTmbs() {

	configuration = tmbs.LoadConfiguration()

	tmbs.StartPushListener(configuration)
	tmbs.StartBuildMaster(configuration)

	for {
		// Run all goroutines till they press enter.
		var input string
		fmt.Scanln(&input)
		tmbs.FakePush()
	}
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

func setupTmbs() {
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

/*
package tmbs

import (
	"fmt"
)

type BuildAgent struct {
	RepoUrl  string
	RepoType string
}

type BuildAgent interface {
	CloneRepository(url string, path string) bool
}

func CloneRepository(url string, path string) bool {
	fmt.Println("Trying to clone repository from", url)
	gitCloneOpts := git.CloneOptions{CheckoutBranch: "master"}
	repo, err := git.Clone(url, path, &gitCloneOpts)
	if err != nil {
		fmt.Println("Error")
		panic(err)
	}
	fmt.Println("Repository:", repo)
*/
