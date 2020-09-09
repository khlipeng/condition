package conditionexpr_test

import (
	"github.com/davecgh/go-spew/spew"
	"github.com/khlipeng/conditionexpr"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestExpr_NOT(t *testing.T) {
	var eqExpr = &conditionexpr.Expr{
		Op: "NOT",
		Values: []interface{}{
			&conditionexpr.Expr{
				Op: "GET",
				Values: []interface{}{
					"name",
				},
			},
			"zs",
			"ls",
			"leo",
			"snail",
		},
	}

	var fields = map[string][]string{
		"name": []string{"lp"},
	}
	v, s, err := eqExpr.LoadExOperator(conditionexpr.DefaultExOperators).Validator(fields)
	spew.Dump(v, s)
	require.NoError(t, err)
}
