package config

import (
	"encoding/json"
	"fmt"
)

type Parser struct{}

func NewParser() Parser {
	return Parser{}
}
func (p Parser) Parse(format Format, inputData string) (string, error) {
	a := parsersMap[format]
	data, err := a.Parser(inputData)
	if err != nil {
		return "", err
	}

	j, err := json.Marshal(data)
	if err != nil {
		return "", fmt.Errorf("parse json: %w", err)
	}

	return string(j), err
}
