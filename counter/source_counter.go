package counter

import (
	"encoding/json"
	"fmt"
	"sync"

	"sc/config"
)

var (
	sc               *SourceCounter
	once             sync.Once
	CountTypeCode    = "Code"
	CountTypeBlank   = "Blank"
	CountTypeComment = "Comment"
	CountTypeSum     = "Sum"
)

type fileTypeCounter struct {
	Code    int // 代码行数
	Blank   int // 空行
	Comment int // 注释
}

// SourceCounter
type SourceCounter struct {
	fc map[string]*fileTypeCounter
}

func init() {
	once.Do(func() {
		sc = &SourceCounter{
			fc: make(map[string]*fileTypeCounter),
		}
	})
	if config.GetInstance().Debug {
		fmt.Println("SourceCounter init done")
	}
}

// GetInstance
func GetInstance() *SourceCounter {
	return sc
}

// Incr
func (s *SourceCounter) Incr(fileType, countType string, count int) {
	fc := s.fc[fileType]
	if fc == nil {
		fc = &fileTypeCounter{
			Code:    0,
			Blank:   0,
			Comment: 0,
		}
		s.fc[fileType] = fc
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

// Sum
func (s *SourceCounter) Sum() *SourceCounter {

	sum := &fileTypeCounter{
		Code:    0,
		Blank:   0,
		Comment: 0,
	}
	for _, c := range s.fc {
		sum.Code += c.Code
		sum.Blank += c.Blank
		sum.Comment += c.Comment
	}

	s.fc[CountTypeSum] = sum

	return s

}

// String
func (s *SourceCounter) String() string {
	bytes, err := json.Marshal(s.fc)
	if err != nil {
		fmt.Println(err.Error())
	}
	if config.GetInstance().Debug {
		fmt.Printf("统计数据:%s", string(bytes))
	}
	return string(bytes)
}
