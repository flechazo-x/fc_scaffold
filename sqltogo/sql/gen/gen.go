/**
 * @Description
 * @Author 张盛钢
 * @Date 2022/11/8 13:40
 **/
package gen

import (
	"bytes"
	"fc_scaffold/sqltogo/ddl-parser/console"
	"fc_scaffold/sqltogo/sql/config"
	"fc_scaffold/sqltogo/sql/model"
	"fc_scaffold/sqltogo/sql/parser"
	"fc_scaffold/sqltogo/sql/stringx"
	"fc_scaffold/sqltogo/sql/template"
	"fc_scaffold/sqltogo/sql/util"
	"fc_scaffold/sqltogo/sql/util/format"
	"fc_scaffold/sqltogo/sql/util/pathx"
	"fmt"
	"github.com/logrusorgru/aurora"
	"os"
	"path/filepath"
	"strings"
)

const pwd = "."

type (
	defaultGenerator struct {
		console.Console
		// source string
		dir           string
		pkg           string
		cfg           *config.Config
		isPostgreSql  bool
		ignoreColumns []string
	}

	// Option defines a function with argument defaultGenerator
	Option func(generator *defaultGenerator)

	code struct {
		importsCode string
		varsCode    string
		typesCode   string
		insertCode  string
		findCode    []string
		findBatch   string
		findRow     string
		updateCode  string
		deleteCode  string
	}

	codeTuple struct {
		modelCode string
		//modelCustomCode string
	}
)

// NewDefaultGenerator creates an instance for defaultGenerator
func NewDefaultGenerator(dir string, cfg *config.Config, opt ...Option) (*defaultGenerator, error) {
	if dir == "" {
		dir = pwd
	}
	dirAbs, err := filepath.Abs(dir)
	if err != nil {
		return nil, err
	}

	dir = dirAbs
	pkg := util.SafeString(filepath.Base(dirAbs))
	err = pathx.MkdirIfNotExist(dir)
	if err != nil {
		return nil, err
	}

	generator := &defaultGenerator{dir: dir, cfg: cfg, pkg: pkg}
	var optionList []Option
	optionList = append(optionList, newDefaultOption())
	optionList = append(optionList, opt...)
	for _, fn := range optionList {
		fn(generator)
	}

	return generator, nil
}

// WithConsoleOption creates a console option.
func WithConsoleOption(c console.Console) Option {
	return func(generator *defaultGenerator) {
		generator.Console = c
	}
}

// WithIgnoreColumns ignores the columns while insert or update rows.
func WithIgnoreColumns(ignoreColumns []string) Option {
	return func(generator *defaultGenerator) {
		generator.ignoreColumns = ignoreColumns
	}
}

// WithPostgreSql marks  defaultGenerator.isPostgreSql true.
func WithPostgreSql() Option {
	return func(generator *defaultGenerator) {
		generator.isPostgreSql = true
	}
}

func newDefaultOption() Option {
	return func(generator *defaultGenerator) {
		generator.Console = console.NewColorConsole()
	}
}

func (g *defaultGenerator) StartFromDDL(filename string, strict bool, database string) error {
	modelList, err := g.genFromDDL(filename, strict, database)
	if err != nil {
		return err
	}

	return g.createFile(modelList)
}

func (g *defaultGenerator) StartFromInformationSchema(tables map[string]*model.Table, strict bool) error {
	m := make(map[string]*codeTuple)
	for _, each := range tables {
		table, err := parser.ConvertDataType(each, strict)
		if err != nil {
			return err
		}

		code, err := g.genModel(*table)
		if err != nil {
			return err
		}

		m[table.Name.Source()] = &codeTuple{
			modelCode: code,
		}
	}

	return g.createFile(m)
}

func (g *defaultGenerator) createFile(modelList map[string]*codeTuple) error {
	dirAbs, err := filepath.Abs(g.dir)
	if err != nil {
		return err
	}

	g.dir = dirAbs
	g.pkg = util.SafeString(filepath.Base(dirAbs))
	err = pathx.MkdirIfNotExist(dirAbs)
	if err != nil {
		return err
	}

	for tableName, codes := range modelList {
		tn := stringx.From(tableName)
		modelFilename, err := format.FileNamingFormat(g.cfg.NamingFormat,
			tn.Source())
		if err != nil {
			return err
		}

		name := util.SafeString(modelFilename) + ".go"
		filename := filepath.Join(dirAbs, name)
		err = os.WriteFile(filename, []byte(codes.modelCode), os.ModePerm)
		if err != nil {
			return err
		}
		fmt.Println(aurora.BgRed(fmt.Sprintf("dao output directory-------------------->%s", filename)))
		if pathx.FileExists(filename) {
			fmt.Printf("%s already exists, ignored.\n", name)
			continue
		}
	}

	fmt.Println(aurora.BrightYellow("Done."))
	return nil
}

