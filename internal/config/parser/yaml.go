package parser

type YamlParser struct{}

func NewYamlParser() *YamlParser {
	return &YamlParser{}
}

func (y *YamlParser) Parser(raw string) (map[string]interface{}, error) {
	return nil, nil
}
