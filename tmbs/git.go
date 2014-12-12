package main

import (
	"fmt"
	git "github.com/libgit2/git2go"
)

func main() {

	co := git.CloneOptions{}

	url := "https://tommyvyo@bitbucket.org/tommyvyo/mini-build.git"

	repo, err := git.Clone(url, "/home/thomas/code/whoa-mini-build", &co)

	if err != nil {
		fmt.Println("Error:", err)
	} else {

		fmt.Println("Cloned repository!")

		sl, err := repo.StatusList(nil)

		if err != nil {
			fmt.Println("Status error:", err)
		}

		fmt.Println(sl)
	}

}
