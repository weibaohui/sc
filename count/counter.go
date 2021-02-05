package count

import (
	"encoding/json"
	"fmt"

	"sc/config"
	"sc/file"
)

type Counter struct {
	Code    int // 代码行数
	Blank   int // 空行
	Comment int // 注释
}

func (c *Counter) Execute(f *file.Folder, config *config.Config) {
	c.countFolder(f, config)
	bytes, _ := json.Marshal(c)
	fmt.Println(string(bytes))
}

// 按文件夹统计
func (c *Counter) countFolder(f *file.Folder, config *config.Config) {
	if config.IgnoreHide && f.Hidden {
		return
	}
	fileList, folderList := f.List(f.FullPath, config)
	// 计算文件里的代码行数
	c.countFile(fileList, config)

	// 按文件夹迭代，计算文件里的代码行数
	for _, folder := range folderList {
		c.countFolder(folder, config)
	}

}

// 按文件统计
func (c *Counter) countFile(fileList []*file.File, config *config.Config) {
	for _, f := range fileList {
		codeCount, blankCount, commentCount, err := f.CountLines(config)
		if err != nil {
			// fmt.Println(err.Error())
			continue
		}
		c.Code += codeCount
		c.Blank += blankCount
		c.Comment += commentCount
	}
}
