package expression

import "github.com/stuartmscott/pncgo"

type Multiply struct {
	A, B pncgo.Expression
}

func (e *Multiply) Resolve() (float64, error) {
	a, err := e.A.Resolve()
	if err != nil {
		return 0, err
	}
	b, err := e.B.Resolve()
	if err != nil {
		return 0, err
	}
	return a * b, nil
}
