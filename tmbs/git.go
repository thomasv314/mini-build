package tmbs

import (
	"fmt"
	git "github.com/libgit2/git2go"
	"net/url"
)

func CloneRepository(gitURL *url.URL, path string, bare bool) (*git.Repository, error) {
	fmt.Println("Cloning")
	repo, err := git.Clone(gitURL.String(), path, cloneOptionsForURL(gitURL, bare))
	fmt.Println("repo", repo)
	fmt.Println("err", err)
	fmt.Println("\n")
	return repo, err
}

func FetchRepository(repository *git.Repository) error {
	// list the remotes in the repository (origin by default)
	// TODO think about when this wont be origin and what to do
	var (
		remotes []string
		err     error
	)

	remotes, err = repository.ListRemotes()

	if err != nil {
		fmt.Println("wtf?")
		fmt.Println(err)
		return err // error listing remotes
	} else {
		// lookup the first remote returned from the repo (origin)
		origin, err := repository.LookupRemote(remotes[0])
		origin.SetCallbacks(&remoteCallbacks)
		if err != nil {
			Log.Print("Cannot lookup remote [0]")
			return err // error looking up remote
		} else {
			refspec := make([]string, 0)
			err = origin.Fetch(refspec, nil, "")
			if err != nil {
				Log.Print("Error fetching the origin from the refspec")
				Log.Print(refspec)
				return err // Error fetching
			} else {
				Log.Print("Cannot fetch repo")
				return nil // Cool, we fetched the repo
			}
		}
	}
}
