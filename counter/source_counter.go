package counter

import (
	"encoding/json"
	"fmt"
	"sync"

	"sc/config"
)

type fileTypeCounter struct {
	Code    int // 代码行数
	Blank   int // 空行
	Comment int // 注释
}

// 统计器
type sourceCounter struct {
	fc map[string]*fileTypeCounter
}

var sc *sourceCounter
var once sync.Once

func init() {
	once.Do(func() {
		sc = &sourceCounter{
			fc: make(map[string]*fileTypeCounter),
		}
	})
	if config.Instant().Debug {
		fmt.Println("sourceCounter init done")
	}
}

func Instant() *sourceCounter {
	return sc
}

func (s *sourceCounter) Add(fileType, countType string, count int) {
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
	case "Code":
		fc.Code += count
	case "Blank":
		fc.Blank += count
	case "Comment":
		fc.Blank += count
	}
}

func (s *sourceCounter) Sum() *sourceCounter {

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

	s.fc["sum"] = sum

	return s

}
func (s *sourceCounter) String() string {
	bytes, err := json.Marshal(s.fc)
	if err != nil {
		fmt.Println(err.Error())
	}
	if config.Instant().Debug {
		fmt.Println(string(bytes))
	}
	return string(bytes)
}
