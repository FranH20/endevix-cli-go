package config

import (
	"fmt"
	"strings"

	"github.com/santhosh-tekuri/jsonschema/v6"
)

func Validate(schema string, inputJson string) error {
	schemaOfProject, err := jsonschema.UnmarshalJSON(strings.NewReader(schema))
	if err != nil {
		return fmt.Errorf("unmarshal schema: %w", err)
	}
	inputOfProject, err := jsonschema.UnmarshalJSON(strings.NewReader(inputJson))
	if err != nil {
		return fmt.Errorf("unmarshal input: %w", err)
	}

	c := jsonschema.NewCompiler()
	if err := c.AddResource("schema.json", schemaOfProject); err != nil {
		return fmt.Errorf("compile schema: %w", err)
	}
	sch, err := c.Compile("schema.json")
	if err != nil {
		return fmt.Errorf("compile schema: %w", err)
	}
	err = sch.Validate(inputOfProject)
	fmt.Println("valid:", err == nil)
	return err
}
