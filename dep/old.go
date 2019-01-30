package eh

import (
	"log"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"

	"github.com/marcusolsson/tui-go"
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

func findInList(item string, list []string) int {
	for i := 0; i < len(list); i++ {
		if list[i] == item {
			return i
		}
	}
	return -1
}

func goLeft(currentDir string, l *tui.List) (string, []os.FileInfo) {
	parent := ""
	currentDir, parent = goBackDir(currentDir)
	files := listFolder(currentDir)
	fileNames := getFileNames(files)
	l.RemoveItems()
	l.AddItems(fileNames...)
	l.SetSelected(findInList(parent, fileNames))

	return currentDir, files
}

func goRight(currentDir string, files []os.FileInfo, l *tui.List) (string, []os.FileInfo) {
	lastIndex := l.Selected()
	if files[lastIndex].IsDir() {
		currentDir = currentDir + l.SelectedItem() + "/"
		files = listFolder(currentDir)
		l.RemoveItems()
		l.AddItems(getFileNames(files)...)
		l.SetSelected(0)
	} else {
		_, err := exec.Command("/bin/bash", "-c", "xdg-open '"+currentDir+files[lastIndex].Name()+"'").Output()
		if err != nil { log.Fatal(err) }
	}
	return currentDir, files
}

func goUp(l *tui.List) {
	tempSel := l.Selected()
	if tempSel > 0 {
		l.Select(tempSel-1)
	}
}

func goDown(filesLen int, l *tui.List) {
	tempSel := l.Selected()
	if tempSel < filesLen-1 {
		l.Select(tempSel+1)
	}
}



func main() {
	t := tui.NewTheme()
	normal := tui.Style{Bg: tui.ColorWhite, Fg: tui.ColorBlack}
	t.SetStyle("normal", normal)

	currentDir := "/home/josh/Important images/"

	files, err := ioutil.ReadDir(currentDir)
	if err != nil { log.Fatal(err) }

	// A simple label.

	// A list with some items selected.
	l := tui.NewList()
	l.SetFocused(true)
	l.AddItems(getFileNames(files)...)
	l.SetSelected(0)


	t.SetStyle("list.item.selected", tui.Style{Bg: tui.ColorBlue, Fg: tui.ColorWhite})

	box := tui.NewHBox(l)
	box.SetBorder(true)
	root := tui.NewScrollArea(box)

	ui, err := tui.New(root)
	if err != nil {
		log.Fatal(err)
	}

	ui.SetTheme(t)
	ui.SetKeybinding("q", func() { ui.Quit() })

	ui.SetKeybinding("h", func() { currentDir, files = goLeft(currentDir, l) })
	ui.SetKeybinding("Left", func() { currentDir, files = goLeft(currentDir, l) })

	ui.SetKeybinding("l", func() { currentDir, files = goRight(currentDir, files, l) })
	ui.SetKeybinding("Right", func() { currentDir, files = goRight(currentDir, files, l) })

	ui.SetKeybinding("k", func() { goUp(l) })
	ui.SetKeybinding("Up", func() { goUp(l) })

	ui.SetKeybinding("j", func() { goDown(len(files), l) })
	ui.SetKeybinding("Down", func() { goDown(len(files), l) })

	if err := ui.Run(); err != nil {
		log.Fatal(err)
	}
}
