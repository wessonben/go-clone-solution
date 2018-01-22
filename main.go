package main

import (
	"fmt"
	"io/ioutil"
)

func main() {

	indir := getDirectory("Enter source directory")
	outdir := getDirectory("Enter output directory")
	replace := getText("Enter text to replace")
	inject := getText("Enter text to inject")

	cloneSolution(indir, outdir, replace, inject)
}

func getDirectory(prompt string) string {
	directory := ""
	fmt.Print(prompt)
	fmt.Scanln(&directory)
	return directory
}

func getText(prompt string) string {
	text := ""
	fmt.Print(prompt)
	fmt.Scanln(&text)
	return text
}

func cloneSolution(indir string, outdir string, replace string, inject string) {

	// any solution level stuff before...
	cloneFolder(indir, outdir, replace, inject)
}

func cloneFolder(sourcedir string, targetdir string, replace string, inject string) {

	directories, files := getDirectoryItems(sourcedir)

	for _, name := range directories {
		sourcesubdir := sourcedir + "\\" + name
		targetsubdir := "" // need to substitute text
		// create the new targetsubdir
		cloneFolder(sourcesubdir, targetsubdir, replace, inject)
	}

	for _, name := range files {
		sourcefile := sourcedir + "\\" + name
		cloneFile(sourcefile, targetdir, replace, inject)
	}
}

func cloneFile(sourcefile string, targetDir string, replace string, inject string) {
	// copy file to target directory
	// replace text in file name
	// replace text in file content
}

func getDirectoryItems(path string) ([]string, []string) {
	directories := make([]string, 0)
	files := make([]string, 0)
	fileInfos, _ := ioutil.ReadDir(path)
	for _, fileInfo := range fileInfos {
		if fileInfo.IsDir() {
			directories = append(directories, fileInfo.Name())
		} else {
			files = append(files, fileInfo.Name())
		}
	}
	return directories, files
}