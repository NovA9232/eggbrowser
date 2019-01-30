package fList

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
  f.NextFiles.Dir = f.CurrFiles.Dir+f.CurrFiles.Names[f.CurrFiles.ListObj.SelectedRow]
  if f.CurrFiles.Info[f.CurrFiles.ListObj.SelectedRow].IsDir() {
    f.NextFiles.Dir = f.NextFiles.Dir+"/"
    f.NextFiles.UpdateList()
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
  currDir := f.CurrFiles.Dir
  //name := f.CurrFiles.Names[f.CurrFiles.ListObj.SelectedRow]
  copyFileList(f.CurrFiles, f.LastFiles)
  back, _ := goBackDir(currDir)
  //println(back)
  f.LastFiles = NewFileList(back)
  f.updateAllFileLists()
}

func (f *MainFList) GoRight() {
  copyFileList(f.LastFiles, f.CurrFiles)
  currDir := f.CurrFiles.Dir
  name := f.CurrFiles.Names[f.CurrFiles.ListObj.SelectedRow]
  fileObj := f.CurrFiles.Info[f.CurrFiles.ListObj.SelectedRow]
  copyFileList(f.CurrFiles, f.NextFiles)
  f.CurrFiles.ListObj.SelectedRow = 0
  if fileObj.IsDir() {
    f.NextFiles = NewFileList(currDir+name+"/"+)
    println(currDir+name+"/")
  } else {
    f.NextFiles = NewFileList("/")
  }
  f.updateAllFileLists()
  //println(f.CurrFiles.Dir, f.NextFiles.Dir, f.CurrFiles.Names[0], f.NextFiles.Names[0])
}
