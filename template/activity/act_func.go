// Package activity
// @description
// @author      张盛钢
// @datetime    2023/4/10 16:22
package activity

const ActFunc = `
package {{.pkg}}

import (
	"casino/module/business/service/activity/base"
	"casino/module/business/service/activity/base/static"
	"casino/module/business/service/activity/manager"
	"casino/module/business/service/activity/{{.pkgName}}"
	"github.com/aaagame/common/log"
	"github.com/aaagame/common/middleware/trace"
	"runtime/debug"
)

// GetObject 获取对象
func GetObject(ctx *trace.Context, container *manager.Container, delivery *base.Delivery) (oneself *{{.structName}}, ok bool) {
	// 获取、识别对象
	template, err := container.Get(ctx, static.Act{{.actID}}, delivery)
	if base.ScanActErr(err, static.Act{{.actID}}) {
		return
	}
	if !template.Verify(ctx) { // 活动没开启
		return
	}
	if oneself, ok = template.(*{{.structName}}); !ok { // 断言失败
		log.Errorf(static.Act{{.actID}}, "[GetObject]template.(*{{.structName}}) assert failed %s", debug.Stack())
		return
	}
	return
}

`
