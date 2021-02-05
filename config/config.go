package config

type Config struct {
	IgnoreHide bool
	Debug      bool
	Exclude    []string // 排除文件夹,逗号分割
}

func New(ignoreHide bool, debug bool) *Config {
	return &Config{
		IgnoreHide: ignoreHide,
		Debug:      debug,
		Exclude:    []string{"node_modules", "vendor"},
	}
}
