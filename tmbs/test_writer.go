package tmbs

import (
	"fmt"
	"os"
)

type TestWriter struct {
	FilePath string
}

func getLogPath(folder string, file string) string {
	return GetTmbsDirectory() + "/logs/" + folder + "/" + file + ".txt"
}

func (t TestWriter) Setup() {

	fmt.Println("Creating file", t.FilePath)

	file, err := os.Create(t.FilePath)
	defer file.Close()

	if err != nil {
		fmt.Println("Could not create file:", err)
	} else {
		fmt.Println("Created file", file)
	}
}

func (t TestWriter) Write(str string) {

	fmt.Println("Opening", t.FilePath)

	file, err := os.OpenFile(t.FilePath, os.O_APPEND|os.O_WRONLY, 0660)
	defer file.Close()

	if err != nil {
		fmt.Println("Cannot write. err:", err)
	} else {
		fmt.Println("[log]" + str)
		_, err2 := file.WriteString(str + "\n")
		if err2 != nil {
			fmt.Println("[error]", str)
			fmt.Println("[error]", err2)
		}
	}
}
