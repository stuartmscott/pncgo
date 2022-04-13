package expression

type Literal struct {
	Value float64
}

func (e *Literal) Resolve() (float64, error) {
	return e.Value, nil
}
