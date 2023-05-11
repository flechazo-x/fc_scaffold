package template

const (
	Insert = `
func ({{.name}} *{{.upperStartCamelObject}}) Insert(ctx *trace.Context, _ db.KeyValue) (sql.Result, error) {
	query := fmt.Sprintf("INSERT INTO %s ({{.tableField}})  VALUES ({{.expression}})","{{.tableName}}")
	args := {{.fields}}
	return db.Insert(ctx, query, args)
}
`
)
