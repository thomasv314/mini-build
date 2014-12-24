package main

import (
	"fmt"
	"github.com/thomasv314/mini-build/tmbs"
	"log/syslog"
	"os"
)

var AppConfig tmbs.Configuration

func main() {

	logger, err := syslog.NewLogger(syslog.LOG_DEBUG, 0)
	logger.Print("Hello, log file!")

	AppConfig, err := tmbs.LoadConfiguration()

	// Start logging tmbs_log.go
	tmbs.StartLogs()

	if err != nil {
		fmt.Println("Could not find an existing application configuration. Running setup.")

		tmbs.Setup()

		fmt.Println("")
		displayHelp()

	} else {
		args := os.Args[1:]
		if len(args) == 0 {
			displayHelp()
		} else {
			switch args[0] {
			case "start":
				{
					// tmbs.go
					tmbs.Start(&AppConfig)
				}
			case "setup":
				{
					// setup.go
					tmbs.Setup()
				}
			case "add":
				{
					// add.go
					// TODO add error catching if args > 1 || 0
					if len(args) != 3 {
						fmt.Println("Usage: mini-build add <repo-name> <repo-url>")
					} else {
						tmbs.AddRepository(args[1], args[2])
					}
				}
			case "test-add-repo":
				{
					url := "git@bitbucket.org:tommyvyo/mini-build.git"
					//url := "http://tommyvyo@bitbucket.org/tommyvyo/mini-build.git"
					//					url := "https://tommyvyo@bitbucket.org/tommyvyo/mini-build.git"
					tmbs.AddRepository("mini-build", url)
				}
			case "test":
				{

					var config tmbs.Configuration = tmbs.Configuration{}

					err := tmbs.LoadJSONFile(tmbs.GetTmbsDirectory()+"/config.json", &config)

					if err != nil {
						fmt.Println("error", err)
					} else {
						fmt.Println(config)
					}
				}
			default:
				{
					displayHelp()
				}
			}
		}
	}
}

func displayHelp() {
	fmt.Println("** Thomas' Mini Build Server **") // Sup?
	fmt.Println("")
	fmt.Println("  Commands:")
	fmt.Println("\t start   - Starts the build server")
	fmt.Println("\t add     - Add a repository")
	fmt.Println("\t help    - Display this message")
}
