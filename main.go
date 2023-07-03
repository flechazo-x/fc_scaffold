package main

import (
	"fc_scaffold/gen/activities"
	"fc_scaffold/httpx"
	"fc_scaffold/sqltogo"
	"fc_scaffold/static"
	"fmt"
	"github.com/logrusorgru/aurora"
)

func main() {
	var err error
	c := httpx.Run()
	err = c.Check()
	if err != nil {
		panic(err)
	}

	err = activities.CreateFile(c)
	if err != nil {
		panic(err)
	}
	if c.IsTable == 1 {
		fmt.Println(aurora.BgBlue("generate table ~~~"))
		if err1 := sqltogo.Run(c.OutputPath+static.SqlGoPath, c.TableName...); err1 != nil {
			panic(err1)
		}
	}
	fmt.Println(aurora.BgBlue("File generated successfully ~~~"))

}

//E:\GoProject\src\work\test_server\branches\beta\module\business\dao\user
//venture,levelcharge
