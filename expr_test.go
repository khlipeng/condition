package condition_test

import (
	"encoding/json"
	"github.com/khlipeng/condition"
	"github.com/davecgh/go-spew/spew"
	"github.com/stretchr/testify/require"
	"testing"
)

var Fields = map[string][]string{
	"user":      []string{"user_02"},
	"check":     []string{"back"},
	"last_node": []string{"node_1"},
	"back_node": []string{"node_2"},
}

var ExprCond = &condition.Expr{
	Op: "AND",
	Values: []interface{}{
		&condition.Expr{
			Op: "AND",
			Values: []interface{}{
				&condition.Expr{
					Op: "EQ",
					Values: []interface{}{
						&condition.Expr{
							Op: "GET",
							Values: []interface{}{
								"check",
							},
						},
						"back",
					},
				},
				&condition.Expr{
					Op: "EQ",
					Values: []interface{}{
						&condition.Expr{
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
		&condition.Expr{
			Op: "EQ",
			Values: []interface{}{
				&condition.Expr{
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
	spew.Dump(ExprCond.LoadExOperator(condition.DefaultExOperators))
}

func TestExprText(t *testing.T) {
	str := `
{"op":"EQ","values":[{"op":"GET","values":["check"]},"pass"]}
`
	var fields = map[string][]string{
		"check": []string{"pass"},
	}

	ev := &condition.Expr{}
	err := json.Unmarshal([]byte(str), ev)
	require.NoError(t, err)

	f, _, _ := ev.LoadExOperator(condition.DefaultExOperators).Validator(fields)
	spew.Dump(f)
}
