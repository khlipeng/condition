package conditionexpr

type CondAND struct{}

func init() {
	DefaultExOperators.Register(&CondAND{})
}

func (expr *CondAND) Condition(conds ...bool) bool {
	for _, v := range conds {
		if !v {
			return false
		}
	}
	return true
}

func (expr *CondAND) Op() string {
	return "AND"
}
