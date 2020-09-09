package condition_test

import (
	"github.com/khlipeng/condition"
	"github.com/davecgh/go-spew/spew"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestExpr_GT(t *testing.T) {
	var eqExpr = &condition.Expr{
		Op: "GT",
		Values: []interface{}{
			&condition.Expr{
				Op: "GET",
				Values: []interface{}{
					"age",
				},
			},
			"10",
		},
	}

	var fields = map[string][]string{
		"age": []string{"20"},
	}
	v, s, err := eqExpr.LoadExOperator(condition.DefaultExOperators).Validator(fields)
	spew.Dump(v, s)
	require.NoError(t, err)
}
