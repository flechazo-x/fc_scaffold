package gen

import (
	"fc_scaffold/sqltogo/sql/template"
	"fc_scaffold/sqltogo/sql/util"
	"fc_scaffold/sqltogo/sql/util/pathx"
)

func genTypes(table Table) (string, error) {
	fields := table.Fields
	fieldsString, err := genFields(table, fields)
	if err != nil {
		return "", err
	}

	text, err := pathx.LoadTemplate(category, typesTemplateFile, template.Types)
	if err != nil {
		return "", err
	}

	output, err := util.With("types").
		Parse(text).
		Execute(map[string]interface{}{
			"upperStartCamelObject": table.Name.ToCamel(),
			"fields":                fieldsString,
		})
	if err != nil {
		return "", err
	}

	return output.String(), nil
}
