package file

import (
	"bufio"
	"encoding/hex"
	"fmt"
	"io"
	"os"
	"path"
	"strings"

	"sc/config"
	"sc/counter"
)

// File 文件
type File struct {
	Name     string
	FullPath string
	Hidden   bool
	Suffix   string
}

// 按文件统计
func countFileList(fileList []*File) {
	for _, f := range fileList {
		err := f.CountLines()
		if err != nil {
			// fmt.Println(err.Error())
			continue
		}
	}
}

// CountLines 统计行数
func (f *File) CountLines() error {
	sc := counter.GetInstance()
	cfg := config.GetInstance()

	if cfg.IgnoreHide && f.Hidden {
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
				if cfg.Debug {
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
			sc.Incr(ext, counter.CountTypeCode, 1)
		} else {
			sc.Incr(ext, counter.CountTypeBlank, 1)
		}
		codeCount++
		if err != nil {
			if err == io.EOF {
				if cfg.Debug {
					fmt.Printf("文件 %s \t  行数 %d \t魔法数 %s \n", f.FullPath, codeCount, head)
				}
				return nil
			}
			return err
		}
	}

}
