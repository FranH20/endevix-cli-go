package config

import (
	"endevix-cli-go/internal/config/parser"
)

type Format string

//go:generate mockery --name ParserFn
type ParserFn interface {
	Parser(string) (map[string]interface{}, error)
}

const (
	FormatJSON Format = "json"
	FormatYAML Format = "yaml"
)

var parsersMap = map[Format]ParserFn{
	FormatJSON: parser.NewJsonParser(),
	FormatYAML: parser.NewYamlParser(),
}
