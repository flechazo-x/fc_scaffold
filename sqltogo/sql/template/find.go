package template

const (
	// FindOne defines find row by id.
	FindOne = `
func ({{.name}} *{{.upperStartCamelObject}}) SelectOne(ctx *trace.Context, fields []string, where db.KeyValue) *sql.Row {
	fieldsW, args := where.Transport()
	query := fmt.Sprintf("SELECT %s FROM %s WHERE %s", strings.Join(fields, db.FieldsSplit), "{{.tableName}}",
		strings.Join(fieldsW, db.WhereSplitAnd))
	return db.Select(ctx, query, args)
}
`

	FindBatch = `
// SelectBatch ！！！注意：上层调用需要手动释放defer rows.Close()
func ({{.name}} *{{.upperStartCamelObject}}) SelectBatch(ctx *trace.Context, fields []string, where db.KeyValue) *sql.Rows {
	fieldsW, args := where.Transport()
	query := fmt.Sprintf("SELECT %s FROM %s WHERE %s", strings.Join(fields, db.FieldsSplit), "{{.tableName}}", strings.Join(fieldsW, db.WhereSplitAnd))
	return db.SelectBatch(ctx, query, args)
}
`
	FindRow = `
func ({{.name}} *{{.upperStartCamelObject}}) ScanRow(row *sql.Row) error {
	return row.Scan({{.fields}})
}
`
)
