package fList

import (
	"os"
	w "github.com/gizak/termui/widgets"
	ui "github.com/gizak/termui"
)

type FileList struct {
	ListObj *w.List
	Names []string
	Info []os.FileInfo
	Dir string
}

func NewFileList(dir string) *FileList {
	f := new(FileList)
	f.Dir = dir
	f.Info = listFolder(dir)
	f.Names = getFileNames(f.Info)

	f.ListObj = w.NewList()
	f.ListObj.TextStyle = ui.NewStyle(ui.ColorGreen)
	f.ListObj.WrapText = true
	f.ListObj.Rows = f.Names

	return f
}

func (f *FileList) UpdateList() {
	f.Info = listFolder(f.Dir)
	f.Names = getFileNames(f.Info)
	f.ListObj.Rows = f.Names
}
