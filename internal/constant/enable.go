package constant

import "github.com/ijidan/jsocial/api/proto_build"

const (
	NotEnable = 0 //未启用
	Enable    = 1 //启用
)

var EnableDbToProto = map[int32]proto_build.IsEnable{
	NotEnable: proto_build.IsEnable_NotEnable,
	Enable:    proto_build.IsEnable_Enable,
}

var EnableProtoToDb = map[proto_build.IsEnable]int32{
	proto_build.IsEnable_NotEnable: NotEnable,
	proto_build.IsEnable_Enable:    Enable,
}
