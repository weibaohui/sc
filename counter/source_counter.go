package counter

import (
	"encoding/json"
	"fmt"
	"sync"

	"github.com/weibaohui/sc/config"
)

var (
	sc               *SourceCounter
	once             sync.Once
	CountTypeCode    = "Code"    // type code
	CountTypeBlank   = "Blank"   // blank
	CountTypeComment = "Comment" // comment
)

type fileTypeCounter struct {
	Code    int    // 代码行数
	Blank   int    // 空行
	Comment int    // 注释
	Name    string // 类型名称
}

// SourceCounter  contains the file type,and it's count
type SourceCounter struct {
	FileTypeCounterList []*fileTypeCounter
}

func init() {
	once.Do(func() {
		sc = &SourceCounter{
			FileTypeCounterList: []*fileTypeCounter{},
		}
	})
	if config.GetInstance().Debug {
		fmt.Println("SourceCounter init done")
	}
}

// GetInstance get an Instance
func GetInstance() *SourceCounter {
	return sc
}

// Incr increase a counter
func (s *SourceCounter) Incr(fileType, countType string, count int) {
	var fc *fileTypeCounter
	for _, counter := range s.FileTypeCounterList {
		if counter.Name == fileType {
			fc = counter
		}
	}

	if fc == nil {
		fc = &fileTypeCounter{
			Code:    0,
			Blank:   0,
			Comment: 0,
			Name:    fileType,
		}
		s.FileTypeCounterList = append(s.FileTypeCounterList, fc)
	}
	switch countType {
	case CountTypeCode:
		fc.Code += count
	case CountTypeBlank:
		fc.Blank += count
	case CountTypeComment:
		fc.Blank += count
	}
}

// Sum sum all the count
func (s *SourceCounter) Sum() *SourceCounter {

	sum := &fileTypeCounter{
		Code:    0,
		Blank:   0,
		Comment: 0,
		Name:    "ALL",
	}
	for _, c := range s.FileTypeCounterList {
		sum.Code += c.Code
		sum.Blank += c.Blank
		sum.Comment += c.Comment
	}

	s.FileTypeCounterList = append(s.FileTypeCounterList, sum)
	return s

}

// String
func (s *SourceCounter) String() string {
	bytes, err := json.Marshal(s.FileTypeCounterList)
	if err != nil {
		fmt.Println(err.Error())
	}
	if config.GetInstance().Debug {
		fmt.Printf("统计数据:%s", string(bytes))
	}
	return string(bytes)
}

func (s *SourceCounter) Result() []*fileTypeCounter {
	return s.FileTypeCounterList
}
