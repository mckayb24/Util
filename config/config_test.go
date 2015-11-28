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
		{
			&struct {
				Uint   uint   `default:"7"`
				Uint8  uint8  `default:"8"`
				Uint16 uint16 `default:"16"`
				Uint32 uint32 `default:"32"`
				Uint64 uint64 `default:"64"`
			}{},
			&struct {
				Uint   uint   `default:"7"`
				Uint8  uint8  `default:"8"`
				Uint16 uint16 `default:"16"`
				Uint32 uint32 `default:"32"`
				Uint64 uint64 `default:"64"`
			}{
				7,
				8,
				16,
				32,
				64,
			},
			"struct with uint types",
		},
		{
			&struct {
				Float32 float32 `default:"32.32"`
				Float64 float64 `default:"64.64"`
			}{},
			&struct {
				Float32 float32 `default:"32.32"`
				Float64 float64 `default:"64.64"`
			}{
				32.32,
				64.64,
			},
			"struct with float types",
		},
	}

	for _, test := range tests {
		assert.NoError(t, Create(test.input), "Error", test.reason)
		assert.Equal(t, test.expected, test.input, test.reason)
	}
}
