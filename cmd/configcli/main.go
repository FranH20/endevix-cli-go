package main

import (
	"endevix-cli-go/internal/cli"
	"endevix-cli-go/internal/config"
	"fmt"
	"os"
)

func main() {

	if err := run(); err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}

}

func run() error {
	// 1. Obtener el input data
	inputData, err := cli.GetInput()
	if err != nil {
		return err
	}

	// 2. Obtener el input json format
	jsonSchemaData, err := cli.GetJsonFormatData()
	if err != nil {
		return err
	}

	// 3. Parsear el input a json
	jsonInput, err := config.Parse(config.FormatJSON, string(inputData))
	if err != nil {
		return err
	}

	// 4. Validar el json con el schema
	err = config.Validate(string(jsonSchemaData), jsonInput)
	if err != nil {
		return err
	}

	return err
}
