package test

import (
	log "github.com/sirupsen/logrus"
	"testing"
	"time"
	pb "validator-demo/proto"
	"validator-demo/tivalidator"
)

func TestRequest(t *testing.T) {
	validate := tivalidator.NewValidator()
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
	tivalidator.RegisterTag(validate, "version", tivalidator.Version)
	errorMessage := tivalidator.ValidateStruct(validate, req)
	if errorMessage != "" {
		// 业务层自己处理
		log.Errorf(errorMessage)
	}
}
