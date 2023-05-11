package static

import "fmt"

//活动前缀

const (
	ActivityPrefix = "A" //
)

// 活动文件名
const (
	ModelFile  = "Model.go"
	ConfigFile = "Config.go"
	ConstFile  = "Const.go"
	FuncFile   = "Function.go"
	Action     = "Action.go"
	Task       = "Task.go"
)

var (
	TestFile = func(s string) string {
		return fmt.Sprintf("%s_test.go", s)
	}
)

// 活动路径
const (
	ActivityPath    = "module/business/service/activity"
	ActionPath      = "module/business/service/action/activities"
	StaticPath      = "module/business/service/activity/base/static/Static.go"
	RouterPath      = "module/business/controller/router.go"
	RegisterFactory = "module/business/service/activity/manager/Manager.go"
	Init            = "module/business/service/initapp/Init.go"
	SqlGoPath       = "/module/business/dao/activities"
)

func GetProtoPath(id string) string {
	return fmt.Sprintf("proto/act%s/act%s.proto", id, id)
}

func GetProtoJsonPath(id string) string {
	return fmt.Sprintf("proto/act%s/act%s.json", id, id)
}

var (
	Welcome = "Welcome to [fc_scaffold]!\n" +
		"The ultimate scaffold tool for developing your projects quickly and easily.\n" +
		"Our tool provides advanced features and comfortable templates to get you started in development in no time.\n" +
		"Let's build something great together!\n"
	ProtoHint = "----------Please make sure you have generated the proto file and executed the <!!! BuildProto2.bat !!!> script before using this tool.~~~ ----------\n"
	Options   = "----------Input your option: 0 for parameter mode, 1 for interactive mode,2 for http mode.----------\n"
)
