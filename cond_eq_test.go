package conditionexpr_test

import (
	"github.com/davecgh/go-spew/spew"
	"github.com/khlipeng/conditionexpr"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestExpr_EQ(t *testing.T) {
	var eqExpr = &conditionexpr.Expr{
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
	}

	var fields = map[string][]string{
		"check": []string{"back"},
	}
	v, s, err := eqExpr.LoadExOperator(conditionexpr.DefaultExOperators).Validator(fields)
	spew.Dump(v, s)
	require.NoError(t, err)
}
