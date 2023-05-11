package activities

import (
	"bufio"
	"errors"
	"fc_scaffold/config"
	"fc_scaffold/parser"
	"fc_scaffold/static"
	"fc_scaffold/template/activity"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"time"
)

func GenModel(c *config.Config) (string, error) {
	if c == nil {
		return "", errors.New("config is nil")
	}
	var prefix = "v" //默认是v
	if len(c.StructName) > 0 {
		prefix = strings.ToLower(c.StructName[0:1])
	}

	var (
		text      = activity.Model //内容
		importStr = activity.Import
	)
	if c.IsPaySupport == 1 {
		text = activity.Model + activity.Support
		importStr = activity.Import1
	}

	output, err := parser.With("text").Parse(text).GoFmt(true).Execute(map[string]interface{}{
		"pkg":        c.PkgName,
		"author":     c.Author,
		"time":       time.Now().Format("2006-01-02 15:04:05"),
		"actID":      c.Id,
		"structName": c.StructName,
		"prefix":     prefix,
		"import":     importStr,
	})
	if err != nil {
		return "", fmt.Errorf("[GenModel] Execute  Error:%s", err.Error())
	}
	return output.String(), nil
}

func GenTask(c *config.Config) (string, error) {
	if c == nil {
		return "", errors.New("config is nil")
	}
	output, err := parser.With("Task").Parse(activity.Task).GoFmt(true).Execute(map[string]interface{}{
		"pkg":   c.PkgName,
		"actID": c.Id,
	})
	if err != nil {
		return "", fmt.Errorf("[GenTask] Execute  Error:%s", err.Error())
	}
	return output.String(), nil
}

func GenConfig(c *config.Config) (string, error) {
	if c == nil {
		return "", errors.New("config is nil")
	}
	output, err := parser.With("Cofing").Parse(activity.Cofing).GoFmt(true).Execute(map[string]interface{}{
		"pkg":   c.PkgName,
		"actID": c.Id,
	})
	if err != nil {
		return "", fmt.Errorf("[GenConfig] Execute  Error:%s", err.Error())
	}
	return output.String(), nil
}

func GenConst(c *config.Config) (string, error) {
	if c == nil {
		return "", errors.New("config is nil")
	}
	output, err := parser.With("Const").Parse(activity.Const).GoFmt(true).Execute(map[string]interface{}{
		"pkg": c.PkgName,
	})
	if err != nil {
		return "", fmt.Errorf("[GenConst] Execute  Error:%s", err.Error())
	}
	return output.String(), nil
}

func GenFunc(c *config.Config) (string, error) {
	if c == nil {
		return "", errors.New("config is nil")
	}
	output, err := parser.With("Function").Parse(activity.Function).GoFmt(true).Execute(map[string]interface{}{
		"pkg": c.PkgName,
	})
	if err != nil {
		return "", fmt.Errorf("[GenFunc] Execute  Error:%s", err.Error())
	}
	return output.String(), nil
}

func GenTest(c *config.Config) (string, error) {
	if c == nil {
		return "", errors.New("config is nil")
	}
	output, err := parser.With("Test").Parse(activity.Test).GoFmt(true).Execute(map[string]interface{}{
		"pkg": c.PkgName,
	})
	if err != nil {
		return "", fmt.Errorf("[GenTest] Execute  Error:%s", err.Error())
	}
	return output.String(), nil
}

func GenAction(c *config.Config, info *config.ProtoData) (string, error) {
	if c == nil || info == nil {
		return "", errors.New("config is nil")
	}
	var funcAction string
	for idx, req := range info.Req {
		output, err := parser.With("FuncAction").Parse(activity.FuncAction).GoFmt(true).Execute(map[string]interface{}{
			"funcName":  req[:len(req)-3],
			"actID":     c.Id,
			"protoReq":  req,
			"protoResp": info.Resp[idx],
		})
		if err != nil {
			return "", fmt.Errorf("[FuncAction] Execute  Error:%s", err.Error())
		}
		funcAction += output.String() + "\n\n"
	}

	output, err := parser.With("Action").Parse(activity.Action).GoFmt(true).Execute(map[string]interface{}{
		"pkg":        static.ActivityPrefix + c.Id,
		"author":     c.Author,
		"time":       time.Now().Format("2006-01-02 15:04:05"),
		"actID":      c.Id,
		"structName": c.PkgName + "." + (c.StructName),
		"funcAction": funcAction,
	})
	if err != nil {
		return "", fmt.Errorf("[GenAction] Execute  Error:%s", err.Error())
	}
	return output.String(), nil
}

