package conditionexpr

var DefaultExOperators = NewExOperator()

// ex
type Ex interface {
	Op() string
}

// 条件验证
type ExprValidator interface {
	Validator(values ...string) bool
}

// 条件表达
type ExprCondition interface {
	Condition(conds ...bool) bool
}

type ExOperator struct {
	exprSet map[string]Ex
}

func NewExOperator() *ExOperator {
	return &ExOperator{
		exprSet: map[string]Ex{},
	}
}

func (e *ExOperator) Register(exs ...Ex) {
	for _, ex := range exs {
		e.exprSet[ex.Op()] = ex
	}
}
