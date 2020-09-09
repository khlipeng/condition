package condition

import (
	"strconv"
)

type CondLTE struct{}

func init() {
	DefaultExOperators.Register(&CondLTE{})
}

func (expr *CondLTE) Validator(values ...string) bool {
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
	return a <= b
}

func (expr *CondLTE) Op() string {
	return "LTE"
}
