package tivalidator

import (
	"fmt"
	"reflect"
	"strconv"
)

func Port(fl FieldLevel) bool {
	field := fl.Field()
	switch field.Kind() {
	case reflect.String:
		if portNum, err := strconv.ParseInt(field.String(), 10, 32); err != nil || portNum > 65535 || portNum < 1 {
			return false
		}
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64, reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return field.Int() >= 0 && field.Int() <= 65535
	default:
		return false
	}
	panic(fmt.Sprintf("Bad field type %T", field.Interface()))
}
