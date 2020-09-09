package conditionexpr_test

import (
	"github.com/davecgh/go-spew/spew"
	"github.com/khlipeng/conditionexpr"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestExpr_GT(t *testing.T) {
	var eqExpr = &conditionexpr.Expr{
		Op: "GT",
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
		"age": []string{"20"},
	}
	v, s, err := eqExpr.LoadExOperator(conditionexpr.DefaultExOperators).Validator(fields)
	spew.Dump(v, s)
	require.NoError(t, err)
}
