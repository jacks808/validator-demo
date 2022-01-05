package tivalidator

import (
	"github.com/go-playground/validator/v10"
	"reflect"
	"strings"
)

func Version(fl validator.FieldLevel) bool {
	field := fl.Field()
	switch field.Kind() {
	case reflect.String:
		// example: 1.0.0  2.1.23
		s := field.String()
		split := strings.Split(s, ".")
		for _, str := range split {
			if len(str) == 0 {
				return false
			}
			for i := 0; i < len(str); i++ {
				if str[i] < '0' || str[i] > '9' {
					return false
				}
			}
		}
		return true
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return field.Int() > 0
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return true
	default:
		return false
	}
}
