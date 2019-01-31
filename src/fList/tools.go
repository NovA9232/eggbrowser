package fList

import (
	"log"
	"os"
	"io/ioutil"
	"strings"
)

func findInList(item string, list []string) int {
	for i := 0; i < len(list); i++ {
		if list[i] == item {
			return i
		}
	}
	return -1
}

func goBackDir(currentDir string) (string, string) {
	currentDirSplit := strings.Split(currentDir, "/")
	return strings.Join(currentDirSplit[:len(currentDirSplit)-2], "/")+"/", currentDirSplit[len(currentDirSplit)-2] // Returns parent directory
}

func getFileNames(files []os.FileInfo) []string {
	var out []string
	for i := 0; i < len(files); i++ {
		out = append(out, files[i].Name())
	}
	return out
}

func listFolder(folderDir string) []os.FileInfo {
	files, err := ioutil.ReadDir(folderDir)
	if err != nil { log.Output(0, "ERROR: "+err.Error()) } // Probably just a permission error.
	return files
}
