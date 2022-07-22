package initialize

import (
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	enTranslations "github.com/go-playground/validator/v10/translations/en"
	zhTranslations "github.com/go-playground/validator/v10/translations/zh"
	"github.com/hail-pas/GinStartKit/global"
	"reflect"
	"strings"
)

var (
	uni *ut.UniversalTranslator
)

func ValidateWithTranslation(locale string) {
	//注册翻译器
	zhTranslation := zh.New()
	uni = ut.New(zhTranslation, zhTranslation)

	global.Translator, _ = uni.GetTranslator("zhTranslation")

	//获取gin的校验器
	global.Validate = binding.Validator.Engine().(*validator.Validate)
	// 注册一个获取json的字段名的自定义方法
	global.Validate.RegisterTagNameFunc(func(field reflect.StructField) string {
		name := strings.SplitN(field.Tag.Get("label"), ",", 2)[0]
		if name == "-" {
			return ""
		}
		if name == "" {
			name := strings.SplitN(field.Tag.Get("json"), ",", 2)[0]
			if name == "-" {
				return ""
			}
			return name
		}
		return name
	})
	//注册翻译器
	switch locale {
	case "en":
		_ = enTranslations.RegisterDefaultTranslations(global.Validate, global.Translator)
	case "zh":
		_ = zhTranslations.RegisterDefaultTranslations(global.Validate, global.Translator)
	default:
		_ = enTranslations.RegisterDefaultTranslations(global.Validate, global.Translator)
	}
	return
}
