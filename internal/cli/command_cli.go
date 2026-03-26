package cli

import (
	"fmt"
	"io"
	"os"
)

func GetInput() ([]byte, error) {
	if len(os.Args) > 1 {

		data, err := os.ReadFile(os.Args[1])
		if err != nil {
			return nil, fmt.Errorf("read input file %q: %w", os.Args[1], err)
		}

		return data, nil
	}

	data, err := io.ReadAll(os.Stdin)
	if err != nil {
		return nil, fmt.Errorf("read stdin: %w", err)
	}

	return data, nil
}

func GetJsonFormatData() ([]byte, error) {
	if len(os.Args) > 2 {

		data, err := os.ReadFile(os.Args[2])
		if err != nil {
			return nil, fmt.Errorf("read input file %q: %w", os.Args[1], err)
		}

		return data, nil
	}
	return nil, fmt.Errorf("no json format data provided")
}
