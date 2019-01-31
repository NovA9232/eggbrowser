package main

import (
	"log"
	//"fmt"
	ui "github.com/gizak/termui"

	"fList"
)

var (
	currentDir string

	fileView *fList.MainFList
	grid *ui.Grid
)


func setupWidgets() {
	fileView = fList.NewMainFList("/home/josh/programming/")

	grid = ui.NewGrid()
	termWidth, termHeight := ui.TerminalDimensions()
	grid.SetRect(0, 0, termWidth, termHeight)

	grid.Set(
		ui.NewRow(1.0, // One row
			ui.NewCol(1.0/3, fileView.LastFiles.ListObj),  // 3 columns
			ui.NewCol(1.0/3, fileView.CurrFiles.ListObj),
			ui.NewCol(1.0/3, fileView.NextFiles.ListObj),
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
			case "l", "<Right>":
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

func main() {
	if err := ui.Init(); err != nil {
		log.Fatalf("failed to initialize termui: %v", err)
	}
	defer ui.Close()

	setupWidgets()
	ui.Render(grid)

	// l := widgets.NewList()
	// l.Title = "List"
	// var list []string
	// for i := 0; i < 100; i++ {
	// 	list = append(list, fmt.Sprintf("Num: %d", i))
	// }
	//
	// l.Rows = list
	// l.TextStyle = ui.NewStyle(ui.ColorBlue)
	// l.WrapText = false
	// l.SetRect(0, 0, 400, 50)
	//
	// ui.Render(l)

	mainLoop()
}
