package conditionexpr

type CondOR struct{}

func init() {
	DefaultExOperators.Register(&CondOR{})
}

func (expr *CondOR) Condition(conds ...bool) bool {
	for _, v := range conds {
		if v {
			return true
		}
	}
	return false
}

func (expr *CondOR) Op() string {
	return "OR"
}
