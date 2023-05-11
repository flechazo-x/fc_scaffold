package gen

import (
	"fc_scaffold/sqltogo/sql/stringx"
	"fc_scaffold/sqltogo/sql/template"
	"fc_scaffold/sqltogo/sql/util"
	"fc_scaffold/sqltogo/sql/util/pathx"
)

func genFindBatch(table Table) (string, error) {
	camel := table.Name.ToCamel()
	text, err := pathx.LoadTemplate(category, findOneTemplateFile, template.FindBatch)
	if err != nil {
		return "", err
	}

	output, err := util.With("findBatch").
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

func genFindRows(table Table) (string, error) {
	var tableField string
	camel := table.Name.ToCamel()
	name := stringx.From(camel).Untitle()[:1]
	var count int
	for _, field := range table.Fields {
		camel := util.SafeString(field.Name.ToCamel())
		if table.isIgnoreColumns(field.Name.Source()) {
			continue
		}

		if field.Name.Source() == table.PrimaryKey.Name.Source() {
			if table.PrimaryKey.AutoIncrement {
				continue
			}
		}

		count += 1

		if count == len(table.Fields) {
			tableField = tableField + "&" + name + "." + camel
		} else {
			tableField = tableField + "&" + name + "." + camel + ","
		}

	}

	text, err := pathx.LoadTemplate(category, findRowTemplateFile, template.FindRow)
	if err != nil {
		return "", err
	}

	output, err := util.With("findRow").
		Parse(text).
		Execute(map[string]interface{}{
			"upperStartCamelObject": camel,
			"fields":                tableField,
			"name":                  name,
		})
	if err != nil {
		return "", err
	}
	return output.String(), nil
}
