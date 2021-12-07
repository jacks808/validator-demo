package main

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"strings"
	"validator-demo/proto"
	"validator-demo/tivalidator"
)

// use a single instance of Validate, it caches struct info
var validate *validator.Validate

func main() {
	validate = validator.New()

	validate.RegisterValidation("rocket", tivalidator.RocketValidator, true)
	validateStruct()
}

func validateStruct() {

	user := &proto.CreateUserRequest{
		FirstName:      "Badger",
		LastName:       "Smith",
		Age:            135,
		Email:          "Badger.Smith@gmail.com",
		FavouriteColor: "#000-",
	}

	// returns nil or ValidationErrors ( []FieldError )
	err := validate.Struct(user)
	if err != nil {
		msg := buildErrorMessage(err)
		fmt.Println(msg)
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

		sb.WriteString("InvalidParameterError\n")

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
			//fmt.Println()

			// 构建标准Error
			msg := fmt.Sprintf("error %d, field: '%v', current value: '%v', field require: '%v %v'\n",
				count, err.Field(), err.Value(), err.Tag(), err.Param())
			sb.WriteString(msg)
		}

		// from here you can create your own error messages in whatever language you wish
		return sb.String()
	}
	return ""
}