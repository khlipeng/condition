package condition

type ExprGET struct{}

func init() {
	DefaultExOperators.Register(&ExprGET{})
}

func (expr *ExprGET) Op() string {
	return "GET"
}

func (expr *ExprGET) GetValues(key string, fields map[string][]string) []string {
	return fields[key]
}
