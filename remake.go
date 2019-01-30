package main

import (
	"strings"
	"log"
	"io/ioutil"
	"os"
	"github.com/rivo/tview"
)

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
	if err != nil { log.Fatal(err) }
	return files
}

func loadList(folderDir string, list *tview.List) {
	l := getFileNames(listFolder(folderDir))
	for _, element := range l {
		list.AddItem(element, "", 'a', func(){})
	}
}


type MainApp struct {
	rootGrid *tview.Grid
	listOfFiles *tview.List
	listOfLastFiles *tview.List
	listOfNextFiles *tview.List
	fileList []os.FileInfo
	lastFilesList []os.FileInfo
	nextFilesList []os.FileInfo
}


func main() {
	grid := tview.NewGrid().SetBorders(true).SetSize(0, 1, 0, 1)
	listOfLastFiles := tview.NewList().AddItem("Eef", "Secondary", ' ', func(){}).
																 AddItem("Bob", "Sec", ' ', func(){})

	listOfFiles := tview.NewList()//.AddItem("Egg", "Secondary", '1', func(){}).
																// AddItem("Beep", "m", '2', func(){})

	listOfNextFiles := tview.NewList().AddItem("Egg", "Secondary", '1', func(){}).
																 AddItem("Beep", "m", '2', func(){})

	loadList("/home/josh/", listOfFiles)

	grid.AddItem(listOfFiles, 0, 2, 1, 1, 10, 10, true)
	grid.AddItem(listOfLastFiles, 0, 1, 1, 1, 10, 10, true)
	grid.AddItem(listOfNextFiles, 0, 3, 1, 1, 10, 10, true)

	if err := tview.NewApplication().SetRoot(grid, true).Run(); err != nil {
		panic(err)
	}
}
