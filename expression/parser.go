package expression

import (
	"fmt"
	"github.com/stuartmscott/pncgo"
	"strconv"
	"unicode"
)

func Parse(parts []string) (pncgo.Expression, error) {
	var stack []pncgo.Expression
	// Start from right hand side
	for i := len(parts) - 1; i >= 0; i-- {
		part := parts[i]
		if unicode.IsDigit(([]rune(part))[0]) {
			// If part is a number, add Literal to stack
			f, err := strconv.ParseFloat(part, 64)
			if err != nil {
				return nil, err
			}
			stack = append(stack, &Literal{f})
		} else {
			// If part is an operator, add Operation to stack
			size := len(stack)
			switch part {
			case "+":
				if size < 2 {
					return nil, fmt.Errorf("Add takes 2 operands")
				}
				a, b := stack[size-1], stack[size-2]
				stack = append(stack[:size-2], &Add{a, b})
			case "-":
				if size < 2 {
					return nil, fmt.Errorf("Subtract takes 2 operands")
				}
				a, b := stack[size-1], stack[size-2]
				stack = append(stack[:size-2], &Subtract{a, b})
			case "*":
				if size < 2 {
					return nil, fmt.Errorf("Multiply takes 2 operands")
				}
				a, b := stack[size-1], stack[size-2]
				stack = append(stack[:size-2], &Multiply{a, b})
			case "/":
				if size < 2 {
					return nil, fmt.Errorf("Divide takes 2 operands")
				}
				a, b := stack[size-1], stack[size-2]
				stack = append(stack[:size-2], &Divide{a, b})
			case "%":
				if size < 2 {
					return nil, fmt.Errorf("Modulo takes 2 operands")
				}
				a, b := stack[size-1], stack[size-2]
				stack = append(stack[:size-2], &Modulo{a, b})
			default:
				return nil, fmt.Errorf("Unrecognized operation: %s", part)
			}
		}
	}
	if len(stack) != 1 {
		return nil, fmt.Errorf("Unused Parameter")
	}
	return stack[0], nil
}
