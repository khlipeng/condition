package conditionexpr_test

import (
	"github.com/khlipeng/condition"
	"github.com/davecgh/go-spew/spew"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestExpr_LTE(t *testing.T) {
	var eqExpr = &conditionexpr.Expr{
		Op: "LTE",
		Values: []interface{}{
			&conditionexpr.Expr{
				Op: "GET",
				Values: []interface{}{
					"age",
				},
			},
			"10",
		},
	}

	var fields = map[string][]string{
		"age": []string{"5"},
	}
	v, s, err := eqExpr.LoadExOperator(conditionexpr.DefaultExOperators).Validator(fields)
	spew.Dump(v, s)
	require.NoError(t, err)
}
