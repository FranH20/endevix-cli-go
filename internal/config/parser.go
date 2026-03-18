package config

import (
	"encoding/json"
	"fmt"
)

func ParserJson(jsonData string) {

	var structJsonData Model

	err := json.Unmarshal([]byte(jsonData), &structJsonData)
	if err != nil {
		fmt.Println("error:", err)
		panic(err)
	}

	fmt.Println("json parsed successfully: ", structJsonData)
}
