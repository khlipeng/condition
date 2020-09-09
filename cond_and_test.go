package condition_test

import (
	"github.com/khlipeng/condition"
	"github.com/davecgh/go-spew/spew"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestExprAND(t *testing.T) {
	var eqExpr = &condition.Expr{
		Op: "AND",
		Values: []interface{}{
			&condition.Expr{
				Op: "EQ",
				Values: []interface{}{
					&condition.Expr{
						Op: "GET",
						Values: []interface{}{
							"check_node",
						},
					},
					"back",
				},
			},

			&condition.Expr{
				Op: "LT",
				Values: []interface{}{
					&condition.Expr{
						Op: "GET",
						Values: []interface{}{
							"age",
						},
					},
					"100",
				},
			},
		},
	}

	var fields = map[string][]string{
		"check_node": []string{"back"},
		"age":        []string{"70"},
	}
	v, s, err := eqExpr.LoadExOperator(condition.DefaultExOperators).Validator(fields)
	spew.Dump(v, s)
	require.NoError(t, err)
}
