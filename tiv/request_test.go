package tiv

import (
	"fmt"
	"testing"
	"time"
	pb "validator-demo/proto"
)

func TestRequest(t *testing.T) {
	req := pb.CreateUserRequest{
		FirstName:      "Badger",
		LastName:       "Smith",
		Age:            135,
		Email:          "Badger.Smith@gmail.com",
		FavouriteColor: "#000-",
		// format : 2006-01-02 15:04:05
		CreateTime: time.Now().Format("2006-01-02 15:04:05"),
		Version:    "1.0.2.",
	}
	RegisterValidator("version", Version)
	errorMessage := ValidateStruct(req)
	for _, err := range errorMessage {
		fmt.Printf("InvalidParameter Error ,field:'%v', current value: '%v', field require: '%v %v'\n",
			err.Field, err.Value, err.Tag, err.Param)
	}
}
