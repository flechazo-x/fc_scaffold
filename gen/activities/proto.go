package activities

import (
	"errors"
	"fc_scaffold/config"
	"regexp"
)

func GetProtoInfo(content string) *config.ProtoData {
	var protoData = new(config.ProtoData)

	// 创建正则表达式并匹配文本
	pattern := regexp.MustCompile(`message\s+(\w+)(Req|Resp)\s*\{`)

	// 提取匹配到的数据
	match := pattern.FindAllStringSubmatch(content, -1)
	for _, m := range match {
		switch m[2] {
		case "Req":
			protoData.Req = append(protoData.Req, m[1]+"Req")
		case "Resp":
			protoData.Resp = append(protoData.Resp, m[1]+"Resp")
		}
	}
	if len(protoData.Resp) != len(protoData.Req) {
		config.PrintErr(errors.New("proto req length != resp length"))
	}
	return protoData
}
