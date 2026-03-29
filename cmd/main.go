package main

import (
	"endevix-cli-go/internal/command_cli"
	"fmt"
	"os"

	"endevix-cli-go/internal/config"
)

func main() {
	if err := run(); err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}
}

func run() error {
	cli := command_cli.NewCommandCli()
	err := cli.GetInput()
	if err != nil {
		return err
	}
	err = cli.GetJsonFormatData()
	if err != nil {
		return err
	}
	parser := config.NewParser()
	jsonInput, err := parser.Parse(config.FormatJSON, string(cli.InputData))
	if err != nil {
		return err
	}
	err = config.Validate(string(cli.FormatData), jsonInput)
	if err != nil {
		return err
	}

	return err
}
