package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
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
	fmt.Print(prompt + ": ")
	fmt.Scanln(&directory)
	return directory
}

func getText(prompt string) string {
	text := ""
	fmt.Print(prompt + ": ")
	fmt.Scanln(&text)
	return text
}

func cloneSolution(indir string, outdir string, replace string, inject string) {

	// ensure the output directory is empty

	// begin cloning
	cloneFolder(indir, outdir, replace, inject)
}

func cloneFolder(sourcedir string, targetdir string, replace string, inject string) {

	directories, files := getDirectoryItems(sourcedir)

	for _, name := range directories {
		// ignore the VS build output directories
		if (name != "bin" && name != "obj") {
			sourcesubdir := sourcedir + "\\" + name
			targetsubdir := targetdir + "\\" + name
			targetsubdir = strings.Replace(targetsubdir, replace, inject, -1)
			os.Mkdir(targetsubdir, os.ModeDir)
			cloneFolder(sourcesubdir, targetsubdir, replace, inject)
		}
	}

	for _, name := range files {
		if (!strings.HasSuffix(name, ".vspscc") && !strings.HasSuffix(name, ".user")) {
			cloneFile(sourcedir, targetdir, name, replace, inject)
		}
	}
}

func cloneFile(sourcedir string, targetdir string, filename string, replace string, inject string) {

	// construnct full file paths
	sourcefile := sourcedir + "\\" + filename
	targetfile := targetdir + "\\" + filename
	targetfile = strings.Replace(targetfile, replace, inject, -1)

	// read the content of the source file and convert to string
	contentbytes, _ := ioutil.ReadFile(sourcefile)
	contentstring := string(contentbytes)

	// do string substitution in file content and convert back to bytes
	contentstring = strings.Replace(contentstring, replace, inject, -1)
	contentbytes = []byte(contentstring)

	// create the new destination file
	ioutil.WriteFile(targetfile, contentbytes, os.ModePerm)
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