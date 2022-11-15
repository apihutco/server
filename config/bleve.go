package config

type Bleve struct {
	Index     string `mapstructure:"index"`
	SetupPath string `mapstructure:"setup_path"`
}
