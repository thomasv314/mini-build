package main

import (
	"fmt"
	"github.com/thomasv314/mini-build/tmbs"
	"os"
)

var configuration tmbs.Configuration

func main() {

	configuration = tmbs.LoadConfiguration()

	args := os.Args[1:]
	if len(args) == 0 {
		displayHelp()
	} else {
		switch args[0] {
		case "start":
			{
				// tmbs.go
				tmbs.Start()
			}
		case "setup":
			{
				// setup.go
				tmbs.Setup()
			}
		case "add":
			{
				// manager.go
				// TODO add error catching if args > 1 || 0
				if len(args) != 3 {
					fmt.Println("Usage: mini-build add <repo-name> <repo-url>")
				} else {
					tmbs.AddRepository(args[1], args[2])
				}
			}
		case "test":
			{

				tmbs.Setup()
				url := "https://tommyvyo@bitbucket.org/tommyvyo/mini-build.git"
				tmbs.AddRepository("ahous", url)

			}
		default:
			{
				displayHelp()
			}
		}
	}
}

func displayHelp() {
	fmt.Println("  Commands:")
	fmt.Println("\t start   - Starts the build server")
	fmt.Println("\t setup   - Initializes the tmbs directory")
	fmt.Println("\t add     - Add a repository")
	fmt.Println("\t help    - Display this message")
}
