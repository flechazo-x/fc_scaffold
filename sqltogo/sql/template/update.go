package template

const (
	// Update defines a template for generating update codes
	Update = `
func ({{.name}} *{{.upperStartCamelObject}}) Update(ctx *trace.Context,  kv db.KeyValue, where db.KeyValue) (sql.Result, error) {
	fields, args := kv.Transport()
	fieldsW, fieldsValueW := where.Transport()
	args = append(args, fieldsValueW...)
	query := fmt.Sprintf("UPDATE %s SET %s WHERE %s", "{{.tableName}}", strings.Join(fields, db.FieldsSplit), strings.Join(fieldsW, db.WhereSplitAnd))
	return db.Update(ctx, query, args)
}
`
)
