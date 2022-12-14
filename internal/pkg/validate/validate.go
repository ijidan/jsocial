package validate

import (
	mapset "github.com/deckarep/golang-set"
	"github.com/go-playground/validator/v10"
	"github.com/ijidan/jsocial/internal/constant"
	"regexp"
)

func MobileValidator(f validator.FieldLevel) bool {
	value := f.Field().String()
	result, _ := regexp.MatchString(`^(1[3|4|5|8][0-9]\d{4,8})$`, value)
	if result {
		return true
	} else {
		return false
	}

}

func ImageValidator(f validator.FieldLevel) bool {
	value := f.Field().String()
	imageFormatSlice := []interface{}{constant.ImageFormatOfJpg, constant.ImageFormatOfJpeg, constant.ImageFormatOfGif, constant.ImageFormatOfPng, constant.ImageFormatOfBmp}
	imageFormatSet := mapset.NewSetFromSlice(imageFormatSlice)
	return imageFormatSet.Contains(value)

}
