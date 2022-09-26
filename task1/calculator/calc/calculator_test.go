package calculator

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCalculateSuccess(t *testing.T) {
	tests := map[string]struct {
		input string
		result float64
	} {
		"addition":  {
			input: "3+4",
			result: 7,
		},
		"subtraction": {
			input: "10-5",
			result: 5,
		},
		"division": {
			input: "9/2",
			result: 4.5,
		},
		"multiplication": {
			input: "3*8",
			result: 24,
		},
		"operationsWithDifferentPriority":  {
			input: "2+2*2",
			result: 6,
		},
		"withBrackets": {
			input: "(2+2)*2",
			result: 8,
		},
	}
	  
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			gotRes, err := Calculate(test.input)
			require.NoError(t, err)
			require.Equal(t, gotRes, test.result)
		})
	}
}

func TestCalculateFailed(t *testing.T) {
	tests := map[string]struct {
		input string
		result float64
	} {
		"incorrect expression: unknown symbol":  {
			input: "1 5",
			result: 0,
		},
		"incorrect expression: not enough operands":  {
			input: "1-",
			result: 0,
		},
		"division by zero": {
			input: "43/0",
			result: 0,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			gotRes, err := Calculate(test.input)
			require.Error(t, err)
			require.Equal(t, gotRes, test.result)
		})
	}
}

