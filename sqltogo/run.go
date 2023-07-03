package sqltogo

import (
	"database/sql"
	"fc_scaffold/sqltogo/sql/config"
	"fc_scaffold/sqltogo/sql/gen"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/logrusorgru/aurora"
	"os"
)

func Run(outputPath string, tables ...string) error {
	fileName, err := GenFile(tables...)
	if err != nil {
		return err
	}
	defer os.Remove(fileName) //删除文件
	if err = SqlToGo(fileName, outputPath); err != nil {
		return err
	}
	return nil
}

func GenFile(tables ...string) (string, error) {
	db, err := sql.Open("mysql", "devpc:dev55571965@tcp(192.168.1.8:3306)/slots2021")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	var (
		ddlStatements []string
	)
	for _, tableName := range tables {
		rows, err := db.Query(fmt.Sprintf("SHOW CREATE TABLE %s", tableName))
		if err != nil {
			return "", fmt.Errorf("表不存在:%s,err:%s", tableName, err.Error())
		}
		defer rows.Close()

		for rows.Next() {
			var ddlStatement string
			err = rows.Scan(&tableName, &ddlStatement)
			if err != nil {
				return "", fmt.Errorf("scan err:%s", err.Error())
			}
			ddlStatements = append(ddlStatements, ddlStatement)
		}
	}

	// 将DDL语句保存到临时文件
	tmpFile, err := os.CreateTemp("", "ddl-*.sql")
	if err != nil {
		panic(err.Error())
	}
	defer tmpFile.Close()

	for _, ddlStatement := range ddlStatements {
		_, err = tmpFile.WriteString(ddlStatement + ";\n")
		if err != nil {
			return "", fmt.Errorf("临时文件写入本地失败 err:%s", err.Error())
		}
	}
	return tmpFile.Name(), nil
}

func SqlToGo(sqlFileName, outputPath string) error {
	_ = gen.Clean()

	fmt.Println(aurora.BgRed(fmt.Sprintf("sql input directory------------------->%s", sqlFileName)))

	g, err := gen.NewDefaultGenerator(outputPath, &config.Config{
		NamingFormat: namingFormat,
	})
	if err != nil {
		return fmt.Errorf("生成器创建失败：%s", err.Error())
	}
	err = g.StartFromDDL(sqlFileName, false, dataBase)
	if err != nil {
		return fmt.Errorf("sql创建失败：%s,请检查表中是否有联合主键", err.Error())
	}
	return nil
}

const (
	namingFormat = "GoZero" //命名规则,驼峰命名
	dataBase     = "slots2021"
) //命名规则,驼峰命名
//levelcharge
