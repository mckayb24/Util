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
				Int   int   `default:"7"`
				Int8  int8  `default:"8"`
				Int16 int16 `default:"16"`
				Int32 int32 `default:"32"`
				Int64 int64 `default:"64"`
			}{},
			&struct {
				Int   int   `default:"7"`
				Int8  int8  `default:"8"`
				Int16 int16 `default:"16"`
				Int32 int32 `default:"32"`
				Int64 int64 `default:"64"`
			}{
				7,
				8,
				16,
				32,
				64,
			},
			"struct with int types",
		},
	}

	for _, test := range tests {
		assert.NoError(t, Create(test.input), "Error", test.reason)
		assert.Equal(t, test.expected, test.input, test.reason)
	}
}
