package pathx

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReadLink(t *testing.T) {
	dir, err := ioutil.TempDir("", "slots")
	assert.Nil(t, err)
	symLink := filepath.Join(dir, "test")
	pwd, err := os.Getwd()
	assertError(err, t)

	err = os.Symlink(pwd, symLink)
	assertError(err, t)
}

func assertError(err error, t *testing.T) {
	if err != nil {
		t.Fatal(err)
	}
}
