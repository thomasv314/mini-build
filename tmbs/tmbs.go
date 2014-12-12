package tmbs

import (
	"fmt"
)

func StartBuildServer() {

	configuration := LoadConfiguration()

	StartPushListener(configuration)
	StartBuildMaster(configuration)

	for {
		// Run all goroutines til they press enter.
		var input string
		fmt.Scanln(&input)
		FakePush()
	}

}
