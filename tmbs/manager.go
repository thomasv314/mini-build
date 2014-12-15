package tmbs

import (
	"fmt"
	"net/url"
	"os"
	"os/user"
	"path"
	"strings"
)

func AddRepository(name string, gitUrl string) {
	// Parse the gitURL to work w/ libgit2
	url := parseURL(gitUrl)

	// Set the clone path into ~/.tmbs/repos/reponame
	usr, err := user.Current()
	exitIfError(err, "Could not load current user.")
	path := usr.HomeDir + "/.tmbs/repos/" + name

	// Clone the repository
	repository, err := CloneRepository(url, path)
	exitIfError(err, "Could not clone repository.", url.String(), path)

	// Fetch HEAD
	err = FetchRepository(repository)
	exitIfError(err, "Could not fetch repository.")

	// If we make it this far, everything works.
	fmt.Println("Fetched repository", name, gitUrl)
	os.Exit(0)
}

// Taken from https://github.com/jwaldrip/git-get commit: be9ab16
// https://github.com/jwaldrip/git-get/blob/master/parsers.go
func parseURL(str string) *url.URL {
	u, err := url.Parse(str)
	exitIfError(err, "URL Invalid")
	if u.Scheme == "" {
		u = parseURL("ssh://" + str)
		parts := strings.Split(u.Host, ":")
		u.Host = parts[0]
		u.Path = path.Join(parts[1], u.Path)
	} else {
	}
	return u
}
