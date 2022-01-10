package tiv

import (
	"errors"
	"github.com/go-playground/validator/v10"
)

var (
	// uni      *ut.UniversalTranslator
	validate *validator.Validate
	// trans    ut.Translator
)

func init() {
	validate = validator.New()

	RegisterValidator("public", Public)
	RegisterValidator("port", Port)
	RegisterValidator("version", Version)
	RegisterAlias("keyMax", "max=10")
	// init translator
	//enTranslator := en.New()
	//uni = ut.New(enTranslator, enTranslator)
	//trans, _ := uni.GetTranslator("en")
	//
	//en_translations.RegisterDefaultTranslations(validate, trans)
}

// RegisterValidator 注册自定义验证函数
func RegisterValidator(tag string, fn func(fl FieldLevel) bool) error {
	if len(tag) == 0 {
		return errors.New("function Key cannot be empty")
	}
	if fn == nil {
		return errors.New("function cannot be empty")
	}
	err := validate.RegisterValidation(tag, convertFieldLevelFunc(fn), true)
	if err != nil {
		return err
	}
	return nil
}

// RegisterAlias 给tag注册别名
// Example: RegisterAlias("keymax", "max=10")
func RegisterAlias(alias string, tag string) {
	validate.RegisterAlias(alias, tag)
}

func convertFieldLevelFunc(f func(fl FieldLevel) bool) validator.Func {
	return func(fl validator.FieldLevel) bool {
		return f(fl)
	}
}

func ValidateStruct(s interface{}) []ValidationError {
	return convert(validate.Struct(s))
}

//func InitTrans(fn ) {
//
//}
//
//func Trans(e ValidationError) string {
//
//}

// build error message
func convert(err error) []ValidationError {
	if err != nil {

		result := make([]ValidationError, 1)

		currError := ValidationError{
			RealError: err,
		}

		if _, ok := err.(*validator.InvalidValidationError); ok {
			result = append(result, currError)
			return result
		}

		for _, err := range err.(validator.ValidationErrors) {
			currError.Tag = err.Tag()
			currError.ActualTag = err.ActualTag()
			currError.Namespace = err.Namespace()
			currError.StructNamespace = err.StructNamespace()
			currError.Field = err.Field()
			currError.StructField = err.StructField()
			currError.Value = err.Value()
			currError.Param = err.Param()
			currError.Kind = err.Kind()
			currError.Type = err.Type()
			currError.Error = err.Error()

			result = append(result, currError)
		}
		return result
	}

	return nil
}
