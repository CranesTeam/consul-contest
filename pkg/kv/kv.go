package kv

type TaskConfig struct {
	Tasks []struct {
		Name  string `yaml:"name"`
		Limit int    `yaml:"limit"`
	} `yaml:"tasks"`
}
