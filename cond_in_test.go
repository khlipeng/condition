package conditionexpr_test

import (
	"github.com/davecgh/go-spew/spew"
	"github.com/khlipeng/conditionexpr"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestExpr_IN(t *testing.T) {
	var eqExpr = &conditionexpr.Expr{
		Op: "IN",
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
		"name": []string{"snail"},
	}
	v, s, err := eqExpr.LoadExOperator(conditionexpr.DefaultExOperators).Validator(fields)
	spew.Dump(v, s)
	require.NoError(t, err)
}
