package main

import (
	"log"
	"io/ioutil"
	"os"
	"strings"

	"github.com/marcusolsson/tui-go"
)

func goBackDir(currentDir string) string {
	currentDirSplit := strings.Split(currentDir, "/")
	return strings.Join(currentDirSplit[:len(currentDirSplit)-2], "/")+"/"
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


func main() {
	t := tui.NewTheme()
	normal := tui.Style{Bg: tui.ColorWhite, Fg: tui.ColorBlack}
	t.SetStyle("normal", normal)

	currentDir := "/home/josh/"
	lastIndex := 0

	files, err := ioutil.ReadDir(currentDir)
	if err != nil { log.Fatal(err) }

	// A simple label.
	okay := tui.NewLabel("Everything is fine.")

	// A list with some items selected.
	l := tui.NewList()
	l.SetFocused(true)
	l.AddItems(getFileNames(files)...)
	l.SetSelected(0)


	t.SetStyle("list.item.selected", tui.Style{Bg: tui.ColorBlue, Fg: tui.ColorWhite})

	root := tui.NewHBox(okay, tui.NewScrollArea(l))
	root.SetBorder(true)

	ui, err := tui.New(root)
	if err != nil {
		log.Fatal(err)
	}

	ui.SetTheme(t)
	ui.SetKeybinding("q", func() { ui.Quit() })

	ui.SetKeybinding("h", func() {
		currentDir = goBackDir(currentDir)
		files = listFolder(currentDir)
		l.RemoveItems()
		l.AddItems(getFileNames(files)...)
		l.SetSelected(0)
	})

	ui.SetKeybinding("l", func() {
		lastIndex = l.Selected()
		if files[lastIndex].IsDir() {
			currentDir = currentDir + l.SelectedItem() + "/"
			files = listFolder(currentDir)
			l.RemoveItems()
			l.AddItems(getFileNames(files)...)
			l.SetSelected(0)
		}
	})

	ui.SetKeybinding("k", func() {
		tempSel := l.Selected()
		if tempSel > 0 {
			l.Select(tempSel-1)
		}
	})

	ui.SetKeybinding("j", func() {
		tempSel := l.Selected()
		if tempSel < len(files)-1 {
			l.Select(tempSel+1)
		}
	})

	if err := ui.Run(); err != nil {
		log.Fatal(err)
	}
}
