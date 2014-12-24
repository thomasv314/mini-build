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

	// Load the configuration to add to
	config, err := LoadConfiguration()

	if err != nil {
		exitIfError(err, "Cannot load configuration to add repository.")
	} else {

		// Parse the gitURL to work w/ libgit2
		url := parseURL(gitUrl)

		// Set the clone path into ~/.tmbs/repos/reponame
		usr, err := user.Current()
		exitIfError(err, "Could not load current user.")
		path := usr.HomeDir + "/.tmbs/repos/" + name

		// Clone a bare repository to /repos/
		repository, err := CloneRepository(url, path, true)
		exitIfError(err, "Could not clone repository.")

		// The repository already exists because we tried to clone b4 without auth
		// so let's delete it and only exit if we still have problems (prob perms?)
		//	err = os.RemoveAll(path)
		exitIfError(err, "Could not clear repository path. Check permissions on", GetTmbsDirectory())

		fmt.Println("Cloned. Fetching repository now.")

		fmt.Println(repository)

		// Fetch HEAD
		err = FetchRepository(repository)
		exitIfError(err, "Could not fetch repository.")

		// If we make it this far, everything works.
		// Load the configuration, add the WatchedRepo, save configuration
		newRepo := WatchedRepository{Directory: path, Name: name}
		config.Repositories = append(config.Repositories, newRepo)

		SaveConfiguration(&config)

		os.Exit(0)

	}
}

// Taken from https://github.com/jwaldrip/git-get commit: be9ab16
// https://github.com/jwaldrip/git-get/blob/master/parsers.go
func parseURL(str string) *url.URL {
	u, err := url.Parse(str)
	exitIfError(err, "URL Invalid")
	fmt.Println("Parsed:", str)
	if u.Scheme == "" {
		u = parseURL("git://" + str)
		parts := strings.Split(u.Host, ":")
		u.Host = parts[0]
		u.Path = path.Join(parts[1], u.Path)
	}
	return u
}
