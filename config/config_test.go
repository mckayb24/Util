package config

import "testing"
import "github.com/stretchr/testify/assert"

func TestCreate(t *testing.T) {
	tests := []struct {
		input    interface{}
		expected interface{}
		reason   string
	}{
		{
			&struct {
				String string `default:"string"`
				Int    int    `default:"7"`
				Int8   int8   `default:"8"`
			}{},
			&struct {
				String string `default:"string"`
				Int    int    `default:"7"`
				Int8   int8   `default:"8"`
			}{
				"string",
				7,
				8,
			},
			"struct with string, int, and int8",
		},
	}

	for _, test := range tests {
		assert.NoError(t, Create(test.input), "Error", test.reason)
		assert.Equal(t, test.expected, test.input, test.reason)
	}
}
