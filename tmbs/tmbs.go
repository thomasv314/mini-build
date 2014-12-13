package tmbs

import (
	"fmt"
)

var configuration Configuration

func Start() {

	configuration = LoadConfiguration()

	StartHttpListener(configuration)
	StartBuildMaster(configuration)

	for {
		// Run all goroutines til they press enter.
		var input string
		fmt.Scanln(&input)
		FakePush()
	}

}
