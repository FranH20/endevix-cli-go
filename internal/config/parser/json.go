package parser

import (
	"encoding/json"
	"fmt"
)

func JsonToMap(jsonData string) (map[string]interface{}, error) {
	var structJsonData map[string]interface{}

	err := json.Unmarshal([]byte(jsonData), &structJsonData)
	if err != nil {
		fmt.Println("error:", err)
		return nil, err
	}

	fmt.Println("Parsed JSON data:", structJsonData)
	return structJsonData, nil
}
