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
	Name         string
	FullPath     string
	Hidden       bool
	Suffix       string
	commentCount int
	codeCount    int
	blankCount   int
}

func (f *File) Count() int {
	return f.codeCount + f.blankCount
}

func (f *File) CommentCount() int {
	return f.commentCount
}
func (f *File) CodeCount() int {
	return f.codeCount
}
func (f *File) BlankCount() int {
	return f.blankCount
}

type Folder struct {
	Name     string
	FullPath string
	Hidden   bool
}

// 列出指定路径下的文件和文件夹
func List(pwd string) (fileList []*File, folderList []*Folder) {
	fileInfos, err := ioutil.ReadDir(pwd)
	if err != nil {
		log.Fatal(err)
	}

	for _, fi := range fileInfos {
		hidden := strings.HasPrefix(fi.Name(), ".")
		if fi.IsDir() {
			folderList = append(folderList,
				&Folder{
					Name:     fi.Name(),
					FullPath: path.Join(pwd, fi.Name()),
					Hidden:   hidden,
				})
		} else {
			fileList = append(fileList,
				&File{
					Name:     fi.Name(),
					FullPath: path.Join(pwd, fi.Name()),
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

func (f *File) CountLines(config *config.Config) (int, error) {
	if config.IgnoreHide && f.Hidden {
		return 0, nil
	}
	sf, err := os.Open(f.FullPath)
	b := make([]byte, 30)
	head := ""
	if _, err := sf.Read(b); err == nil {
		head = hex.EncodeToString(b)
		head = strings.ToUpper(head)
		for _, magicType := range Types {
			if strings.HasPrefix(head, magicType.Magic) {
				if config.Debug {
					fmt.Printf("识别到【%s】类型文件 %s,跳过\n", magicType.Name, f.FullPath)
				}
				return 0, nil
			}
		}

	}

	if err != nil {
		return 0, err
	}
	buf := bufio.NewReader(sf)
	for {
		bytes, _, err := buf.ReadLine()
		line := strings.TrimSpace(string(bytes))
		if len(line) != 0 {
			f.codeCount++
		} else {
			f.blankCount++
		}

		if err != nil {
			if err == io.EOF {
				if config.Debug {
					fmt.Printf("文件 %s \t 类型 %s \t 行数 %d \t魔法数 %s \n", f.FullPath, f.Suffix, f.Count(), head)
				}
				return f.Count(), nil
			}
			return 0, err
		}
	}

}
