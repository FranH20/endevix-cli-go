package config

import (
	"encoding/json"
	"fmt"
)

func ParserJson(jsonData string) {

	var structJsonData map[string]interface{}

	err := json.Unmarshal([]byte(jsonData), &structJsonData)
	if err != nil {
		fmt.Println("error:", err)
		panic(err)
	}

	for key, value := range structJsonData {
		fmt.Println("key:", key, "value:", value)
	}

}
