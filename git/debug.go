package git

import (
	"fmt"

	"github.com/weibaohui/sc/config"
)

func Debug(a ...interface{}) {
	if config.GetInstance().Debug {
		fmt.Println(a...)
	}
}
func Debugf(f string, a ...interface{}) {
	if config.GetInstance().Debug {
		fmt.Printf(f, a...)
	}
}
