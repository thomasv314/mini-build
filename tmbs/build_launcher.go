package tmbs

import (
	"fmt"
	//	"github.com/libgit2/git2go"
)

// latestCommand := BuildCommand{ "bitbucket", "a380d3a", "http://bitbucket.org/thomas/myrepo.git"
type BuildCommand struct {
	Type       string
	Commit     string
	Repository string
}

var link chan BuildCommand

func FakePush() {
	fmt.Println("Sending fake push")
	repo := "http://bitbucket.org/tommyvyo/arthouse.git"
	commit := "aeb8430c"
	link <- BuildCommand{"bitbucket", commit, repo}
}

func StartBuildLauncher() {
	link = make(chan BuildCommand)
	go start()
}

func start() {
	for {
		build := <-link
		fmt.Println("Recieved push", build.Commit, "from", build.Type)
	}
}
