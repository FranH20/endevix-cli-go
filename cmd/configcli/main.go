package main

import (
	"endevix-cli-go/internal/config"
	"fmt"
	"os"
)

func readFile(path string) string {
	data, err := os.ReadFile(path)
	if err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		panic(err)
	}
	fmt.Println("File read successfully")
	return string(data)
}

func main() {
	stringFile := readFile("/Users/franklinhuichi/GolandProjects/endevix-cli-go/testdata/valid.json")
	config.ParserJson(stringFile)
}
