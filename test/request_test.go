package test

import (
	"github.com/go-playground/validator/v10"
	"testing"
	"validator-demo/tivalidator"
)

func TestRequest(t *testing.T) {
	validate := validator.New()
	validate.RegisterValidation("public", tivalidator.Public, true)

}
