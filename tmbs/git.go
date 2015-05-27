package tmbs

import (
	"fmt"
	git "github.com/libgit2/git2go"
	"net/url"
	"os"
	"os/user"
	"path"
	"strconv"
	"strings"
	"time"
)

type WatchedRepository struct {
	Directory string
	Name      string
	Commits   []GitCommit
}

type GitCommit struct {
	Id             string
	RepositoryName string
	Author         string
	Message        string
	Timestamp      time.Time
	Status         string
	Source         string
}

func (commit GitCommit) Slug() string {
	return strconv.FormatInt(commit.Timestamp.Unix(), 10) + "-" + commit.Id
}

/*
*  Adds a repository to tmbs
 */
func AddRepository(name string, gitUrl string) {
	// Load the configuration to add to
	config, err := LoadConfiguration()

	if err != nil {
		exitIfError(err, "Cannot load configuration to add repository.")
	} else {

		var mainLog string = GetTmbsDirectory() + "/logs/" + "main.txt"

		writer := TestWriter{mainLog}
		writer.Setup()
		writer.Write("Adding repository " + name + " - " + gitUrl)

		// Parse the gitURL to work w/ libgit2
		url := parseURL(gitUrl)

		// Set the clone path into ~/.tmbs/repos/reponame
		usr, err := user.Current()
		exitIfError(err, "Could not load current user.")
		path := usr.HomeDir + "/.tmbs/repos/" + name

		checkOrCreateDir(GetTmbsDirectory()+"/logs/"+name, 0777)

		// Clone a bare repository to /repos/
		repository, err := CloneRepository(url, path, true, writer)

		fmt.Println("Repo", repository, "Err:", err)

		// The repository already exists because we tried to clone b4 without auth
		// so let's delete it and only exit if we still have problems (prob perms?)
		// err = os.RemoveAll(path)

		if err != nil {
			fmt.Println("Error adding repository", err)
			os.Exit(1)
		} else {
			//	exitIfError(err, "Could not clear repository path. Check permissions on", GetTmbsDirectory())

			fmt.Println("Cloned. Fetching repository now.")

			fmt.Println(repository)

			// If we make it this far, everything works.
			// Load the configuration, add the WatchedRepo, save configuration
			newRepo := WatchedRepository{Directory: path, Name: name}
			config.Repositories = append(config.Repositories, newRepo)

			SaveConfiguration(&config)
		}
	}
}

func parseURL(str string) *url.URL {
	u, _ := url.Parse(str)

	if u.Scheme == "" {
		u = parseURL("ssh://" + str)
		parts := strings.Split(u.Host, ":")
		u.Host = parts[0]
		u.Path = path.Join(parts[1], u.Path)
	}
	return u
}

var remoteCallbacks git.RemoteCallbacks

func repoCloneOptions(url *url.URL, bare bool, writer TestWriter) *git.CloneOptions {

	remoteCallbacks = git.RemoteCallbacks{
		SidebandProgressCallback: buildSidebandProgressCallback(writer),
		TransferProgressCallback: transferProgressCallback,
		CredentialsCallback:      buildCredCallback(url, writer),
		CertificateCheckCallback: buildCertCallback(url, writer),
	}

	return &git.CloneOptions{
		Bare:            bare,
		RemoteCallbacks: &remoteCallbacks,
	}
}

func CloneRepository(url *url.URL, path string, bare bool, writer TestWriter) (*git.Repository, error) {
	repo, err := git.Clone(url.String(), path, repoCloneOptions(url, bare, writer))
	return repo, err
}

func FetchRepository(repository *git.Repository) error {
	// list the remotes in the repository (origin by default)
	var (
		remotes []string
		err     error
	)

	remotes, err = repository.ListRemotes()

	fmt.Println("Remotes", remotes)
	if err != nil {
		fmt.Println("wtf?")
		fmt.Println(err)
		return err // error listing remotes
	} else {
		// lookup the first remote returned from the repo (origin)
		origin, err := repository.LookupRemote(remotes[0])
		origin.SetCallbacks(&remoteCallbacks)
		if err != nil {
			fmt.Println("Cannot lookup remote [0]")
			return err // error looking up remote
		} else {
			refspec := make([]string, 0)
			err = origin.Fetch(refspec, nil, "")
			if err != nil {
				fmt.Println("Error fetching the origin from the refspec")
				fmt.Println(refspec)
				return err // Error fetching
			} else {
				fmt.Println("Cannot fetch repo")
				return nil // Cool, we fetched the repo
			}
		}
	}
}
