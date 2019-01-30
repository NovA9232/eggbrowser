package fList

//import "fmt"

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

  return f
}

func (f *MainFList) updateAllFileLists() {
  f.LastFiles.UpdateList()
  f.NextFiles.UpdateList()
  f.CurrFiles.UpdateList()
}

func (f *MainFList) updateNextFilesAfterScroll() {
  if f.CurrFiles.Info[f.CurrFiles.ListObj.SelectedRow].IsDir() {
    f.NextFiles.Dir = f.CurrFiles.Dir+f.CurrFiles.Names[f.CurrFiles.ListObj.SelectedRow]+"/"
    f.NextFiles.UpdateList()
    //fmt.Println(f.NextFiles.Names)
  }
}

func (f *MainFList) ScrollDown() {
	f.CurrFiles.ListObj.ScrollDown()
  f.updateNextFilesAfterScroll()
}

func (f *MainFList) ScrollUp() {
	f.CurrFiles.ListObj.ScrollUp()
  f.updateNextFilesAfterScroll()
}

func (f *MainFList) GoLeft() {
  copyFileList(f.NextFiles, f.CurrFiles)
  copyFileList(f.CurrFiles, f.LastFiles)

  back, _ := goBackDir(f.CurrFiles.Dir)
  //println(back, "back")
  f.LastFiles.Dir = back
  f.updateAllFileLists()
}

func (f *MainFList) GoRight() {
  f.updateAllFileLists()
  //println(f.CurrFiles.Dir, f.NextFiles.Dir, f.CurrFiles.Names[0], f.NextFiles.Names[0])
}
