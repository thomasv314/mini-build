package tmbs

import (
	"fmt"
	"os"
	"os/user"
)

func AddRepository(name string, gitUrl string) {

	fmt.Println("Adding", gitUrl, "to", name)

	usr, err := user.Current()
	if err != nil {
		panic(err)
	}

	path := usr.HomeDir + "/.tmbs/repos/" + name

	repository, err := CloneRepository(gitUrl, path)

	if err != nil {
		fmt.Println("Error cloning repository:", err)
		os.Exit(1)
	} else {
		fmt.Println("Cloned", gitUrl, "in to", path)
		err := FetchLatestRepository(repository)
		if err != nil {
			fmt.Println("Error fetching repository:", err)
			os.Exit(1)
		} else {
			fmt.Println("Fetched repository")
			os.Exit(0)
		}
	}

}
