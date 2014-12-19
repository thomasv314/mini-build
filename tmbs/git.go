package tmbs

import (
	"fmt"
	git "github.com/libgit2/git2go"
	"net/url"
	"time"
)

type GitPushNotification struct {
	PushType string
	Commits  []GitCommit
}

type GitCommit struct {
	Id        string
	Author    string
	Message   string
	Timestamp time.Time
	Status    string
}

func CloneRepository(gitURL *url.URL, path string, bare bool) (*git.Repository, error) {
	fmt.Println("Cloning")
	repo, err := git.Clone(gitURL.String(), path, cloneOptionsForURL(gitURL, bare))
	fmt.Println("\n")
	return repo, err
}

func FetchRepository(repository *git.Repository) error {
	// list the remotes in the repository (origin by default)
	// TODO think about when this wont be origin and what to do
	remotes, err := repository.ListRemotes()
	if err != nil {
		return err // error listing remotes
	} else {
		// lookup the first remote returned from the repo (origin)
		origin, err := repository.LookupRemote(remotes[0])
		origin.SetCallbacks(&remoteCallbacks)
		if err != nil {
			return err // error looking up remote
		} else {
			refspec := make([]string, 0)
			err = origin.Fetch(refspec, nil, "")
			if err != nil {
				return err // Error fetching
			} else {
				return nil // Cool, we fetched the repo
			}
		}
	}
}
