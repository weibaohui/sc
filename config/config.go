package config

import (
	"runtime"
	"sync"
)

var defaultConcurrency = runtime.GOMAXPROCS(0)

// Config 配置
type Config struct {
	InitPath    string // 初始化路径
	IgnoreHide  bool
	Debug       bool
	Exclude     []string // 排除文件夹,逗号分割
	SkipSuffix  []string // 跳过文件后缀
	Concurrency int
	Force       bool // 使用自定义配置覆盖默认配置
}

var c *Config
var once sync.Once

func init() {
	once.Do(func() {
		c = &Config{
			InitPath:    ".",
			IgnoreHide:  true,
			Debug:       false,
			SkipSuffix:  []string{".ico", ".tf", ".xsl", ".log", ".dapper", ".json", ".3", ".2", ".1"},
			Exclude:     []string{"node_modules", "vendor", "pod", "dist", "target", "bin", "asset", "img", "assets"},
			Concurrency: defaultConcurrency,
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
