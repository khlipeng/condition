package conditionexpr_test

import (
	"github.com/davecgh/go-spew/spew"
	"github.com/khlipeng/conditionexpr"
	"github.com/stretchr/testify/require"
	"testing"
)

var Fields = map[string][]string{
	"user":      []string{"user_02"},
	"check":     []string{"back"},
	"last_node": []string{"node_1"},
	"back_node": []string{"node_2"},
}

var ExprCond = &conditionexpr.Expr{
	Op: "AND",
	Values: []interface{}{
		&conditionexpr.Expr{
			Op: "AND",
			Values: []interface{}{
				&conditionexpr.Expr{
					Op: "EQ",
					Values: []interface{}{
						&conditionexpr.Expr{
							Op: "GET",
							Values: []interface{}{
								"check",
							},
						},
						"back",
					},
				},
				&conditionexpr.Expr{
					Op: "EQ",
					Values: []interface{}{
						&conditionexpr.Expr{
							Op: "GET",
							Values: []interface{}{
								"user",
							},
						},
						"user_02",
					},
				},
			},
		},
		&conditionexpr.Expr{
			Op: "EQ",
			Values: []interface{}{
				&conditionexpr.Expr{
					Op: "GET",
					Values: []interface{}{
						"last_node",
					},
				},
				"node_01",
			},
		},
	},
}

func TestExpr(t *testing.T) {
	spew.Dump(ExprCond.LoadExOperator(conditionexpr.DefaultExOperators))
}

func TestExprText(t *testing.T) {
	var fields = map[string][]string{
		"check": []string{"pass"},
	}

	expr,err:= conditionexpr.ParseExpr(`
{"op":"EQ","values":[{"op":"GET","values":["check"]},"pass"]}
`)
	require.NoError(t, err)
	f, v, err := expr.LoadExOperator(conditionexpr.DefaultExOperators).Validator(fields)
	require.NoError(t, err)
	spew.Dump(f)
	spew.Dump(v)

}
