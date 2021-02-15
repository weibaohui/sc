package config

import (
	"sync"
)

// config 配置
type config struct {
	IgnoreHide bool
	Debug      bool
	Exclude    []string // 排除文件夹,逗号分割
}

var c *config
var once sync.Once

func init() {
	once.Do(func() {
		c = &config{
			IgnoreHide: true,
			Debug:      false,
			Exclude:    []string{"node_modules", "vendor", "pod"},
		}
	})
}

// New 创建默认的配置
func (c *config) SetConfig(ignoreHide bool, debug bool) *config {
	c.IgnoreHide = ignoreHide
	c.Debug = debug
	return c
}

func Instant() *config {
	return c
}
