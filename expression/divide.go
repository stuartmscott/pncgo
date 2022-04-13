package expression

import (
	"fmt"
	"github.com/stuartmscott/pncgo"
)

type Divide struct {
	A, B pncgo.Expression
}

func (e *Divide) Resolve() (float64, error) {
	a, err := e.A.Resolve()
	if err != nil {
		return 0, err
	}
	b, err := e.B.Resolve()
	if err != nil {
		return 0, err
	}
	if b == 0 {
		return 0, fmt.Errorf("Division by Zero")
	}
	return a / b, nil
}
