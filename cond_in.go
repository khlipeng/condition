package condition

type CondIN struct{}

func init() {
	DefaultExOperators.Register(&CondIN{})
}

func (expr *CondIN) Validator(values ...string) bool {
	if len(values) <= 1 {
		return false
	}

	for i := 1; i < len(values); i++ {
		if values[i] == values[0] {
			return true
		}
	}

	return false
}

func (expr *CondIN) Op() string {
	return "IN"
}
