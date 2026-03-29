package parser_test

import (
	"testing"

	"endevix-cli-go/internal/config/parser"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestJsonParser_ToJson(t *testing.T) {
	t.Run("Successful parsing of simple JSON", func(t *testing.T) {
		p := parser.NewJsonParser()
		raw := `{"name": "test", "version": "1.0.0"}`
		want := map[string]any{
			"name":    "test",
			"version": "1.0.0",
		}

		err := p.ToJson(raw)
		assert.NoError(t, err)
		assert.Equal(t, want, p.FormatedData)
	})

	t.Run("Successful parsing of nested JSON", func(t *testing.T) {
		p := parser.NewJsonParser()
		raw := `{
				"user": {
					"id": 1,
					"active": true
				},
				"tags": ["go", "json"]
			}`
		want := map[string]any{
			"user": map[string]any{
				"id":     1.0,
				"active": true,
			},
			"tags": []any{"go", "json"},
		}

		err := p.ToJson(raw)
		assert.NoError(t, err)
		assert.Equal(t, want, p.FormatedData)
	})

	t.Run("Invalid JSON - Malformed", func(t *testing.T) {
		p := parser.NewJsonParser()
		raw := `{"name": "test", "version": "1.0.0"` // missing closing brace

		err := p.ToJson(raw)
		assert.Error(t, err)
	})

	t.Run("Invalid JSON - Empty string", func(t *testing.T) {
		p := parser.NewJsonParser()
		raw := ""

		err := p.ToJson(raw)
		assert.Error(t, err)
	})

	t.Run("Invalid JSON - Not an object", func(t *testing.T) {
		p := parser.NewJsonParser()
		raw := `["item1", "item2"]`

		err := p.ToJson(raw)
		assert.Error(t, err)
	})
}

func TestJsonParser_Parser(t *testing.T) {
	t.Run("Parser should call ToJson and return FormatedData", func(t *testing.T) {
		p := parser.NewJsonParser()
		raw := `{"key": "value"}`
		want := map[string]any{"key": "value"}

		data, err := p.Parser(raw)

		require.NoError(t, err)
		assert.Equal(t, want, data)
		assert.Equal(t, want, p.FormatedData)
	})

	t.Run("Parser should return error on invalid JSON", func(t *testing.T) {
		p := parser.NewJsonParser()
		_, err := p.Parser(`invalid`)

		assert.Error(t, err)
	})
}
