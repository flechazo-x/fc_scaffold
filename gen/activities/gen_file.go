package activities

import (
	"errors"
	"fc_scaffold/config"
	"fc_scaffold/parser"
	"fc_scaffold/static"
	"os"
	"path/filepath"
)

func CreateFile(c *config.Config) error {
	if c == nil {
		return errors.New("[CreateFile] config==nil")
	}

	var (
		genModel, genConfig, genConst, genFunc, genTest, filename string
		err                                                       error
		protoInfo                                                 []byte
	)
	//读取Proto文件内容
	protoInfo, err = os.ReadFile(filepath.Join(c.ProtoPath, static.GetProtoPath(c.Id)))
	if err != nil {
		return err
	}
	info := GetProtoInfo(string(protoInfo))
	activityFile := filepath.Join(c.OutputPath, filepath.Join(static.ActivityPath, c.PkgName))
	actionFile := filepath.Join(c.OutputPath, filepath.Join(static.ActionPath, static.ActivityPrefix+c.Id))

	err = parser.MkdirIfNotExist(activityFile)
	if err != nil {
		return err
	}
	err = parser.MkdirIfNotExist(actionFile)
	if err != nil {
		return err
	}

	filename = filepath.Join(activityFile, static.ModelFile)
	if !parser.FileExists(filename) {
		genModel, err = GenModel(c)
		if err != nil {
			return err
		}
		err = os.WriteFile(filename, []byte(genModel), os.ModePerm)
		if err != nil {
			return err
		}
	}

	filename = filepath.Join(activityFile, static.ConfigFile)
	if !parser.FileExists(filename) {
		genConfig, err = GenConfig(c)
		if err != nil {
			return err
		}
		err = os.WriteFile(filename, []byte(genConfig), os.ModePerm)
		if err != nil {
			return err
		}
	}
	filename = filepath.Join(activityFile, static.ConstFile)
	if !parser.FileExists(filename) {
		genConst, err = GenConst(c)
		if err != nil {
			return err
		}
		err = os.WriteFile(filename, []byte(genConst), os.ModePerm)
		if err != nil {
			return err
		}
	}
	filename = filepath.Join(activityFile, static.FuncFile)
	if !parser.FileExists(filename) {
		genFunc, err = GenFunc(c)
		if err != nil {
			return err
		}
		err = os.WriteFile(filename, []byte(genFunc), os.ModePerm)
		if err != nil {
			return err
		}
	}
	filename = filepath.Join(activityFile, static.TestFile(c.StructName))
	if !parser.FileExists(filename) {
		genFunc, err = GenTest(c)
		if err != nil {
			return err
		}
		err = os.WriteFile(filename, []byte(genFunc), os.ModePerm)
		if err != nil {
			return err
		}
	}
	if c.IsSupportTasks == 1 {
		filename = filepath.Join(activityFile, static.Task)
		if !parser.FileExists(filename) {
			genFunc, err = GenTask(c)
			if err != nil {
				return err
			}
			err = os.WriteFile(filename, []byte(genFunc), os.ModePerm)
			if err != nil {
				return err
			}
		}
	}
	filename = filepath.Join(actionFile, static.Action)
	if !parser.FileExists(filename) {
		genTest, err = GenAction(c, info)
		if err != nil {
			return err
		}
		err = os.WriteFile(filename, []byte(genTest), os.ModePerm)
		if err != nil {
			return err
		}
	}
	filename = filepath.Join(actionFile, static.FuncFile)
	if !parser.FileExists(filename) {
		genTest, err = GenActionFunc(c)
		if err != nil {
			return err
		}
		err = os.WriteFile(filename, []byte(genTest), os.ModePerm)
		if err != nil {
			return err
		}
	}
	err = GenRouter(c, info) //生成路由
	if err != nil {
		return err
	}

	err = GenActId(c) //生成活动id
	if err != nil {
		return err
	}
	err = GenRegisterFactory(c) //注册工厂生成
	if err != nil {
		return err
	}
	err = GenRegisterActConfig(c) //活动配置
	if err != nil {
		return err
	}
	return nil
}
