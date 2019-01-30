package fList

type MainFList struct {
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

func (f *MainFList) updateNextFilesAfterScroll() {
  f.NextFiles.Dir = f.CurrFiles.Dir+f.CurrFiles.Names[f.CurrFiles.ListObj.SelectedRow]
  if f.CurrFiles.Info[f.CurrFiles.ListObj.SelectedRow].IsDir() {
    f.NextFiles.Dir = f.NextFiles.Dir+"/"
    //print(f.NextFiles.Dir)
  }
  f.NextFiles.UpdateList()
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

}

func (f *MainFList) GoRight() {
  f.LastFiles = f.CurrFiles
  f.CurrFiles = f.NextFiles
}
