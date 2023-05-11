package template

const (
	// Delete defines a delete template
	Delete = `
func ({{.name}} *{{.upperStartCamelObject}}) Delete(ctx *trace.Context, where db.KeyValue) (sql.Result, error) {
	fields, args := where.Transport()
	query := fmt.Sprintf("DELETE FROM %s WHERE %s","{{.tableName}}", strings.Join(fields, db.WhereSplitAnd))
	return db.Delete(ctx, query, args)
}
`
)
