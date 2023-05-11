package parser

import (
	_ "embed"
	"io/ioutil"
	"path/filepath"
	"sqlToGo/sql/util"
	"sqlToGo/sql/util/pathx"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParsePlainText(t *testing.T) {
	sqlFile := filepath.Join(pathx.MustTempDir(), "tmp.sql")
	err := ioutil.WriteFile(sqlFile, []byte("plain text"), 0o777)
	assert.Nil(t, err)

	_, err = Parse(sqlFile, "go_zero", false)
	assert.NotNil(t, err)
}

func TestParseSelect(t *testing.T) {
	sqlFile := filepath.Join(pathx.MustTempDir(), "tmp.sql")
	err := ioutil.WriteFile(sqlFile, []byte("select * from user"), 0o777)
	assert.Nil(t, err)

	tables, err := Parse(sqlFile, "go_zero", false)
	assert.Nil(t, err)
	assert.Equal(t, 0, len(tables))
}

//go:embed testdata/user.sql
var user string

func TestParseCreateTable(t *testing.T) {
	sqlFile := filepath.Join(pathx.MustTempDir(), "tmp.sql")
	err := ioutil.WriteFile(sqlFile, []byte(user), 0o777)
	assert.Nil(t, err)

	tables, err := Parse(sqlFile, "go_zero", false)
	assert.Equal(t, 1, len(tables))
	table := tables[0]
	assert.Nil(t, err)
	assert.Equal(t, "test_user", table.Name.Source())
	assert.Equal(t, "id", table.PrimaryKey.Name.Source())
	assert.Equal(t, true, table.ContainsTime())
	assert.Equal(t, 2, len(table.UniqueIndex))
	assert.True(t, func() bool {
		for _, e := range table.Fields {
			if e.Comment != util.TrimNewLine(e.Comment) {
				return false
			}
		}

		return true
	}())
}
