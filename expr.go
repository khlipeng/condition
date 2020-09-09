package conditionexpr

import "encoding/json"

type Expr struct {
	ef     *ExOperator   `json:"-"`
	Op     string        `json:"op"`
	Values []interface{} `json:"values"`
}

func (e *Expr) LoadExOperator(ef *ExOperator) *Expr {
	e.ef = ef
	return e
}

func (e *Expr) Validator(fields map[string][]string) (bool, []string, error) {
	evs := []string{}
	vList := []bool{}
	for _, ev := range e.Values {
		switch ev.(type) {
		case string:
			evs = append(evs, ev.(string))
			break
		case *Expr, Expr:
			isValidator, vl, err := ev.(*Expr).LoadExOperator(e.ef).Validator(fields)
			if err != nil {
				return false, []string{}, err
			}
			vList = append(vList, isValidator)
			evs = append(evs, vl...)
			break
		case map[string]interface{}:
			mev := ev.(map[string]interface{})
			newExpr := &Expr{}
			if _, ok := mev["op"].(string); ok {
				newExpr.Op = mev["op"].(string)
			}
			if _, ok := mev["values"].([]interface{}); ok {
				newExpr.Values = mev["values"].([]interface{})
			}

			isValidator, vl, err := newExpr.LoadExOperator(e.ef).Validator(fields)
			if err != nil {
				return false, []string{}, err
			}
			vList = append(vList, isValidator)
			evs = append(evs, vl...)
			break
		}
	}

	if _, ok := e.ef.exprSet[e.Op].(*ExprGET); ok {
		values := e.ef.exprSet[e.Op].(*ExprGET).GetValues(evs[0], fields)
		return false, values, nil
	}

	if _, ok := e.ef.exprSet[e.Op].(ExprValidator); ok {
		return e.ef.exprSet[e.Op].(ExprValidator).Validator(evs...), []string{}, nil
	}

	if _, ok := e.ef.exprSet[e.Op].(ExprCondition); ok {
		return e.ef.exprSet[e.Op].(ExprCondition).Condition(vList...), []string{}, nil
	}

	return false, []string{}, nil
}

func  ParseExpr(exprStr string) (*Expr, error) {
	expr:=&Expr{}
	err := json.Unmarshal([]byte(exprStr), expr)
	if err!=nil{
		return nil,err
	}
	return expr,nil
}
