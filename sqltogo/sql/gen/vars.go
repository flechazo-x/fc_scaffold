package gen

import (
	"fc_scaffold/sqltogo/sql/stringx"
	"fmt"
	"strings"
)

const (
	defaultField = "Field"
)

func genVars(table Table) (string, error) {
	keys := make([]string, 0)
	keys = append(keys, table.PrimaryCacheKey.VarExpression)
	for _, v := range table.UniqueCacheKey {
		keys = append(keys, v.VarExpression)
	}

	var output strings.Builder
	output.WriteString(fmt.Sprintf("var _ db.INative = (*%s)(nil)\n", table.Name.ToCamel()))
	output.WriteString("const(")
	//字段名
	for _, field := range table.Fields {
		name := stringx.From(field.NameOriginal).ToCamel()
		output.WriteString(defaultField + name + "=" + `"` + name + `"` + "\n") //表名
	}
	output.WriteString(")")
	return output.String(), nil
}
