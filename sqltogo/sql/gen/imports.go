package gen

import (
	"fc_scaffold/sqltogo/sql/template"
	"fc_scaffold/sqltogo/sql/util"
	"fc_scaffold/sqltogo/sql/util/pathx"
)

func genImports(table Table, timeImport bool) (string, error) {
	text, err := pathx.LoadTemplate(category, importsTemplateFile, template.Imports)
	if err != nil {
		return "", err
	}

	output, err := util.With("import").Parse(text).Execute(map[string]interface{}{
		"time":       timeImport,
		"containsPQ": table.ContainsPQ,
		"data":       table,
	})
	if err != nil {
		return "", err
	}
	return output.String(), nil
}
