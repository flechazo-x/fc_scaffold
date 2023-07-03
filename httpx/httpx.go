package httpx

import (
	"fc_scaffold/config"
	"fc_scaffold/gen/activities"
	"fc_scaffold/htmlx"
	"fc_scaffold/sqltogo"
	"fc_scaffold/static"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/logrusorgru/aurora"
	"net/http"
	"strconv"
	"strings"
)

func Http() {
	gin.SetMode(gin.ReleaseMode)
	// 创建默认的 Gin 引擎
	r := gin.Default()

	//box := packr.New("htmlx", "../htmlx")
	//htmlBytes, err := box.Find("index.html")
	//if err != nil {
	//	panic(err)
	//}

	r.GET("/index", func(c *gin.Context) {
		c.Status(http.StatusOK)
		c.Header("Content-Type", "text/html; charset=utf-8")
		c.String(http.StatusOK, htmlx.SqlToGo)
		//http.ServeContent(c.Writer, c.Request, "index.html", time.Time{}, bytes.NewReader(htmlBytes))
	})
	handleScaffold(r) //处理脚手架
	sqlToGo(r)        //处理sql转go
	jumpScaffold(r)   //处理页面跳转到脚手架
	fmt.Println(aurora.BrightYellow("fc_scaffold: http://127.0.0.1:8080/index"))
	err := r.Run(":8080")
	if err != nil {
		panic(err)
	}
}

func handleScaffold(r *gin.Engine) {
	var cfgs struct {
		PkgName             string `json:"PkgName,omitempty"`             //包名,结构体的包名
		Id                  string `json:"Id,omitempty"`                  //活动或者机器id
		StructName          string `json:"StructName,omitempty"`          //结构体名
		Author              string `json:"Author,omitempty"`              //作者
		OutputPath          string `json:"OutputPath,omitempty"`          //输出路径 填写工作目录绝对路径
		IsSupportTasks      int    `json:"IsSupportTasks,omitempty"`      //是否支持任务
		IsPaySupport        int    `json:"IsPaySupport,omitempty"`        //是否支持支付
		ProtoPath           string `json:"ProtoPath,omitempty"`           //proto路径
		ActivityChineseName string `json:"ActivityChineseName,omitempty"` //活动中文名字
		IsTable             int    `json:"IsTable,omitempty"`             //是否生成dao层文件
		TableName           string `json:"TableName,omitempty"`           //表名字
	}

	// POST 请求处理
	r.POST("/fc_scaffold", func(c *gin.Context) {
		err := c.BindJSON(&cfgs)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		}
		tableName := strings.TrimSpace(cfgs.TableName)
		tableNameList := strings.Split(tableName, ",")
		cfg := &config.Config{
			PkgName:             cfgs.PkgName,
			Id:                  cfgs.Id,
			StructName:          cfgs.StructName,
			Author:              cfgs.Author,
			OutputPath:          cfgs.OutputPath,
			IsSupportTasks:      cfgs.IsSupportTasks,
			IsPaySupport:        cfgs.IsPaySupport,
			ProtoPath:           cfgs.ProtoPath,
			ActivityChineseName: cfgs.ActivityChineseName,
			IsTable:             cfgs.IsTable,
			TableName:           tableNameList,
		}
		err = genFile(cfg)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"result": fmt.Sprintf("出错啦:%s", err.Error()),
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"result": "生成成功",
			})
		}
	})
}

func jumpScaffold(r *gin.Engine) {
	r.GET("/scaffold", func(c *gin.Context) {
		c.Status(http.StatusOK)
		c.Header("Content-Type", "text/html; charset=utf-8")
		c.String(http.StatusOK, htmlx.HtmlStr)
		//http.ServeContent(c.Writer, c.Request, "index.html", time.Time{}, bytes.NewReader(htmlBytes))
	})
}

func sqlToGo(r *gin.Engine) {
	var cfgs struct {
		Filepath  string `json:"filepath,omitempty"`  //输出文件路径
		TableName string `json:"TableName,omitempty"` //表名字
	}

	// POST 请求处理
	r.POST("/sqltogo", func(c *gin.Context) {
		err := c.BindJSON(&cfgs)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		}
		tableName := strings.TrimSpace(cfgs.TableName)
		tableNameList := strings.Split(tableName, ",")
		err = sqltogo.Run(cfgs.Filepath, tableNameList...)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"result": fmt.Sprintf("出错啦:%s", err.Error()),
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"result": "生成成功",
			})
		}
	})
}

func genFile(cfg *config.Config) error {
	var err error
	err = cfg.Check()
	if err != nil {
		return err
	}
	err = activities.CreateFile(cfg)
	if err != nil {
		return err
	}
	if cfg.IsTable == 1 {
		if err1 := sqltogo.Run(cfg.OutputPath+static.SqlGoPath, cfg.TableName...); err1 != nil {
			return err1
		}
	}
	return nil
}

func Run() *config.Config {
	var (
		options string
		c       = new(config.Config)
	)
	config.TimeDelayPrint(static.Welcome, 0, aurora.BrightCyan)
	config.TimeDelayPrint(static.ProtoHint, 5, aurora.BgRed)
	config.TimeDelayPrint(static.Options, 0, aurora.BrightGreen)
	_, err := fmt.Scanln(&options)
	if err != nil {
		panic(err)
	}
	optionsInt, err := strconv.Atoi(options)
	if err != nil {
		panic(err)
	}
	if optionsInt == 0 {
		c = config.EntryType()
	} else if optionsInt == 1 {
		c = config.Interactive()
	} else if optionsInt == 2 {
		// 创建默认的 Gin 引擎
		// 启动 HTTP 服务器
		Http()
	}
	config.TimeDelayPrint(fmt.Sprintf("The configuration you entered is:------>%+v\n", c), 5, aurora.BgMagenta)
	return c
}
