package config

import (
	"runtime"
	"sync"
)

var defaultConcurrency = runtime.GOMAXPROCS(0)

// Config 配置
type Config struct {
	IgnoreHide  bool
	Debug       bool
	Exclude     []string // 排除文件夹,逗号分割
	Concurrency int
}

var c *Config
var once sync.Once

func init() {
	once.Do(func() {
		c = &Config{
			IgnoreHide:  true,
			Debug:       false,
			Concurrency: defaultConcurrency,
			Exclude:     []string{"node_modules", "vendor", "pod"},
		}
	})
}

// SetConfig 配置Config
func (c *Config) SetConfig(ignoreHide bool, debug bool) *Config {
	c.IgnoreHide = ignoreHide
	c.Debug = debug
	return c
}

// GetInstance get an Instance
func GetInstance() *Config {
	return c
}
