package config

import (
	"fmt"
)

func (f Format) Parser(data string) (map[string]interface{}, error) {
	parserFunc := parsersMap[f]
	return parserFunc(data)
}

func Parse(format Format, data string) {
	// el switch pero con map
	doc, err := format.Parser(data)

	if err != nil {
		fmt.Println("Error parsing data:", err)
	}

	fmt.Println("Parsed data:", doc)
}
