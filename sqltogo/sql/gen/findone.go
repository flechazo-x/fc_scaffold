package gen

import (
	"fc_scaffold/sqltogo/sql/stringx"
	"fc_scaffold/sqltogo/sql/template"
	"fc_scaffold/sqltogo/sql/util"
	"fc_scaffold/sqltogo/sql/util/pathx"
)

func genFindOne(table Table) (string, error) {
	camel := table.Name.ToCamel()
	text, err := pathx.LoadTemplate(category, findOneTemplateFile, template.FindOne)
	if err != nil {
		return "", err
	}

	output, err := util.With("findOne").
		Parse(text).
		Execute(map[string]interface{}{
			"upperStartCamelObject": camel,
			"tableName":             table.Name.Source(),
			"name":                  stringx.From(camel).Untitle()[:1],
		})
	if err != nil {
		return "", err
	}
	return output.String(), nil
}
