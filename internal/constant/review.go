package constant

import "github.com/ijidan/jsocial/api/proto_build"

const (
	ReviewInit = 0 //未提交审核
	ReviewTo   = 1 //已提交，待审核
	ReviewPass = 2 //审核通过
	ReviewFail = 3 //审核未通过
)

var ReviewStatusDbToProto = map[int32]proto_build.ReviewStatus{
	ReviewInit: proto_build.ReviewStatus_ReviewInit,
	ReviewTo:   proto_build.ReviewStatus_ReviewTo,
	ReviewPass: proto_build.ReviewStatus_ReviewPass,
	ReviewFail: proto_build.ReviewStatus_ReviewFail,
}
var ReviewStatusProtoToDb = map[proto_build.ReviewStatus]int32{
	proto_build.ReviewStatus_ReviewInit: ReviewInit,
	proto_build.ReviewStatus_ReviewTo:   ReviewTo,
	proto_build.ReviewStatus_ReviewPass: ReviewPass,
	proto_build.ReviewStatus_ReviewFail: ReviewFail,
}
