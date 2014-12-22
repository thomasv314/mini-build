package tmbs

import (
	"fmt"
	"time"
	//	"github.com/libgit2/git2go"
)

// latestCommand := BuildCommand{ "bitbucket", "a380d3a", "http://bitbucket.org/thomas/myrepo.git"
type BuildCommand struct {
	Type       string
	Commit     string
	Repository string
}

var link chan GitCommit

func FakePush() {
	fmt.Println("Sending fake push")
	commit := GitCommit{"23c93b66e188ded46a5d1d0b37add82d21cd05b9", "tommyvyo", "Commit Format Test Two'\n", time.Now(), "received", "bitbucket"}
	fmt.Println(commit)
	link <- commit
}

func StartBuildLauncher() {
	link = make(chan GitCommit)
	go start()
}

func start() {
	for {
		commit := <-link
		BuildNewCommit(commit)
	}
}

func BuildNewCommit(commit GitCommit) {

}
