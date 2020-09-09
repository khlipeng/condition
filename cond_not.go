package condition

type CondNOT struct{}

func init() {
	DefaultExOperators.Register(&CondNOT{})
}

func (expr *CondNOT) Validator(values ...string) bool {
	if len(values) <= 1 {
		return true
	}

	for i := 1; i < len(values); i++ {
		if values[0] == values[i] {
			return false
		}
	}

	return true
}

func (expr *CondNOT) Op() string {
	return "NOT"
}
