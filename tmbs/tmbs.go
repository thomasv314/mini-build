package tmbs

import (
	"fmt"
)

var config *Configuration

func Start(configuration *Configuration) {

	config = configuration

	StartHttpListener()
	StartBuildLauncher()

	for {
		// Run all goroutines til they press enter.
		var input string
		fmt.Scanln(&input)
		FakePush()
	}

}
