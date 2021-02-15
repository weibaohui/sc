package file

// 文件夹
type Folder struct {
	Name     string
	FullPath string
	Hidden   bool
}

func (f *Folder) Execute() {
	listFolder(f)
}
