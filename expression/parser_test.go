package expression_test

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stuartmscott/pncgo/expression"
	"testing"
)

func TestParseOperation(t *testing.T) {
	t.Run("Supported", func(t *testing.T) {
		for name, tt := range map[string]struct {
			input    []string
			expected float64
		}{
			"add_int": {
				input:    []string{"+", "2", "1"},
				expected: 3,
			},
			"subtract_int": {
				input:    []string{"-", "2", "1"},
				expected: 1,
			},
			"multiply_int": {
				input:    []string{"*", "2", "1"},
				expected: 2,
			},
			"divide_int": {
				input:    []string{"/", "2", "1"},
				expected: 2,
			},
			"modulo_int": {
				input:    []string{"%", "2", "1"},
				expected: 0,
			},
			"add_float": {
				input:    []string{"+", "2.1", "1.2"},
				expected: 3.3,
			},
			"subtract_float": {
				input:    []string{"-", "2.1", "1.2"},
				expected: 0.9000000000000001, // Floating point weirdness
			},
			"multiply_float": {
				input:    []string{"*", "2.1", "1.2"},
				expected: 2.52,
			},
			"divide_float": {
				input:    []string{"/", "2.1", "1.2"},
				expected: 1.7500000000000002, // Floating point weirdness
			},
			"modulo_float": {
				input:    []string{"%", "2.1", "1.2"},
				expected: 0.9000000000000001, // Floating point weirdness
			},
			"combination_subtractmulitply": {
				// Infix
				// (5 - 6) * 7
				// Polish
				// - 5 6 = -1
				// * -1 7 = -7
				input:    []string{"*", "-", "5", "6", "7"},
				expected: -7,
			},
			"combination_mulitplysubtract": {
				// Infix
				// 5 - (6 * 7)
				// Polish
				// * 6 7 = 42
				// - 5 42 = -37
				input:    []string{"-", "5", "*", "6", "7"},
				expected: -37,
			},
			"combination_all": {
				// Polish
				// + 2 1 = 3
				// - 3 1 = 2
				// * 2 4 = 8
				// / 8 1 = 8
				// % 8 3 = 2
				input:    []string{"%", "/", "*", "-", "+", "2", "1", "1", "4", "1", "3"},
				expected: 2,
			},
		} {
			t.Run(name, func(t *testing.T) {
				exp, err := expression.Parse(tt.input)
				assert.NoError(t, err)
				result, err := exp.Resolve()
				assert.NoError(t, err)
				assert.Equal(t, tt.expected, result)
			})
		}
	})
	t.Run("Unsupported", func(t *testing.T) {
		exp, err := expression.Parse([]string{"@"})
		assert.Nil(t, exp)
		assert.Error(t, fmt.Errorf("Unrecognized Operation: @"), err)
	})
}
