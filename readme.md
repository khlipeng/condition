# 自定义条件表达式

* 表达式示例

```json
{
    "op": "EQ",
    "values": [
        {
            "op": "GET",
            "values": [
                "check"
            ]
        },
        "pass"
    ]
}
```

* 代码
```go
    import (
    	"github.com/khlipeng/conditionexpr"
    	"fmt"
    )
	var fields = map[string][]string{
		"check": []string{"pass"},
	}

	expr,err:= conditionexpr.ParseExpr(`
{"op":"EQ","values":[{"op":"GET","values":["check"]},"pass"]}
`)
	if err!=nil{
    	panic(err)
    }   

	f, _, err := expr.LoadExOperator(conditionexpr.DefaultExOperators).Validator(fields)
	if err!=nil{
    	panic(err)
    }   
    fmt.Println(f)
```