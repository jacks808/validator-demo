package main

import (
	"database/sql"
	"fmt"
	"github.com/go-playground/validator/v10"
	proto "validator-demo/proto"
	"validator-demo/tivalidator"
)

// DbBackedUser User struct
type DbBackedUser struct {
	Name sql.NullString `validate:"required"`
	Age  sql.NullInt64  `validate:"required"`
}

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

		// this check is only needed when your code could produce
		// an invalid value for validation such as interface with nil
		// value most including myself do not usually have code like this.
		if _, ok := err.(*validator.InvalidValidationError); ok {
			fmt.Println(err)
			return
		}

		for _, err := range err.(validator.ValidationErrors) {

			fmt.Printf("Namespace: %v\n", err.Namespace())
			fmt.Printf("Field: %v\n", err.Field())
			fmt.Printf("StructNamespace: %v\n", err.StructNamespace())
			fmt.Printf("StructField: %v\n", err.StructField())
			fmt.Printf("Tag: %v\n", err.Tag())
			fmt.Printf("ActualTag: %v\n", err.ActualTag())
			fmt.Printf("Kind: %v\n", err.Kind())
			fmt.Printf("Type: %v\n", err.Type())
			fmt.Printf("Value: %v\n", err.Value())
			fmt.Printf("Param: %v\n", err.Param())
			fmt.Println()

			// 构建标准Error
			msg := fmt.Sprintf("=========InvalidParameterError. " +
				"field: '%v', current value: '%v', field require: '%v %v'", err.Field(), err.Value(), err.Tag(), err.Param())
			fmt.Println(msg)
		}

		// from here you can create your own error messages in whatever language you wish
		return
	}

	// save user to database
}

func validateVariable() {

	myEmail := "joeybloggs.gmail.com"

	errs := validate.Var(myEmail, "required,email")

	if errs != nil {
		fmt.Println(errs) // output: Key: "" Error:Field validation for "" failed on the "email" tag
		return
	}

	// email ok, move on
}

