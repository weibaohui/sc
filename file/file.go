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

	"sc/config"
)

type File struct {
	Name     string
	FullPath string
	Hidden   bool
	Suffix   string
}

type Folder struct {
	Name     string
	FullPath string
	Hidden   bool
}

// 列出指定路径下的文件和文件夹
func (f *Folder) List(fullPath string, c *config.Config) (fileList []*File, folderList []*Folder) {
	fileInfos, err := ioutil.ReadDir(fullPath)
	if err != nil {
		log.Fatal(err)
	}

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
					Suffix:   ext(fi.Name()),
				})
		}
	}

	return
}

// 取后缀
func ext(pwd string) string {
	tmp := strings.Split(pwd, ".")
	if len(tmp) >= 2 && tmp[0] != "" {
		return strings.ToLower(tmp[len(tmp)-1])
	}
	return ""
}

func (f *File) CountLines(config *config.Config) (codeCount, blankCount, commentCount int, err error) {

	if config.IgnoreHide && f.Hidden {
		return
	}
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
				return
			}
		}
	}

	if err != nil {
		return
	}
	buf := bufio.NewReader(sf)
	for {
		bytes, _, err := buf.ReadLine()
		line := strings.TrimSpace(string(bytes))
		if len(line) != 0 {
			codeCount++
		} else {
			blankCount++
		}

		if err != nil {
			if err == io.EOF {
				if config.Debug {
					fmt.Printf("文件 %s \t  行数 %d \t魔法数 %s \n", f.FullPath, codeCount, head)
				}
				return codeCount, blankCount, commentCount, nil
			}
			return 0, 0, 0, err
		}
	}

}
