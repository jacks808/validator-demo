package tivalidator

import (
	_ "github.com/go-playground/locales/en"
	_ "github.com/go-playground/universal-translator"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	_ "github.com/go-playground/validator/v10/translations/en"
	"reflect"
)

var (
	uni      *ut.UniversalTranslator
	validate *validator.Validate
	trans    ut.Translator
	fn       validator.Func
)

type ValidationError struct {

	// Tag returns the validation tag that failed. if the
	// validation was an alias, this will return the
	// alias name and not the underlying tag that failed.
	//
	// eg. alias "iscolor": "hexcolor|rgb|rgba|hsl|hsla"
	// will return "iscolor"
	Tag string

	// ActualTag returns the validation tag that failed, even if an
	// alias the actual tag within the alias will be returned.
	// If an 'or' validation fails the entire or will be returned.
	//
	// eg. alias "iscolor": "hexcolor|rgb|rgba|hsl|hsla"
	// will return "hexcolor|rgb|rgba|hsl|hsla"
	ActualTag string

	// Namespace returns the namespace for the field error, with the tag
	// name taking precedence over the field's actual name.
	//
	// eg. JSON name "User.fname"
	//
	// See StructNamespace() for a version that returns actual names.
	//
	// NOTE: this field can be blank when validating a single primitive field
	// using validate.Field(...) as there is no way to extract it's name
	Namespace string

	// StructNamespace returns the namespace for the field error, with the field's
	// actual name.
	//
	// eq. "User.FirstName" see Namespace for comparison
	//
	// NOTE: this field can be blank when validating a single primitive field
	// using validate.Field(...) as there is no way to extract its name
	StructNamespace string

	// Field returns the fields name with the tag name taking precedence over the
	// field's actual name.
	//
	// eq. JSON name "fname"
	// see StructField for comparison
	Field string

	// StructField returns the field's actual name from the struct, when able to determine.
	//
	// eq.  "FirstName"
	// see Field for comparison
	StructField string

	// Value returns the actual field's value in case needed for creating the error
	// message
	Value interface{}

	// Param returns the param value, in string form for comparison; this will also
	// help with generating an error message
	Param string

	// Kind returns the Field's reflect Kind
	//
	// eg. time.Time's kind is a struct
	Kind reflect.Kind

	// Type returns the Field's reflect Type
	//
	// eg. time.Time's type is time.Time
	Type reflect.Type

	// Error returns the FieldError's message
	Error string

	RealError error
}

func Init() {
	validate = validator.New()

	validate.RegisterValidation("public", Public, true)
	validate.RegisterValidation("port", Port, true)
	validate.RegisterValidation("version", Version, true)

	// init translator
	//enTranslator := en.New()
	//uni = ut.New(enTranslator, enTranslator)
	//trans, _ := uni.GetTranslator("en")
	//
	//en_translations.RegisterDefaultTranslations(validate, trans)
}

// 注册函数适配器，自定义，并加校验
//func RegisterTag(tag string, ) {
//
//	validate.RegisterValidation(tag, fn, true)
//}

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
