package file

import (
	"bufio"
	"encoding/hex"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path"
	"strings"

	"sc/count"
)

// 文件
type File struct {
	Name     string
	FullPath string
	Hidden   bool
	Suffix   string
}

// 按文件夹统计
func listFolder(f *Folder) {
	config := count.GetSourceCounter().Config()
	if config.IgnoreHide && f.Hidden {
		return
	}
	fileList, folderList := f.List(f.FullPath)
	// 计算文件里的代码行数
	countFile(fileList)

	// 按文件夹迭代，计算文件里的代码行数
	for _, folder := range folderList {
		listFolder(folder)
	}

}

// 按文件统计
func countFile(fileList []*File) {
	for _, f := range fileList {
		err := f.CountLines()
		if err != nil {
			// fmt.Println(err.Error())
			continue
		}
	}
}

// 列出指定路径下的文件和文件夹
func (f *Folder) List(fullPath string) (fileList []*File, folderList []*Folder) {
	fileInfos, err := ioutil.ReadDir(fullPath)
	if err != nil {
		log.Fatal(err)
	}

	c := count.GetSourceCounter().Config()

	for _, fi := range fileInfos {
		hidden := strings.HasPrefix(fi.Name(), ".")
		if fi.IsDir() {
			// 文件夹是否在排除列表
			for _, s := range c.Exclude {
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

// 统计行数
func (f *File) CountLines() error {
	counter := count.GetSourceCounter()
	config := counter.Config()

	if config.IgnoreHide && f.Hidden {
		return nil
	}
	ext := path.Ext(f.FullPath)

	sf, err := os.Open(f.FullPath)
	defer sf.Close()
	b := make([]byte, 30)
	head := ""
	if _, err = sf.Read(b); err == nil {

		head = hex.EncodeToString(b)
		head = strings.ToUpper(head)
		// fmt.Printf("识别 %s 文件%s \n", f.FullPath, string(b))

		for _, magicType := range Types {
			if strings.HasPrefix(head, magicType.Magic) {
				if config.Debug {
					// fmt.Printf("识别到【%s】类型文件%s,跳过=%t\n", magicType.Name, f.FullPath, magicType.Skip)
				}
				return nil
			}
		}
	}

	if err != nil {
		return nil
	}
	buf := bufio.NewReader(sf)
	codeCount := 0
	for {
		bytes, _, err := buf.ReadLine()
		line := strings.TrimSpace(string(bytes))
		if len(line) != 0 {
			counter.Add(ext, "Code", 1)

		} else {
			counter.Add(ext, "Blank", 1)
		}
		codeCount++
		if err != nil {
			if err == io.EOF {
				if config.Debug {
					fmt.Printf("文件 %s \t  行数 %d \t魔法数 %s \n", f.FullPath, codeCount, head)
				}
				return nil
			}
			return err
		}
	}

}
