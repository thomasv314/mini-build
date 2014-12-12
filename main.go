package main

import (
	"fmt"
	"github.com/thomasv314/mini-build/tmbs"
	"os"
)

// Main method for TMBS
func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		tmbs.StartBuildServer()
	} else {
		if args[0] == "setup" {
			tmbs.Setup()
		} else {
			fmt.Println("What? That's not a command")
		}
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
