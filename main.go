package main

import (
	"log"
	"io/ioutil"
	"os"
	"strings"

	"github.com/marcusolsson/tui-go"
)


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

	ui.SetKeybinding("h", func() {
	})

	ui.SetKeybinding("l", func() {
	})

	ui.SetKeybinding("k", func() {
	})

	ui.SetKeybinding("j", func() {
	})

	if err := ui.Run(); err != nil {
		log.Fatal(err)
	}
}
