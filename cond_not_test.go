package condition_test

import (
	"github.com/khlipeng/condition"
	"github.com/davecgh/go-spew/spew"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestExpr_NOT(t *testing.T) {
	var eqExpr = &condition.Expr{
		Op: "NOT",
		Values: []interface{}{
			&condition.Expr{
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
	v, s, err := eqExpr.LoadExOperator(condition.DefaultExOperators).Validator(fields)
	spew.Dump(v, s)
	require.NoError(t, err)
}
