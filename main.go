package main

import (
	"fmt"
	"github.com/thomasv314/mini-build/tmbs"
	"os"
)

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		displayHelp()
	} else {
		switch args[0] {
		case "start":
			{
				tmbs.Start()
			}
		case "setup":
			{
				tmbs.Setup()
			}
		case "add":
			{
				tmbs.AddRepository(args)
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
