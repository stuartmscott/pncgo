package pncgo

type Expression interface {
	Resolve() (float64, error)
}
