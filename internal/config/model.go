package config

type Model struct {
	ServiceName  string   `yaml:"service_name" json:"service_name"`
	Port         int      `yaml:"port" json:"port"`
	Environment  string   `yaml:"environment" json:"environment"`
	Enabled      bool     `yaml:"enabled" json:"enabled"`
	Dependencies []string `yaml:"dependencies" json:"dependencies"`
}
