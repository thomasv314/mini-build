package tmbs

import (
	git "github.com/libgit2/git2go"
)

func CloneRepository(gitUrl string, path string) (*git.Repository, error) {
	co := git.CloneOptions{}
	repo, err := git.Clone(gitUrl, path, &co)
	return repo, err
}

func FetchLatestRepository(repository *git.Repository) error {
	// list the remotes in the repository (origin by default)
	// TODO think about when this wont be origin and what to do
	remotes, err := repository.ListRemotes()
	if err != nil {
		return err // error listing remotes
	} else {
		// lookup the first remote returned from the repo (origin)
		origin, err := repository.LookupRemote(remotes[0])
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