// ret1: key-table name,value-code
func (g *defaultGenerator) genFromDDL(filename string, strict bool, database string) (
	map[string]*codeTuple, error,
) {
	m := make(map[string]*codeTuple)
	tables, err := parser.Parse(filename, database, strict)
	if err != nil {
		return nil, err
	}

	for _, e := range tables {
		code, err := g.genModel(*e)
		if err != nil {
			return nil, err
		}

		m[e.Name.Source()] = &codeTuple{
			modelCode: code,
		}
	}

	return m, nil
}

// Table defines mysql table
type Table struct {
	parser.Table
	PrimaryCacheKey        Key
	UniqueCacheKey         []Key
	ContainsUniqueCacheKey bool
	ignoreColumns          []string
}

func (t Table) isIgnoreColumns(columnName string) bool {
	for _, v := range t.ignoreColumns {
		if v == columnName {
			return true
		}
	}
	return false
}

// 生成model
func (g *defaultGenerator) genModel(in parser.Table) (string, error) {
	if len(in.PrimaryKey.Name.Source()) == 0 {
		return "", fmt.Errorf("table %s: missing primary key", in.Name.Source())
	}

	primaryKey, uniqueKey := genCacheKeys(in)

	var table Table
	table.Table = in
	table.PrimaryCacheKey = primaryKey
	table.UniqueCacheKey = uniqueKey
	table.ContainsUniqueCacheKey = len(uniqueKey) > 0
	table.ignoreColumns = g.ignoreColumns

	importsCode, err := genImports(table, in.ContainsTime()) //导入代码
	if err != nil {
		return "", err
	}

	varsCode, err := genVars(table) //字段参数const
	if err != nil {
		return "", err
	}

	insertCode, err := genInsert(table) //插入代码
	if err != nil {
		return "", err
	}

	findCode := make([]string, 0)
	findOneCode, err := genFindOne(table) //查询一行代码
	if err != nil {
		return "", err
	}
	findCode = append(findCode, findOneCode)
	findBatchCode, err := genFindBatch(table) // 批量查询
	if err != nil {
		return "", err
	}
	findRowCode, err := genFindRows(table)
	if err != nil {
		return "", err
	}

	updateCode, err := genUpdate(table) //更新代码
	if err != nil {
		return "", err
	}

	deleteCode, err := genDelete(table) //删除代码
	if err != nil {
		return "", err
	}

	typesCode, err := genTypes(table) //生成go结构体
	if err != nil {
		return "", err
	}

	code := &code{
		importsCode: importsCode,
		varsCode:    varsCode,
		typesCode:   typesCode,
		insertCode:  insertCode,
		findCode:    findCode,
		findRow:     findRowCode,
		findBatch:   findBatchCode,
		updateCode:  updateCode,
		deleteCode:  deleteCode,
	}

	output, err := g.executeModel(code)
	if err != nil {
		return "", err
	}

	return output.String(), nil
}

func (g *defaultGenerator) genModelCustom(in parser.Table, withCache bool) (string, error) {
	text, err := pathx.LoadTemplate(category, modelCustomTemplateFile, template.ModelCustom)
	if err != nil {
		return "", err
	}

	t := util.With("model-custom").
		Parse(text).
		GoFmt(true)
	output, err := t.Execute(map[string]interface{}{
		"pkg":                   g.pkg,
		"withCache":             withCache,
		"upperStartCamelObject": in.Name.ToCamel(),
		"lowerStartCamelObject": stringx.From(in.Name.ToCamel()).Untitle(),
	})
	if err != nil {
		return "", err
	}

	return output.String(), nil
}

func (g *defaultGenerator) executeModel(code *code) (*bytes.Buffer, error) {
	text, err := pathx.LoadTemplate(category, modelGenTemplateFile, template.ModelGen)
	if err != nil {
		return nil, err
	}
	t := util.With("model").
		Parse(text).
		GoFmt(true)
	output, err := t.Execute(map[string]interface{}{
		"pkg":       g.pkg,
		"imports":   code.importsCode,
		"vars":      code.varsCode,
		"types":     code.typesCode,
		"insert":    code.insertCode,
		"find":      strings.Join(code.findCode, "\n"),
		"findBatch": code.findBatch,
		"findRow":   code.findRow,
		"update":    code.updateCode,
		"delete":    code.deleteCode,
	})
	if err != nil {
		return nil, err
	}
	return output, nil
}

func wrapWithRawString(v string, postgreSql bool) string {
	if postgreSql {
		return v
	}

	if v == "`" {
		return v
	}

	if !strings.HasPrefix(v, "`") {
		v = "`" + v
	}

	if !strings.HasSuffix(v, "`") {
		v = v + "`"
	} else if len(v) == 1 {
		v = v + "`"
	}

	return v
}
