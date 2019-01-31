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

func (f *FileList) UpdateList() { // Also updates title to the directory
	if f.Dir == "" {
		f.Info = []os.FileInfo{}
		f.Names = []string{}
	} else {
		f.Info = listFolder(f.Dir)
		f.Names = getFileNames(f.Info)
	}
	f.ListObj.Rows = f.Names
	if f.ListObj.Title != "" {  // If it has a title, update it.
		f.ListObj.Title = f.Dir
	}
}

func copyFileList(dst, src *FileList) {
  dst.Dir = src.Dir
	dst.Names = make([]string, len(src.Names))
	dst.Info = make([]os.FileInfo, len(src.Info))
	copy(dst.Names, src.Names)
	copy(dst.Info, src.Info)
	dst.ListObj.SelectedRow = src.ListObj.SelectedRow
}
