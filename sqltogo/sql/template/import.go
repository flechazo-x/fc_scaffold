package template

const (
	// Imports defines an import template for model in cache case
	Imports = `import (
	"casino/module/business/dao/db"
	"database/sql"
	"fmt"
	"strings"
	{{if .time}}"time"{{end}}

   "github.com/aaagame/common/middleware/trace"
)
`
)
