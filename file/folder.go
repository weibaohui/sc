package file

import (
	"io/ioutil"
	"log"
	"path"
	"strings"

	"github.com/weibaohui/sc/config"
)

// Folder 文件夹
type Folder struct {
	Name     string
	FullPath string
	Hidden   bool
}

// Execute Start to execute
func (f *Folder) Execute() {
	countFolderList(f)
}

// 按文件夹统计
func countFolderList(f *Folder) {
	config := config.GetInstance()
	if config.IgnoreHide && f.Hidden {
		return
	}
	fileList, folderList := f.List(f.FullPath)
	// 计算文件里的代码行数
	countFileList(fileList)

	// 按文件夹迭代，计算文件里的代码行数
	for _, folder := range folderList {
		countFolderList(folder)
	}
}

// List 列出指定路径下的文件和文件夹
func (f *Folder) List(fullPath string) (fileList []*File, folderList []*Folder) {
	fileInfos, err := ioutil.ReadDir(fullPath)
	if err != nil {
		log.Fatal(err)
	}

	cfg := config.GetInstance()

	for _, fi := range fileInfos {
		hidden := strings.HasPrefix(fi.Name(), ".")
		if fi.IsDir() {
			// 文件夹是否在排除列表
			for _, s := range cfg.Exclude {
				if s == fi.Name() {
					return
				}
			}
			folderList = append(folderList,
				&Folder{
					Name:     fi.Name(),
					FullPath: path.Join(fullPath, fi.Name()),
					Hidden:   hidden,
				})
		} else {
			fileList = append(fileList,
				&File{
					Name:     fi.Name(),
					FullPath: path.Join(fullPath, fi.Name()),
					Hidden:   hidden,
					Suffix:   path.Ext(fi.Name()),
				})
		}
	}

	return
}
