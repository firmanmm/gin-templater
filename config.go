package gintemplater

type Config struct {
	AutoReload bool
	InputDir   string
	OutputDir  string
}

func NewConfig() *Config {
	instance := new(Config)
	instance.InputDir = "view"
	instance.OutputDir = "cache/view"
	instance.AutoReload = true
	return instance
}
