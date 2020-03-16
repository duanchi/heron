package validate

import (
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator"
	zh_translations "github.com/go-playground/validator/translations/zh"
)

var trans ut.Translator

func Init () {
	binding.Validator = new(defaultValidator)
	zh := zh.New()
	en := en.New()
	uni := ut.New(en, zh)
	trans, _ = uni.GetTranslator("zh")

	zh_translations.RegisterDefaultTranslations(binding.Validator.Engine().(*validator.Validate), trans)

}
