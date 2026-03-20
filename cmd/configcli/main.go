package main

import (
	"endevix-cli-go/internal/config"
	"fmt"
	"io"
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
	var data []byte
	var err error

	if len(os.Args) > 1 {
		data, err = os.ReadFile(os.Args[1])
	} else {
		data, err = io.ReadAll(os.Stdin)
	}

	if err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		panic(err)
	}
	// /Users/franklinhuichi/GolandProjects/endevix-cli-go/testdata/valid.json
	//;stringFile := readFile(string(data))

	config.Parse(config.FormatJSON, string(data))
}
