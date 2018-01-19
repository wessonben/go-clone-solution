package main

import (
	"fmt"
	"os"
	"path"
)

func main() {

	directory := getDirectory()

	fmt.Println(directory)
}

func getDirectory() string {
	directory := ""

	fmt.Print("Enter directory or leave blank for current: ")
	fmt.Scanln(&directory)

	if directory == "" {
		ex, _ := os.Executable()
		directory = path.Dir(ex)
	}

	return directory
}
