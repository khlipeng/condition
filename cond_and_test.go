package conditionexpr_test

import (
	"github.com/davecgh/go-spew/spew"
	"github.com/khlipeng/conditionexpr"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestExprAND(t *testing.T) {
	var eqExpr = &conditionexpr.Expr{
		Op: "AND",
		Values: []interface{}{
			&conditionexpr.Expr{
				Op: "EQ",
				Values: []interface{}{
					&conditionexpr.Expr{
						Op: "GET",
						Values: []interface{}{
							"check_node",
						},
					},
					"back",
				},
			},

			&conditionexpr.Expr{
				Op: "LT",
				Values: []interface{}{
					&conditionexpr.Expr{
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
	v, s, err := eqExpr.LoadExOperator(conditionexpr.DefaultExOperators).Validator(fields)
	spew.Dump(v, s)
	require.NoError(t, err)
}
