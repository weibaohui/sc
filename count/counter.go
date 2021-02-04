package count

import (
	"fmt"

	"sc/config"
	"sc/file"
)

type Counter struct {
}

func (c *Counter) Execute(f *file.Folder, config *config.Config) {
	count := countFolder(f, config)
	fmt.Printf("\n共计【 %d 】行", count)
}

// 按文件夹统计
func countFolder(f *file.Folder, config *config.Config) int {
	if config.IgnoreHide && f.Hidden {
		return 0
	}
	fileList, folderList := file.List(f.FullPath)

	count := countFile(fileList, config)

	for _, folder := range folderList {
		count += countFolder(folder, config)
	}

	return count
}

// 按文件统计
func countFile(fileList []*file.File, config *config.Config) (finalCount int) {
	for _, f := range fileList {
		count, err := f.CountLines(config)
		if err != nil {
			fmt.Println(err.Error())
			continue
		}
		finalCount += count
	}
	return finalCount
}
