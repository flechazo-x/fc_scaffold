// Package activity
// @description
// @author      张盛钢
// @datetime    2023/4/10 17:12
package activity

const Task = `
package {{.pkg}}

import (
	"casino/module/business/service/activity/base"
	"casino/module/business/service/activity/base/static"
	"github.com/aaagame/common/log"
)

type Task struct {
	*base.TaskEntity
}

func NewTask(missionType static.Event, value, arg uint64, missionID int32, state static.State) *Task {
	t := new(Task)
	t.TaskEntity = base.NewTaskEntity(missionID, missionType, value, arg, state)
	return t
}

// Notify 任务单元执行逻辑,增加进度
func (t *Task) Notify(msg interface{}) {
	//todo 需要将配置参数修改,然后将注释删除即可
	//if payLoad, ok := msg.(*base.PayLoad); ok {
	//	if !t.IsMetaExecuteOK() {
	//		return
	//	}
	//todo 判断是否指定机器id 自定义逻辑，如果有需要自己手动增加
	//
	//	config := GetTaskConfigByIDAndType(t.MissionPoolID, t.MissionType)
	//	if config == nil {
	//		return
	//	}
	//	//进度增加
	//	run := base.GetTaskFunc(static.Act{{.actID}}, t.MissionType)
	//	if run == nil {
			log.Errorf(static.Act{{.actID}}, "[{{.pkg}}.notify]unregister task func event=%d", t.MissionType)
	//		return
	//	}
	//	payLoad.ActId, payLoad.Metric = static.Act{{.actID}}, &base.Metric{Target: config.MissionValue, Limit: config.MissionArg} // 参数补充
	//	run(t.TaskEntity, payLoad, t)
	//	//状态判断
	//	t.IsComplete(payLoad, nil)
	//}
}
`
