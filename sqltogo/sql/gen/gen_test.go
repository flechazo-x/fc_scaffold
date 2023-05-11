package gen

import (
	_ "embed"
	"github.com/stretchr/testify/assert"
	"sqlToGo/sql/config"
	"sqlToGo/sql/util/pathx"

	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
)

//go:embed testdata/user.sql
var source string

func TestNamingModel(t *testing.T) {

	_ = Clean()

	sqlFile := filepath.Join(pathx.MustTempDir(), "tmp.sql")
	err := ioutil.WriteFile(sqlFile, []byte(source), 0o777)
	assert.Nil(t, err)

	dir, _ := filepath.Abs("./testmodel")
	camelDir := filepath.Join(dir, "camel")
	//snakeDir := filepath.Join(dir, "snake")
	//defer func() {
	//	_ = os.RemoveAll(dir)
	//}()
	g, err := NewDefaultGenerator(camelDir, &config.Config{
		NamingFormat: "GoZero",
	})
	assert.Nil(t, err)

	err = g.StartFromDDL(sqlFile, true, "go_zero")
	assert.Nil(t, err)
	assert.True(t, func() bool {
		_, err := os.Stat(filepath.Join(camelDir, "TestUserModel.go"))
		return err == nil
	}())
	//g, err = NewDefaultGenerator(snakeDir, &config.Config{
	//	NamingFormat: "go_zero",
	//})
	//assert.Nil(t, err)
	//
	//err = g.StartFromDDL(sqlFile, true, false, "go_zero")
	//assert.Nil(t, err)
	//assert.True(t, func() bool {
	//	_, err := os.Stat(filepath.Join(snakeDir, "test_user_model.go"))
	//	return err == nil
	//}())
	//log.LErrorf("")
}
