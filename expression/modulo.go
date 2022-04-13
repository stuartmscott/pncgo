package expression

import (
	"fmt"
	"github.com/stuartmscott/pncgo"
	"math"
)

type Modulo struct {
	A, B pncgo.Expression
}

func (e *Modulo) Resolve() (float64, error) {
	a, err := e.A.Resolve()
	if err != nil {
		return 0, err
	}
	b, err := e.B.Resolve()
	if err != nil {
		return 0, err
	}
	if b == 0 {
		return 0, fmt.Errorf("Modulo by Zero")
	}
	return math.Mod(a, b), nil
}
