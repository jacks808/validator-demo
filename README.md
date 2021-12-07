# validator-demo
This project is a demo project shows how to use [validator](github.com/go-playground/validator/v10) to validate parameters

# use case
## install requirements
```shell script
go get github.com/favadi/protoc-go-inject-tag
```
see more [details](https://github.com/favadi/protoc-go-inject-tag)
## generate pb.go from proto files
```shell script
make proto

> ./proto/demo.pb.go
> protoc --proto_path=./protos --go_out=paths=source_relative:./proto ./protos/*.proto
> protoc-go-inject-tag -input="././proto/*.pb.go"
```

## run 
```shell script
make run

> InvalidParameterError
> error 1, field: 'Age', current value: '135', field require: 'lte 130'
> error 2, field: 'FavouriteColor', current value: '#000-', field require: 'iscolor '
```