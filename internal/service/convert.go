package service

import (
	"github.com/ijidan/jsocial/api/proto_build"
	"github.com/ijidan/jsocial/internal/constant"
)

var TypeMap = map[uint64]proto_build.IMAGE_TYPE{
	constant.ImageTypeOrigin:    proto_build.IMAGE_TYPE_ORIGIN,
	constant.ImageTypeBig:       proto_build.IMAGE_TYPE_BIG,
	constant.ImageTypeThumbnail: proto_build.IMAGE_TYPE_THUMBNAIL,
}
var FormatMap = map[uint64]proto_build.IMAGE_FORMAT{
	constant.ImageFormatJpg:   proto_build.IMAGE_FORMAT_JPG,
	constant.ImageFormatGif:   proto_build.IMAGE_FORMAT_GIF,
	constant.ImageFormatPng:   proto_build.IMAGE_FORMAT_PNG,
	constant.ImageFormatBmp:   proto_build.IMAGE_FORMAT_BMP,
	constant.ImageFormatOther: proto_build.IMAGE_FORMAT_OTHER,
}