func GenActionFunc(c *config.Config) (string, error) {
	if c == nil {
		return "", errors.New("config is nil")
	}
	output, err := parser.With("ActFunc").Parse(activity.ActFunc).GoFmt(true).Execute(map[string]interface{}{
		"pkg":        static.ActivityPrefix + c.Id,
		"author":     c.Author,
		"time":       time.Now().Format("2006-01-02 15:04:05"),
		"actID":      c.Id,
		"structName": c.PkgName + "." + (c.StructName),
		"pkgName":    c.PkgName,
	})
	if err != nil {
		return "", fmt.Errorf("[GenAction] Execute  Error:%s", err.Error())
	}
	return output.String(), nil
}

func GenActId(c *config.Config) error {
	var (
		err          error
		addActId     = "\n" + "Act" + c.Id + " " + "=" + " " + c.Id + " " + "//" + c.ActivityChineseName
		actnote      = "// 活动ID(此段注释不要修改，正则表达式需要)"
		lineNum, num int64
	)
	// 1. 读取文件内容
	filename := filepath.Join(c.OutputPath, static.StaticPath)
	// 读取文件内容
	f, err := os.OpenFile(filename, os.O_RDWR, 0644)
	if err != nil {
		return err
	}
	var (
		content     string   //读取的文件内容
		contentList []string //读取的文件内容list,是一个切片，代表每一行的内容
	)
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		num++
		if strings.TrimSpace(scanner.Text()) == actnote {
			lineNum = num
		}
		content += scanner.Text() + "\n"
		contentList = append(contentList, scanner.Text())
	}

	if err = scanner.Err(); err != nil {
		return err
	}

	if lineNum <= 0 || lineNum > int64(len(contentList)) {
		return fmt.Errorf("GenActId line number out of range")
	}

	// 查找活动id常量类的定义
	reg := regexp.MustCompile(`Act\d+ = (\d+) \/\/(.*)`)
	lineNum += int64(len(reg.FindAllString(content, -1)) + 1) //查看活动id在第几行

	contentList[lineNum-1] = strings.Join([]string{contentList[lineNum-1], addActId}, "") //添加内容
	newContent := []byte(strings.Join(contentList, "\n"))                                 //将切片转换成字符串
	output, err := parser.With("newContent").Parse(string(newContent)).GoFmt(true).Execute(nil)
	if err != nil {
		return err
	}
	// 将文件指针移到插入位置并写入内容
	_, err = f.WriteAt(output.Bytes(), 0)
	if err != nil {
		return err
	}
	return nil
}

func GenRouter(c *config.Config, info *config.ProtoData) error {
	var (
		err                          error
		lineNum, num                 int64
		actnote                      = "// RegisterBusiness 注册业务接口(此段注释不要修改，正则表达式需要)"
		addcontent, addcontentAction = "\n" + "//" + c.Id + "\n", "\n\n"
	)
	// 1. 读取文件内容
	filename := filepath.Join(c.OutputPath, static.RouterPath)
	// 读取文件内容
	f, err := os.OpenFile(filename, os.O_RDWR, 0644)
	if err != nil {
		return err
	}

	var (
		content     string   //读取的文件内容
		contentList []string //读取的文件内容list,是一个切片，代表每一行的内容
	)
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		num++
		if strings.TrimSpace(scanner.Text()) == actnote {
			lineNum = num
		}
		content += scanner.Text() + "\n"
		contentList = append(contentList, scanner.Text())
	}

	if err = scanner.Err(); err != nil {
		return err
	}

	if lineNum <= 0 || lineNum > int64(len(contentList)) {
		return fmt.Errorf("FuncAction line number out of range")
	}

	re := regexp.MustCompile(`(?s)func RegisterBusiness\(p tcpframework.Processor\) \{(.*?)\}`)
	// 查找所有匹配项
	matches := re.FindAllStringSubmatch(content, -1)
	var lines []string
	for _, match := range matches {
		lines = strings.Split(match[1], "\n")
	}
	lineNum += int64(len(lines) - 1) //查看活动id在第几行

	//生成内容
	for _, req := range info.Req {
		output, err := parser.With("RegisterBusiness").Parse(activity.RegisterBusiness).GoFmt(true).Execute(map[string]interface{}{
			"funcName": req[:len(req)-3],
			"actID":    c.Id,
			"protoReq": req,
		})
		if err != nil {
			return fmt.Errorf("[FuncAction] Execute  Error:%s", err.Error())
		}
		addcontent += output.String()
	}

	contentList[lineNum-1] = strings.Join([]string{contentList[lineNum-1], addcontent}, "") //添加内容

	//生成内容
	for idx, req := range info.Req {
		output, err := parser.With("ConstAction").Parse(activity.ConstAction).GoFmt(true).Execute(map[string]interface{}{
			"funcName":  req[:len(req)-3],
			"actID":     c.Id,
			"protoReq":  req,
			"protoResp": info.Resp[idx],
		})
		if err != nil {
			return fmt.Errorf("[FuncAction] Execute  Error:%s", err.Error())
		}
		addcontentAction += output.String()
	}
	contentList[num-1] = strings.Join([]string{contentList[num-1], addcontentAction}, "") //添加内容
	newContent := []byte(strings.Join(contentList, "\n"))                                 //将切片转换成字符串
	output, err := parser.With("newContent1").Parse(string(newContent)).GoFmt(true).Execute(nil)
	if err != nil {
		return err
	}
	//将文件指针移到插入位置并写入内容
	_, err = f.WriteAt(output.Bytes(), 0)
	if err != nil {
		return err
	}
	return nil
}

