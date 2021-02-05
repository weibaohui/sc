package config

// 配置
type Config struct {
	IgnoreHide bool
	Debug      bool
	Exclude    []string // 排除文件夹,逗号分割
}

// 创建默认的配置
func New(ignoreHide bool, debug bool) *Config {
	return &Config{
		IgnoreHide: ignoreHide,
		Debug:      debug,
		Exclude:    []string{"node_modules", "vendor"},
	}
}
