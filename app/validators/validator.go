package validators

import (
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	enTranslations "github.com/go-playground/validator/v10/translations/en"
	zhTranslations "github.com/go-playground/validator/v10/translations/zh"
	"go_code/gintest/bootstrap/glog"
	"reflect"
)

var trans ut.Translator

func InitTrans(locale string) {
	var err error
	zhT := zh.New() // 中文翻译器
	enT := en.New() // 英文翻译器

	// 第一个参数是备用（fallback）的语言环境
	// 后面的参数是应该支持的语言环境（支持多个）
	// uni := translator.New(zhT, zhT) 也是可以的
	uni := ut.New(enT, zhT, enT)

	// locale 通常取决于 http 请求头的 'Accept-Language'
	// 也可以使用 uni.FindTranslator(...) 传入多个 locale 进行查找
	trans, _ = uni.GetTranslator(locale)

	// 修改 gin 框架中的 Validator 引擎属性，实现自定义
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		// 重点1：注册一个函数，获取 struct tag 里自定义的 label 作为字段名
		v.RegisterTagNameFunc(func(fld reflect.StructField) string {
			label := fld.Tag.Get("label")
			if label == "" {
				return fld.Name
			}
			return label
		})

		// 注册翻译器
		switch locale {
		case "en":
			err = enTranslations.RegisterDefaultTranslations(v, trans)
		case "zh":
			err = zhTranslations.RegisterDefaultTranslations(v, trans)
		default:
			err = enTranslations.RegisterDefaultTranslations(v, trans)
		}
		if err != nil {
			glog.SL.Error("初始化翻译器错误", err)
		}
	}
}

// Translate 翻译工具
func translate(error error, s interface{}) map[string]string {
	r := make(map[string]string)
	t := reflect.TypeOf(s).Elem()
	for _, err := range error.(validator.ValidationErrors) {
		// 重点2：使用反射方法获取 struct 种的 json 标签作为 key
		var k string
		if field, ok := t.FieldByName(err.StructField()); ok {
			k = field.Tag.Get("json")
		}
		if k == "" {
			k = err.StructField()
		}
		r[k] = err.Translate(trans)
	}
	return r
}

// Error 返回单个验证错误
func Error(error error, request interface{}) string {
	for _, v := range translate(error, request) {
		return v
	}
	return ""
}
