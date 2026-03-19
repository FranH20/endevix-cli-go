package config

import "endevix-cli-go/internal/config/parser"

type Format string

const (
	FormatJSON Format = "json"
	FormatYAML Format = "yaml"
)

type ParserFunc func(string) (map[string]interface{}, error)

var parsersMap = map[Format]ParserFunc{
	FormatJSON: parser.JsonToMap,
}

type Model struct {
	ServiceName  string   `yaml:"service_name" json:"service_name"`
	Port         int      `yaml:"port" json:"port"`
	Environment  string   `yaml:"environment" json:"environment"`
	Enabled      bool     `yaml:"enabled" json:"enabled"`
	Dependencies []string `yaml:"dependencies" json:"dependencies"`
}
