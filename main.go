package main

import (
	"fmt"
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
	log "github.com/sirupsen/logrus"
	"strings"
	"time"
	pb "validator-demo/proto"
	"validator-demo/tivalidator"
)

var (
	uni      *ut.UniversalTranslator
	validate *validator.Validate
	trans    ut.Translator
)

func main() {
	validate = ValidatorInit()
	validateStruct()
}

func ValidatorInit() *validator.Validate {
	validate := validator.New()
	validate.RegisterValidation("rocket", tivalidator.RocketValidator, true)
	validate.RegisterStructValidation(UserStructLevelValidation, true)

	validate.RegisterValidation("public", tivalidator.Public, true)
	validate.RegisterValidation("version", tivalidator.Version, true)

	// init translator
	enTranslator := en.New()
	uni = ut.New(enTranslator, enTranslator)
	trans, _ := uni.GetTranslator("en")

	en_translations.RegisterDefaultTranslations(validate, trans)
	return validate
}

func validateStruct() {

	user := pb.CreateUserRequest{
		FirstName:      "Badger",
		LastName:       "Smith",
		Age:            135,
		Email:          "Badger.Smith@gmail.com",
		FavouriteColor: "#000-",
		// format : 2006-01-02 15:04:05
		CreateTime: time.Now().Format("2006-01-02 15:04:05"),
		Version:    "1.0.2.",
	}

	// returns nil or ValidationErrors ( []FieldError )
	err := validate.Struct(user)
	if err != nil {
		log.Errorf(buildErrorMessage(err))
		return
	}

	// do sth such as save user to database...
}

// build error message
func buildErrorMessage(err error) string {
	count := 0

	sb := strings.Builder{}

	if err != nil {

		// this check is only needed when your code could produce
		// an invalid value for validation such as interface with nil
		// value most including myself do not usually have code like this.
		if _, ok := err.(*validator.InvalidValidationError); ok {
			return fmt.Sprintf("%v", err)
		}

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
			fmt.Println(err.Translate(trans))

			// 构建标准Error
			log.Errorf("InvalidParameter Error %d, field:'%v', current value: '%v', field require: '%v %v'\n",
				count, err.Field(), err.Value(), err.Tag(), err.Param())
		}

		// from here you can create your own error messages in whatever language you wish
		return sb.String()
	}
	return ""
}

func UserStructLevelValidation(sl validator.StructLevel) {

	user := sl.Current().Interface().(pb.CreateUserRequest)

	if len(user.FirstName) == 0 && len(user.LastName) == 0 {
		sl.ReportError(user.FirstName, "fname", "FirstName", "fnameorlname", "")
		sl.ReportError(user.LastName, "lname", "LastName", "fnameorlname", "")
	}
}
