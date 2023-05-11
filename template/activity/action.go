// Package activity 生成活动协议
package activity

const Action = `
// Package {{.pkg}}
// @description
// @author      {{.author}}
// @datetime    {{.time}}
package {{.pkg}}

import (
	"casino/module/business/service/action"
	"casino/module/network/task"
	"casino/module/protocol/act{{.actID}}"

	"casino/module/protocol/cmd"
	"github.com/aaagame/common/middleware/trace"
)

{{.funcAction}}

`
const FuncAction = `func {{.funcName}}Action(ctx *trace.Context, u *task.User, _ *act{{.actID}}.{{.protoReq}}) (resp *act{{.actID}}.{{.protoResp}}) {
	//todo (to be coded)
	resp = new(act{{.actID}}.{{.protoResp}})
	resp.Err = cmd.ErrorCode_ActivityClose

	delivery := action.DeliveryArgs(u) //获取对象
	defer action.Retrieve(delivery)    //回收对象

	_, ok := GetObject(ctx, u.Activity, delivery) //_ ==> oneself
	if !ok {
		return resp
	}
	
	return resp
}`
