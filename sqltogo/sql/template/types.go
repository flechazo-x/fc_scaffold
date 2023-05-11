package template

// Types defines a template for types in model.
const Types = `
type {{.upperStartCamelObject}} struct {
		db.DefaultDBINative
		{{.fields}}
	}
`
