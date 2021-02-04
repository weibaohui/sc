package config

type Config struct {
	IgnoreHide bool
	Debug      bool
}

func New(ignoreHide bool, debug bool) *Config {
	return &Config{
		IgnoreHide: ignoreHide,
		Debug:      debug,
	}
}
