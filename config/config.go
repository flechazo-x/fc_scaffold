package config

import (
	"bufio"
	"errors"
	"fc_scaffold/static"
	"flag"
	"fmt"
	"github.com/logrusorgru/aurora"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

type (
	Config struct {
		PkgName             string   `json:"PkgName,omitempty"`             //包名,结构体的包名
		Id                  string   `json:"Id,omitempty"`                  //活动或者机器id
		StructName          string   `json:"StructName,omitempty"`          //结构体名
		Author              string   `json:"Author,omitempty"`              //作者
		OutputPath          string   `json:"OutputPath,omitempty"`          //输出路径 填写工作目录绝对路径
		IsSupportTasks      int      `json:"IsSupportTasks,omitempty"`      //是否支持任务
		IsPaySupport        int      `json:"IsPaySupport,omitempty"`        //是否支持支付
		ProtoPath           string   `json:"ProtoPath,omitempty"`           //proto路径
		ActivityChineseName string   `json:"ActivityChineseName,omitempty"` //活动中文名字
		IsTable             int      `json:"IsTable,omitempty"`             //是否生成dao层文件
		TableName           []string `json:"TableName,omitempty"`           //表名字
	}
	ProtoData struct {
		Req  []string
		Resp []string
	}
)

var (
	InstructionSet = []int32{8, 1, 2, 3, 4, 5, 9, 6, 7, 10, 11}
	Table          = map[int32]func(){
		1: func() { fmt.Println(aurora.BrightYellow("Please enter the pkgName. (Package name):")) },
		2: func() { fmt.Println(aurora.BrightYellow("Please input the Id. (Activity ID 7001...):")) },
		3: func() { fmt.Println(aurora.BrightYellow("Please enter structName (Struct name):")) },
		4: func() { fmt.Println(aurora.BrightYellow("Please enter author (Author name):")) },
		5: func() {
			fmt.Println(aurora.BrightYellow("Please enter outputPath working directory. (Beta absolute path, for example, mine:E:\\GoProject\\src\\work\\server\\branches\\beta):"))
		},
		6: func() {
			fmt.Println(aurora.BrightYellow("Please enter isSupportTasks (Does it support tasks? 0 No, 1 Yes):"))
		},
		7: func() {
			fmt.Println(aurora.BrightYellow("Please enter IsPaySupport (Do you support payment? 0 No, 1 Yes):"))
		},
		8: func() {
			fmt.Println(aurora.BrightYellow("Please enter proto working directory (Beta absolute path E:\\GoProject\\src\\work\\docs\\branches\\beta):"))
		},
		9: func() {
			fmt.Println(aurora.BrightYellow("Please enter activity chinese name (Activity chinese name):"))
		},
		10: func() {
			fmt.Println(aurora.BrightYellow("Whether to generate dao layer files (0 No, 1 Yes):"))
		},
		11: func() {
			fmt.Println(aurora.BrightYellow("Please enter a data table name (Supports multiple,separated by commas):"))
		},
	}
)

func (c *Config) Check() error {
	var err error
	filename := filepath.Join(c.ProtoPath, static.GetProtoPath(c.Id))
	_, err = os.Stat(filename)
	if err != nil {
		return err
	}

	if !filepath.IsAbs(c.OutputPath) {
		return errors.New("outputPath no abs")
	}
	if c.StructName == "" {
		return errors.New("structName is empty")
	}

	jsonName := filepath.Join(c.ProtoPath, static.GetProtoJsonPath(c.Id))
	_, err = os.Stat(jsonName)
	if err != nil {
		return err
	}

	c.StructName = strings.Title(c.StructName)
	c.PkgName = strings.ToLower(c.PkgName)
	return nil
}

func PrintErr(err error) {
	if err != nil {
		panic(err.Error())
	}
}

// 入参式
func EntryType() *Config {
	c := new(Config)
	var tableStr string
	flag.StringVar(&c.PkgName, "PkgName", "", "")
	flag.StringVar(&c.Id, "Id", "", "")
	flag.StringVar(&c.StructName, "StructName", "", "")
	flag.StringVar(&c.Author, "Author", "", "")
	flag.StringVar(&c.OutputPath, "OutputPath", "", "")
	flag.IntVar(&c.IsSupportTasks, "IsSupportTasks", 0, "")
	flag.IntVar(&c.IsPaySupport, "IsPaySupport", 0, "")
	flag.StringVar(&c.ProtoPath, "ProtoPath", "", "")
	flag.StringVar(&c.ActivityChineseName, "ActivityChineseName", "", "")
	flag.IntVar(&c.IsTable, "IsTable", 0, "")
	flag.StringVar(&tableStr, "TableName", "", "")
	flag.Parse()
	c.TableName = strings.Split(tableStr, ",")
	return c
}

// 交互式
func Interactive() *Config {
	scanner := bufio.NewReader(os.Stdin)
	var (
		result    []string
		tableName []string
	)
	for _, i := range InstructionSet {
		v, ok := Table[i]
		if !ok {
			continue
		}
		if i != 11 {
			v()
			input, _ := scanner.ReadString('\n')
			result = append(result, strings.TrimSpace(input))
		}
		if i == 11 {
			if result[len(result)-1] == "0" {
				continue
			}
			v()
			input, _ := scanner.ReadString('\n')
			tableName = strings.Split(input, ",")
		}
	}
	st, err := strconv.Atoi(result[7])
	PrintErr(err)
	ps, err := strconv.Atoi(result[8])
	PrintErr(err)
	it, err := strconv.Atoi(result[9])
	PrintErr(err)
	c := &Config{
		PkgName:             result[1],
		Id:                  result[2],
		StructName:          result[3],
		Author:              result[4],
		OutputPath:          result[5],
		IsSupportTasks:      st,
		IsPaySupport:        ps,
		ProtoPath:           result[0],
		ActivityChineseName: result[6],
		IsTable:             it,
		TableName:           tableName,
	}
	return c
}

// TimeDelayPrint 延时打印
func TimeDelayPrint(str string, tm int64, runfunc func(arg interface{}) aurora.Value) {
	chars := strings.Split(str, "")
	for _, char := range chars {
		if runfunc == nil {
			fmt.Print(char)
		} else {
			fmt.Print(runfunc(char))
		}

		time.Sleep(time.Duration(tm) * time.Millisecond) // 控制打印速度，500毫秒为例
	}
	fmt.Println()
}
