// Package activity 生成活动结构
package activity

const Model = `
// Package {{.pkg}}
// @description 
// @author      {{.author}}
// @datetime    {{.time}}
package {{.pkg}}

{{.import}}

type {{.structName}} struct {
	base.TemplateInstance                //活动模板
}

func New() *{{.structName}} {
	{{.prefix}} := new({{.structName}})
	{{.prefix}}.BaseUnlockBet = bet.NewBaseUnlockBet() //todo 默认生成解锁bet,如果活动没有解锁bet需要将这行删除
	return {{.prefix}}
}

func ({{.prefix}} *{{.structName}}) Id() int32 {
	return static.Act{{.actID}}
}

func ({{.prefix}} *{{.structName}}) Verify(_ *trace.Context) bool {
	return activity.Verify({{.prefix}}.Id())
}

func ({{.prefix}} *{{.structName}}) Init(ctx *trace.Context, delivery *base.Delivery) (error, activity.State) {
	//todo (to be coded)
	return nil, activity.StateCodeOpen
}

func ({{.prefix}} *{{.structName}}) Durable(ctx *trace.Context, delivery *base.Delivery) error {
	//todo (to be coded)
	return nil
}
`

const Support = `
func ({{.prefix}} *{{.structName}}) PaySupport() bool {
	return true
}

// PayNotifyCallback 支付操作
func ({{.prefix}} *{{.structName}}) PayNotifyCallback(ctx *trace.Context, o *bases.Options) error {
	if o.ActId == static.Act{{.actID}} && o.Entrance == "" {  
	//todo (to be coded)
	}
	return nil
}

func ({{.prefix}} *{{.structName}}) PayAfter(ctx *trace.Context, o *bases.Options) error {
		settleOpt := &charge.SettlementOptions{
		AwardOperate:   charge.Operate(o.AwardOperate),
		MergeAwardFunc: exchange.MergeJsonToJson,
		ActId:          o.ActId,
		CallbackId:     o.CallbackId,
		Award:          o.InputProps,
	}
	if o.Delivery != nil {
		settleOpt.AfterBalance = o.Delivery.Balance
	}
	return charge.Settlement(ctx, settleOpt)
}
`

const Import = `
import (
	"casino/module/business/service/activity/base"
	"github.com/aaagame/common/middleware/trace"
	"casino/module/business/service/activity"
	"casino/module/business/service/activity/base/bet"
	"casino/module/business/service/activity/base"
	"casino/module/business/service/activity/base/static"
)
`
const Import1 = `
import (
	"casino/module/business/dao/charge"
	"casino/module/business/service/activity"
	"casino/module/business/service/activity/base"
	"casino/module/business/service/activity/base/bet"
	"casino/module/business/service/activity/base/static"
	bases "casino/module/business/service/base"
	"github.com/aaagame/common/middleware/trace"
	"casino/module/business/service/prop/exchange"
)
`
