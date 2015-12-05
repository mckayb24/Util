package config

import "testing"
import "github.com/stretchr/testify/assert"

func TestParse(t *testing.T) {
	tests := []struct {
		input    interface{}
		expected interface{}
		reason   string
	}{
		{
			&struct {
				Int   int   `default:"7"`
				Int64 int64 `default:"64"`
			}{},
			&struct {
				Int   int   `default:"7"`
				Int64 int64 `default:"64"`
			}{
				7,
				64,
			},
			"struct with int types",
		},
		{
			&struct {
				Uint   uint   `default:"7"`
				Uint64 uint64 `default:"64"`
			}{},
			&struct {
				Uint   uint   `default:"7"`
				Uint64 uint64 `default:"64"`
			}{
				7,
				64,
			},
			"struct with uint types",
		},
		{
			&struct {
				Float64 float64 `default:"64.64"`
			}{},
			&struct {
				Float64 float64 `default:"64.64"`
			}{
				64.64,
			},
			"struct with float types",
		},
		{
			&struct {
				String string `default:"string"`
				Bool   bool   `default:"true"`
			}{},
			&struct {
				String string `default:"string"`
				Bool   bool   `default:"true"`
			}{
				"string",
				true,
			},
			"struct with string and bool types",
		},
	}

	for _, test := range tests {
		assert.NoError(t, Parse(test.input), "Error", test.reason)
		assert.Equal(t, test.expected, test.input, test.reason)
	}
}
