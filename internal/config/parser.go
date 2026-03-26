package config

import (
	"encoding/json"
	"fmt"
)

func (f Format) Parser(data string) (map[string]interface{}, error) {
	parserFunc := parsersMap[f]
	return parserFunc(data)
}

func Parse(format Format, inputData string) (string, error) {
	// el switch pero con map
	jsonInputData, err := format.Parser(inputData)

	if err != nil {
		return "", fmt.Errorf("parse input: %w", err)
	}

	j, err := json.Marshal(jsonInputData)
	if err != nil {
		return "", fmt.Errorf("parse json: %w", err)
	}

	return string(j), err
}
