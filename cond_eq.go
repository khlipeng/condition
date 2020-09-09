package condition

type CondEQ struct{}

func init() {
	DefaultExOperators.Register(&CondEQ{})
}

func (expr *CondEQ) Validator(values ...string) bool {
	if len(values) != 2 {
		return false
	}
	return values[0] == values[1]
}

func (expr *CondEQ) Op() string {
	return "EQ"
}
