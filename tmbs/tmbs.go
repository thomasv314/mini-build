package tmbs

import (
	"fmt"
)

var config *Configuration

func Start(configuration *Configuration) {

	config = configuration

	fmt.Println("Starting HTTP Listener and Build Launcher")

	StartHttpListener()
	StartBuildLauncher()

	fmt.Println("Normally you'd hit enter to quit the application.")
	fmt.Println("Hit enter to send a FakePush signal")

	for {
		// Run all goroutines til they press enter.
		var input string
		fmt.Scanln(&input)
		FakePush()
	}

}
