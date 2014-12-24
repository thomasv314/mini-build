package tmbs

import (
	"fmt"
	//	git "github.com/libgit2/git2go"
	"time"
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
	commit := GitCommit{"23c93b66e188ded46a5d1d0b37add82d21cd05b9", "mini-build", "tommyvyo", "Commit Format Test Two'\n", time.Now(), "received", "bitbucket"}
	fmt.Println(commit)
	link <- commit
}

func StartBuildLauncher() {
	link = make(chan GitCommit)
	// listen for new commits on a new goroutine
	go start()
}

func start() {
	for {
		commit := <-link
		go BuildNewCommit(commit)
	}
}

func BuildNewCommit(commit GitCommit) {
	fmt.Println("Recieved commit ", commit.Id, "on a new go routine")

	path := GetTmbsDirectory() + "/repos/" + commit.RepositoryName

	newRepo, err := CloneRepository(parseURL("file://"+path), GetTmbsDirectory()+"/repos/watwat", false)

	//	repo, err := git.OpenRepository(path)

	if err != nil {
		fmt.Println("Error", err)
	} else {

		fmt.Println("Repository!", newRepo)
	}
}
