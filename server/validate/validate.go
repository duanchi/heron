package validate

import (
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	zh_translations "github.com/go-playground/validator/translations/zh"
	"gopkg.in/go-playground/validator.v9"
)

var trans ut.Translator

var Validators map[string]validator.Func = map[string]validator.Func{}

func Init () {
	binding.Validator = new(defaultValidator)
	zh := zh.New()
	en := en.New()
	uni := ut.New(en, zh)
	trans, _ = uni.GetTranslator("zh")

	zh_translations.RegisterDefaultTranslations(binding.Validator.Engine().(*validator.Validate), trans)

	for tag, fn := range Validators {
		binding.Validator.Engine().(*validator.Validate).RegisterValidation(tag, fn)
	}

}
