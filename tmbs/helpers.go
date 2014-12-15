package tmbs

import (
	"bufio"
	"fmt"
	"github.com/bgentry/speakeasy"
	"github.com/jwaldrip/tint"
	"os"
)

func getInput(question string) string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("%s: ", question)
	text, _ := reader.ReadString('\n')
	if text[len(text)-1] == '\n' {
		text = text[:len(text)-1]
	}
	return text
}

func getMaskedInput(question string) string {
	answer, err := speakeasy.Ask(question + tint.Colorize(" (masked)", tint.LightBlack) + ": ")
	exitIfError(err, "could not get masked input")
	return answer
}
