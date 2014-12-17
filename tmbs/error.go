package tmbs

import (
	"fmt"
	"os"
)

func alertIfError(err error, messages ...string) {
	if err != nil {
		fmt.Println(err)
		fmt.Println("Alert Error:", messages)
	}
}

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
