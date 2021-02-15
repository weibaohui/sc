package count

import (
	"encoding/json"
	"fmt"
	"sync"

	"sc/config"
)

type FileTypeCounter struct {
	Code    int // 代码行数
	Blank   int // 空行
	Comment int // 注释
}

// 统计器
type SourceCounter struct {
	FileTypeCounter map[string]*FileTypeCounter
	Code            int // 代码行数
	Blank           int // 空行
	Comment         int // 注释
	config          *config.Config
}

var sourceCounter *SourceCounter
var once sync.Once

func init() {
	once.Do(func() {
		sourceCounter = &SourceCounter{
			FileTypeCounter: make(map[string]*FileTypeCounter),
		}
	})
	fmt.Println("SourceCounter init done")
}
func NewSourceCounter(config *config.Config) *SourceCounter {
	sourceCounter.config = config
	return sourceCounter
}
func GetSourceCounter() *SourceCounter {
	return sourceCounter
}
func (s *SourceCounter) Config() *config.Config {
	return s.config
}
func (s *SourceCounter) Add(fileType, countType string, count int) {
	fc := s.FileTypeCounter[fileType]
	if fc == nil {
		fc = &FileTypeCounter{
			Code:    0,
			Blank:   0,
			Comment: 0,
		}
		s.FileTypeCounter[fileType] = fc
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

func (s *SourceCounter) Sum() *SourceCounter {

	sum := &FileTypeCounter{
		Code:    0,
		Blank:   0,
		Comment: 0,
	}
	for _, c := range s.FileTypeCounter {
		sum.Code += c.Code
		sum.Blank += c.Blank
		sum.Comment += c.Comment
	}

	s.FileTypeCounter["sum"] = sum

	return s

}
func (s *SourceCounter) String() string {
	bytes, err := json.Marshal(s.FileTypeCounter)
	if err != nil {
		fmt.Println(err.Error())
	}
	if s.config.Debug {
		fmt.Println(string(bytes))
	}
	return string(bytes)
}
