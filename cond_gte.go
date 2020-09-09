package condition

import (
	"strconv"
)

type CondGTE struct{}

func init() {
	DefaultExOperators.Register(&CondGTE{})
}

func (expr *CondGTE) Validator(values ...string) bool {
	if len(values) != 2 {
		return false
	}
	a, err := strconv.ParseFloat(values[0], 64)
	if err != nil {
		return false
	}
	b, err := strconv.ParseFloat(values[1], 64)
	if err != nil {
		return false
	}
	return a >= b
}

func (expr *CondGTE) Op() string {
	return "GTE"
}
