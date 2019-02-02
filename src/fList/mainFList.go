package fList

import (
  "log"
  "os/exec"
)

type MainFList struct {  // More of a manger for the three FileLists
  LastFiles *FileList
  NextFiles *FileList
  CurrFiles *FileList
}

func NewMainFList(dir string) *MainFList {
  f := new(MainFList)
  f.CurrFiles = NewFileList(dir)
  back, name := goBackDir(dir)
  f.LastFiles = NewFileList(back)
  f.LastFiles.ListObj.SelectedRow = uint(findInList(name, f.LastFiles.Names))
  f.NextFiles = NewFileList(dir+f.CurrFiles.Names[0]+"/")

  f.CurrFiles.ListObj.Title = f.CurrFiles.Dir

  return f
}

func (f *MainFList) updateAllFileLists() {
  f.LastFiles.UpdateList()
  f.NextFiles.UpdateList()
  f.CurrFiles.UpdateList()
}

func (f *MainFList) updateNextFiles() {
  //log.Output(0, fmt.Sprintf("File name nextFiles: %s, Is a dir: %s", f.CurrFiles.Info[f.CurrFiles.ListObj.SelectedRow].Name(), f.CurrFiles.Info[f.CurrFiles.ListObj.SelectedRow].IsDir()))
  if f.CurrFiles.Info[f.CurrFiles.ListObj.SelectedRow].IsDir() {
    f.NextFiles.Dir = f.CurrFiles.Dir+f.CurrFiles.Names[f.CurrFiles.ListObj.SelectedRow]+"/"
    f.NextFiles.UpdateList()
    f.NextFiles.ListObj.SelectedRow = 0
  } else {
    f.NextFiles.Dir = ""
    f.NextFiles.UpdateList()
  }
}

func (f *MainFList) ScrollDown() {
  if len(f.CurrFiles.Names) > 0 {
    f.CurrFiles.ListObj.ScrollDown()
    f.updateNextFiles()
  }
}

func (f *MainFList) ScrollUp() {
  if len(f.CurrFiles.Names) > 0 {
  	f.CurrFiles.ListObj.ScrollUp()
    f.updateNextFiles()
  }
}

func (f *MainFList) PageUp() {
  if len(f.CurrFiles.Names) > 0 {
    f.CurrFiles.ListObj.PageUp()
    f.updateNextFiles()
  }
}

func (f *MainFList) PageDown() {
  if len(f.CurrFiles.Names) > 0 {
    f.CurrFiles.ListObj.PageDown()
    f.updateNextFiles()
  }
}

func (f *MainFList) GoLeft() {
	if f.CurrFiles.Dir != "/" {
		copyFileList(f.NextFiles, f.CurrFiles)
		copyFileList(f.CurrFiles, f.LastFiles)

		back, name := goBackDir(f.CurrFiles.Dir)
		f.LastFiles.Dir = back
		f.updateAllFileLists()
		f.LastFiles.ListObj.SelectedRow = uint(findInList(name, f.LastFiles.Names))
	}
}

func (f *MainFList) GoRight() {
  if f.CurrFiles.Info[f.CurrFiles.ListObj.SelectedRow].IsDir() {
    copyFileList(f.LastFiles, f.CurrFiles)
    copyFileList(f.CurrFiles, f.NextFiles)
    if len(f.CurrFiles.Names) > 0 {
      f.updateNextFiles()
    }

    f.updateAllFileLists()
  } else {
		go func() {
			_, err := exec.Command("/bin/bash", "-c", "xdg-open '"+f.CurrFiles.Dir+f.CurrFiles.Names[f.CurrFiles.ListObj.SelectedRow]+"'").Output()
			if err != nil { log.Fatal(err) }
		}()
  }
}
