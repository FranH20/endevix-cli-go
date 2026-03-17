package main

import (
	"fmt"
	"os"
)

func readFile(path string) {
	data, err := os.ReadFile(path)
	if err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		panic(err)
	}
	fmt.Println(string(data))
}

func main() {
	readFile("/Users/franklinhuichi/GolandProjects/endevix-cli-go/testdata/valid.json")
}