func GenRegisterFactory(c *config.Config) error {
	var (
		lineNum, num int
		actnote      = "case static.Act7001: // 免费金币(此段注释不要修改，正则表达式需要)"
		content      string   //读取的文件内容
		contentList  []string //读取的文件内容list,是一个切片，代表每一行的内容
	)
	filename := filepath.Join(c.OutputPath, static.RegisterFactory)
	f, err := os.OpenFile(filename, os.O_RDWR, 0644)
	if err != nil {
		return err
	}
	scanner := bufio.NewScanner(f) // 增加buffer大小
	for scanner.Scan() {
		num++
		if strings.TrimSpace(scanner.Text()) == actnote {
			lineNum = num
		}
		content += scanner.Text() + "\n"
		contentList = append(contentList, scanner.Text())
	}

	if err = scanner.Err(); err != nil {
		return err
	}

	if lineNum <= 0 || lineNum > (len(contentList)) {
		return fmt.Errorf("GenRegisterFactory line number out of range")
	}

	/// 定义正则表达式
	pattern := regexp.MustCompile(`case\s+static\.Act\d+:(?:.|\n)*?\}\s*\n`)

	// 查找匹配结果
	matches := pattern.FindAllStringSubmatch(content, -1)

	var lines []string
	for _, match := range matches {
		lines = strings.Split(match[0], "\n")
	}
	count := len(lines) - 7 //定位到倒数第二个case分支
	lineNum += count
	contentList[lineNum-1] = strings.Join([]string{contentList[lineNum-1], activity.RegisterFactory}, "") //添加内容
	newContent := []byte(strings.Join(contentList, "\n"))                                                 //将切片转换成字符串
	output, err := parser.With("newContent2").Parse(string(newContent)).GoFmt(true).Execute(map[string]interface{}{
		"actID":               c.Id,
		"pkg":                 c.PkgName,
		"activityChineseName": c.ActivityChineseName,
	})
	if err != nil {
		return err
	}
	// 将文件指针移到插入位置并写入内容
	_, err = f.WriteAt(output.Bytes(), 0)
	if err != nil {
		return err
	}
	return nil
}

func GenRegisterActConfig(c *config.Config) error {
	var (
		lineNum, num int
		actnote      = "// 此段注释不要修改，正则表达式需要"
		content      string   //读取的文件内容
		contentList  []string //读取的文件内容list,是一个切片，代表每一行的内容
	)
	filename := filepath.Join(c.OutputPath, static.Init)
	f, err := os.OpenFile(filename, os.O_RDWR, 0644)
	if err != nil {
		return err
	}
	scanner := bufio.NewScanner(f) // 增加buffer大小
	for scanner.Scan() {
		num++
		if strings.TrimSpace(scanner.Text()) == actnote {
			lineNum = num
		}
		content += scanner.Text() + "\n"
		contentList = append(contentList, scanner.Text())
	}

	if err = scanner.Err(); err != nil {
		return err
	}

	if lineNum <= 0 || lineNum > (len(contentList)) {
		return fmt.Errorf("GenRegisterFactory line number out of range")
	}
	lineNum += 3
	addcontent := fmt.Sprintf("\nnew(%s.Config),// %s\t", c.PkgName, c.ActivityChineseName)
	contentList[lineNum-1] = strings.Join([]string{contentList[lineNum-1], addcontent}, "") //添加内容
	newContent := []byte(strings.Join(contentList, "\n"))                                   //将切片转换成字符串

	output, err := parser.With("newContent3").Parse(string(newContent)).GoFmt(true).Execute(nil)
	if err != nil {
		return err
	}
	// 将文件指针移到插入位置并写入内容
	_, err = f.WriteAt(output.Bytes(), 0)
	if err != nil {
		return err
	}
	return nil
}
