package main

import (
	"log"
	"io/ioutil"
	"os"
	"strings"

	"github.com/marcusolsson/tui-go"
)

var (
	currentDir string = "/home/josh/test1/"
	currentFiles = []os.FileInfo{}
	currentFileNames = []string{}
)


func getFileNames(files []os.FileInfo) []string {
	var out []string
	for _, file := range files {
		out = append(out, file.Name())
	}

	return out
}

func changeDir(file os.FileInfo) {
	if file.IsDir() {
		files, err := ioutil.ReadDir(currentDir+file.Name()+"/")
		if err != nil { log.Fatal(err) }
		currentFiles = files
		currentDir = currentDir+file.Name()+"/"
		currentFileNames = getFileNames(files)
	}
}

func getBackDir(dir string) string {
	split := strings.Split(currentDir, "/")
	return strings.Join(split[:len(split)-2], "/")+"/"
}


func goBackDir() {
	currentDir = getBackDir(currentDir)
	files, err := ioutil.ReadDir(currentDir)
	if err != nil { log.Fatal(err) }
	currentFiles = files
	currentFileNames = getFileNames(currentFiles)
}


func main() {
	t := tui.NewTheme()
	normal := tui.Style{Bg: tui.ColorWhite, Fg: tui.ColorBlack}
	t.SetStyle("normal", normal)


	files, err := ioutil.ReadDir(currentDir)
	if err != nil { log.Fatal(err) }
	currentFiles = files
	currentFileNames = getFileNames(currentFiles)

	// A simple label.
	okay := tui.NewLabel("Everything is fine.")

	// A list with some items selected.
	l := tui.NewList()
	l.SetFocused(true)
	l.AddItems(currentFileNames...)
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

	ui.SetKeybinding("Left", func() {
		goBackDir()
		l.RemoveItems()
		l.AddItems(currentFileNames...)
		l.SetSelected(0)
	})

	ui.SetKeybinding("Right", func() {
		changeDir(currentFiles[l.Selected()])
		l.RemoveItems()
		l.AddItems(currentFileNames...)
		l.SetSelected(0)
	})

	ui.SetKeybinding("Up", func() {
		if l.Selected() > 0 {
			l.Select(l.Selected()-1)
		}
	})

	ui.SetKeybinding("Down", func() {
		if l.Selected() < len(currentFileNames)-1 {
			l.Select(l.Selected()+1)
		}
	})

	if err := ui.Run(); err != nil {
		log.Fatal(err)
	}
}
