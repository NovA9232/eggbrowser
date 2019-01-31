package main

import (
	"log"
	//"fmt"
	"os"
	ui "github.com/gizak/termui"

	"fList"
)

var (
	currentDir string

	fileView *fList.MainFList
	grid *ui.Grid

	BordStyle ui.Style
	SelectStyleFolder ui.Style
)


func setupWidgets() {
	fileView = fList.NewMainFList("/home/josh/")

	grid = ui.NewGrid()
	termWidth, termHeight := ui.TerminalDimensions()
	grid.SetRect(0, 0, termWidth, termHeight)

	grid.Set(
		ui.NewRow(1.0, // One row
			ui.NewCol(2.0/10, fileView.LastFiles.ListObj),  // 3 columns
			ui.NewCol(3.0/10, fileView.CurrFiles.ListObj),
			ui.NewCol(5.0/10, fileView.NextFiles.ListObj),
		),
	)
}

func mainLoop() {
	uiEvents := ui.PollEvents()
	for {
		e := <-uiEvents
		//fmt.Println(e)
		switch e.ID {
			case "q", "<C-c>":
				return
			case "j", "<Down>":
				fileView.ScrollDown()
				ui.Render(grid)
			case "k", "<Up>":
				fileView.ScrollUp()
				ui.Render(grid)
			case "h", "<Left>":
				fileView.GoLeft()
				ui.Render(grid)
			case "l", "<Right>", "<Enter>":
				fileView.GoRight()
				ui.Render(grid)
			case "<PageUp>":
				fileView.PageUp()
				ui.Render(grid)
			case "<PageDown>":
				fileView.PageDown()
				ui.Render(grid)
		}
	}
}

func setupLog(dest string) *os.File {
	f, err := os.OpenFile(dest, os.O_RDWR | os.O_CREATE | os.O_APPEND, 0666)
	if err != nil {
	    log.Fatalf("error opening file: %v", err)
	}

	log.SetOutput(f)
	return f
}

func main() {
	logFile := setupLog("out.log")
	defer logFile.Close()

	BordStyle := ui.NewStyle(ui.Color(6))
	SelectStyleFolder := ui.NewStyle(ui.Color(6), ui.Color(7), ui.ModifierReverse)

	fList.BordStyle = BordStyle
	fList.SelectStyleFolder = SelectStyleFolder

	if err := ui.Init(); err != nil {
		log.Fatalf("failed to initialize termui: %v", err)
	}
	defer ui.Close()

	setupWidgets()
	ui.Render(grid)

	mainLoop()
}
