// Package activity
// @description
// @author      张盛钢
// @datetime    2023/4/26 10:39
package activity

const RegisterFactory = `
case static.Act{{.actID}}: // {{.activityChineseName}}
			c.Register(actId, {{.pkg}}.New())
			c.RegisterMsg(actId)`
