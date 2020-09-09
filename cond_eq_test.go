package condition_test

import (
	"github.com/khlipeng/condition"
	"github.com/davecgh/go-spew/spew"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestExpr_EQ(t *testing.T) {
	var eqExpr = &condition.Expr{
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
	}

	var fields = map[string][]string{
		"check": []string{"back"},
	}
	v, s, err := eqExpr.LoadExOperator(condition.DefaultExOperators).Validator(fields)
	spew.Dump(v, s)
	require.NoError(t, err)
}
