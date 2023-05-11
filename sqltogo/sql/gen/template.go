package gen

import (
	"fc_scaffold/sqltogo/sql/template"
	"fc_scaffold/sqltogo/sql/util/pathx"
	"fmt"
)

const (
	category              = "model"
	deleteTemplateFile    = "delete.tpl"
	fieldTemplateFile     = "field.tpl"
	findOneTemplateFile   = "find-one.tpl"
	findBatchTemplateFile = "find-batch.tpl"
	findRowTemplateFile   = "find-row.tpl"

	importsTemplateFile     = "import.tpl"
	insertTemplateFile      = "insert.tpl"
	modelGenTemplateFile    = "model-gen.tpl"
	modelCustomTemplateFile = "model.tpl"
	tagTemplateFile         = "tag.tpl"
	typesTemplateFile       = "types.tpl"
	updateTemplateFile      = "update.tpl"
)

var templates = map[string]string{
	deleteTemplateFile:      template.Delete,
	fieldTemplateFile:       template.Field,
	findOneTemplateFile:     template.FindOne,
	findBatchTemplateFile:   template.FindBatch,
	findRowTemplateFile:     template.FindRow,
	importsTemplateFile:     template.Imports,
	insertTemplateFile:      template.Insert,
	modelGenTemplateFile:    template.ModelGen,
	modelCustomTemplateFile: template.ModelCustom,
	tagTemplateFile:         template.Tag,
	typesTemplateFile:       template.Types,
	updateTemplateFile:      template.Update,
}

// Category returns model const value
func Category() string {
	return category
}

// Clean deletes all template files
func Clean() error {
	return pathx.Clean(category)
}

// GenTemplates creates template files if not exists
func GenTemplates() error {
	return pathx.InitTemplates(category, templates)
}

// RevertTemplate reverts the deleted template files
func RevertTemplate(name string) error {
	content, ok := templates[name]
	if !ok {
		return fmt.Errorf("%s: no such file name", name)
	}

	return pathx.CreateTemplate(category, name, content)
}

// Update provides template clean and init
func Update() error {
	err := Clean()
	if err != nil {
		return err
	}

	return pathx.InitTemplates(category, templates)
}
