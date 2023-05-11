// Package activity
// @description
// @author      张盛钢
// @datetime    2023/4/24 18:41
package activity

const RegisterBusiness = `p.Register((*act{{.actID}}.{{.protoReq}})(nil), Handle{{.funcName}}Req)
`

const ConstAction = `
func Handle{{.protoReq}}(req *act{{.actID}}.{{.protoReq}}, t *task.UserTask) error {
	ctx := trace.Pop().WithIdentifier(&base.Trace{Message: req})
	defer trace.Put(ctx)
	resp := new(act{{.actID}}.{{.protoResp}})
	if ok, code := Verify(t, &Optional{UID: req.UID, Describe: "Handle{{.protoReq}}"}); !ok {
		resp.Err = code
		t.SendMsg(resp)
		return nil
	}
	log.LDebugf("Handle{{.protoReq}} msg=%v", req)

	resp = A{{.actID}}.{{.funcName}}Action(ctx, t.User, req)
	resp.DebugTrace = base.Transport(ctx)
	t.SendMsg(resp)
	return nil
}

`
