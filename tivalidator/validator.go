package tivalidator

import (
	"fmt"
	_ "github.com/go-playground/locales/en"
	_ "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	_ "github.com/go-playground/validator/v10/translations/en"
	log "github.com/sirupsen/logrus"
	"strings"
	"time"
)

//var (
//	uni      *ut.UniversalTranslator
//	validate *validator.Validate
//	trans    ut.Translator
//)

func NewValidator() *validator.Validate {
	validate := validator.New()

	validate.RegisterValidation("public", Public, true)
	validate.RegisterValidation("port", Port, true)

	// init translator
	//enTranslator := en.New()
	//uni = ut.New(enTranslator, enTranslator)
	//trans, _ := uni.GetTranslator("en")
	//
	//en_translations.RegisterDefaultTranslations(validate, trans)
	return validate
}

func RegisterTag(validate *validator.Validate, tag string, fn validator.Func) {
	validate.RegisterValidation(tag, fn, true)
}

func ValidateStruct(validate *validator.Validate, s interface{}) string {
	now := time.Now().UnixNano()
	err := validate.Struct(s)
	errorMessage := buildErrorMessage(err)
	// 性能测试
	log.Infof("validator parse struct cost time %fms", float64(time.Now().UnixNano()-now)/1e6)
	return errorMessage
}

// build error message
func buildErrorMessage(err error) string {
	if err != nil {
		if _, ok := err.(*validator.InvalidValidationError); ok {
			return fmt.Sprintf("%v", err)
		}

		count := 0
		sb := strings.Builder{}
		for _, err := range err.(validator.ValidationErrors) {
			count++
			//fmt.Printf("Namespace: %v\n", err.Namespace())
			//fmt.Printf("Field: %v\n", err.Field())
			//fmt.Printf("StructNamespace: %v\n", err.StructNamespace())
			//fmt.Printf("StructField: %v\n", err.StructField())
			//fmt.Printf("Tag: %v\n", err.Tag())
			//fmt.Printf("ActualTag: %v\n", err.ActualTag())
			//fmt.Printf("Kind: %v\n", err.Kind())
			//fmt.Printf("Type: %v\n", err.Type())
			//fmt.Printf("Value: %v\n", err.Value())
			//fmt.Printf("Param: %v\n", err.Param())

			// 翻译标准错误结果
			// fmt.Println(err.Translate(trans))

			// 构建标准Error
			sprintf := fmt.Sprintf("InvalidParameter Error %d, field:'%v', current value: '%v', field require: '%v %v'\n",
				count, err.Field(), err.Value(), err.Tag(), err.Param())
			sb.Write([]byte(sprintf))
		}

		// from here you can create your own error messages in whatever language you wish
		return sb.String()
	}
	return ""
}
