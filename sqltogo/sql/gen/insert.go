package gen

import (
	"fc_scaffold/sqltogo/sql/stringx"
	"fc_scaffold/sqltogo/sql/template"
	"fc_scaffold/sqltogo/sql/util"
	"fc_scaffold/sqltogo/sql/util/pathx"
	"strings"
)

func genInsert(table Table) (string, error) {
	expressions := make([]string, 0)

	var (
		tableField string
		args       = "[]interface{}{"
	)
	camel := table.Name.ToCamel()
	name := stringx.From(camel).Untitle()[:1]
	var (
		count int
		auto  int
	)
	for _, field := range table.Fields {
		camel := util.SafeString(field.Name.ToCamel())
		if table.isIgnoreColumns(field.Name.Source()) {
			count--
			auto++
			continue
		}

		if field.Name.Source() == table.PrimaryKey.Name.Source() {
			if table.PrimaryKey.AutoIncrement {
				if count != 0 {
					count--
				}
				auto++
				continue
			}
		}

		count += 1
		expressions = append(expressions, "?")

		if count == len(table.Fields)-auto {
			tableField = tableField + camel
			args = args + name + "." + camel
			args = args + "}"
		} else {
			tableField = tableField + camel + ","
			args = args + name + "." + camel + ","
		}

	}
	text, err := pathx.LoadTemplate(category, insertTemplateFile, template.Insert)
	if err != nil {
		return "", err
	}

	output, err := util.With("insert").
		Parse(text).
		Execute(map[string]interface{}{
			"upperStartCamelObject": camel,
			"lowerStartCamelObject": tableField,
			"expression":            strings.Join(expressions, ", "),
			"tableField":            tableField,
			"tableName":             table.Name.Source(),
			"name":                  name,
			"fields":                args,
		})
	if err != nil {
		return "", err
	}
	return output.String(), nil
}
