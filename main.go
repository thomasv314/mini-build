package main

import (
	"fmt"
	//	"github.com/libgit2/git2go"
	"github.com/thomasv314/mini-build/tmbs"
)

var configuration tmbs.Configuration

// Main method for TMBS
func main() {

	configuration = tmbs.LoadConfiguration()
	tmbs.StartPushListener(configuration.ListenPort)

	// Run all goroutines till they press enter.
	var input string
	fmt.Scanln(&input)
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
