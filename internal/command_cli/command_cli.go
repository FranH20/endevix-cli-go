package command_cli

import (
	"fmt"
	"io"
	"os"
)

type CommandCli struct {
	InputData  []byte
	FormatData []byte
}

func NewCommandCli() CommandCli {
	return CommandCli{}
}

func (c *CommandCli) GetInput() error {
	if len(os.Args) > 1 {
		data, err := os.ReadFile(os.Args[1])
		if err != nil {
			return fmt.Errorf("read input file %q: %w", os.Args[1], err)
		}

		c.InputData = data
		return nil
	}

	data, err := io.ReadAll(os.Stdin)
	if err != nil {
		return fmt.Errorf("read stdin: %w", err)
	}

	c.InputData = data
	return nil
}

func (c *CommandCli) GetJsonFormatData() error {
	if len(os.Args) > 2 {
		data, err := os.ReadFile(os.Args[2])
		if err != nil {
			return fmt.Errorf("read input file %q: %w", os.Args[1], err)
		}

		c.FormatData = data
		return nil
	}

	return fmt.Errorf("no json format data provided")
}
