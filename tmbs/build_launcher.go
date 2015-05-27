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
	commit := GitCommit{
		"23c93b66e188ded46a5d1d0b37add82d21cd05b9",
		"mini-build-bb",
		"tommyvyo",
		"Commit Format Test Two'\n",
		time.Now(),
		"received",
		"bitbucket",
	}

	BuildNewCommit(commit)
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

	filePath := getLogPath(commit.RepositoryName, commit.Slug())

	// Start a new test writer to write a log file and send any updates to websockets
	testWriter := TestWriter{filePath}
	testWriter.Setup()

	testWriter.Write("Received commit " + commit.Id)

	path := "file://" + GetTmbsDirectory() + "/repos/" + commit.RepositoryName
	url := parseURL(path)

	clonePath := GetTmbsDirectory() + "/tests/" + commit.Slug()

	newRepo, err := CloneRepository(url, clonePath, false, testWriter)

	fmt.Println(newRepo)

	if err != nil {
		fmt.Println("Error", err)
	} else {
		testWriter.Write("\nCloned the repository to " + path)
	}

}
