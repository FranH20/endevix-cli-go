package parser

import (
	"encoding/json"
)

type JsonParser struct {
	rawJson      string
	FormatedData map[string]interface{}
}

func NewJsonParser() *JsonParser {
	return &JsonParser{}
}

func (j *JsonParser) Parser(raw string) (map[string]interface{}, error) {
	err := j.ToJson(raw)
	if err != nil {
		return nil, err
	}

	return j.FormatedData, nil
}

func (j *JsonParser) ToJson(raw string) error {
	var rawData map[string]interface{}
	err := json.Unmarshal([]byte(raw), &rawData)
	if err != nil {
		return err
	}

	j.FormatedData = rawData
	j.rawJson = raw
	return nil
}
