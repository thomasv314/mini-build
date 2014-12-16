package tmbs

import (
	"fmt"
	"os"
)

func exitWithMessage(messages ...string) {
	fmt.Println(messages)
	os.Exit(1)
}

func exitIfError(err error, message ...string) {
	if err != nil {
		fmt.Println(err)
		fmt.Println("Error:", message)
		os.Exit(1)
	}
}
