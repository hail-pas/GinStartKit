package initialize

import (
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	zh_translations "github.com/go-playground/validator/v10/translations/zh"
	"github.com/hail-pas/GinStartKit/global"
)

var (
	uni *ut.UniversalTranslator
)

func ValidateWithTranslation() {
	//注册翻译器
	zhTranslation := zh.New()
	uni = ut.New(zhTranslation, zhTranslation)

	global.Trans, _ = uni.GetTranslator("zhTranslation")

	//获取gin的校验器
	global.Validate = binding.Validator.Engine().(*validator.Validate)
	//注册翻译器
	err := zh_translations.RegisterDefaultTranslations(global.Validate, global.Trans)
	if err != nil {
		panic(err)
	}
}
