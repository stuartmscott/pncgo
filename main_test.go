package main_test

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stuartmscott/pncgo"
	"testing"
)

func TestParseOperation(t *testing.T) {
	t.Run("Supported", func(t *testing.T) {
		for name, tt := range map[string]struct {
			operation string
			expected  int
		}{
			"add": {
				operation: "+",
				expected:  3,
			},
			"subtract": {
				operation: "-",
				expected:  1,
			},
			"multiply": {
				operation: "*",
				expected:  2,
			},
			"divide": {
				operation: "/",
				expected:  2,
			},
			"modulo": {
				operation: "%",
				expected:  0,
			},
		} {
			t.Run(name, func(t *testing.T) {
				op, err := main.ParseOperation(tt.operation)
				assert.NoError(t, err)
				assert.Equal(t, tt.expected, op(2, 1))
			})
		}
	})
	t.Run("Unsupported", func(t *testing.T) {
		op, err := main.ParseOperation("@")
		assert.Nil(t, op)
		assert.Error(t, fmt.Errorf("Unrecognized Operation: @"), err)
	})
}
