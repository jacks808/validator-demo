# validator-demo

This project is a demo project shows how to use [validator](https://github.com/go-playground/validator/v10) to validate
parameters



# quick start

Step.1: 在proto文件中定义的message中的字段上方，加上@gotags注释，例如

```protobuf
syntax = "proto3";
message Request {
  // @gotags: validate:"gte=0,lte=130"
  int32 Age = 3;
}
```

Step.2:`make proto`生成pb文件[详情](##generate pb.go from proto files)，这样在生成的pb.go文件中的字段后会带上validate属性的tag

```go
package protos

type CreateUserRequest struct {
	Age int32 `protobuf:"varint,1,opt,name=Age,proto3" json:"Age,omitempty" validate:"gte=0,lte=130"`
}
```

"gte=0,lte=130"的含义是大于等于(greater than or equal)0，小于等于(less than or equal)130。

作用于字符串时，保证字符串的utf-8字符个数在[0-130]区间。 作用于slice、map、array时，则保证长度len()在[0-130]区间。 作用于整数、浮点数类型，则保证值的大小在[0-130]区间。

Step.3:在代码中使用

```go
package test

func TestRequest(t *testing.T) {
	// 创建
  req := pb.Request{
		Age:            135,
		CreateTime:     time.Now().Format("2006-01-02 15:04:05"),
	}
	// 调用ValidateStruct方法，对req进行校验
	errorMessage := ValidateStruct(req)
	// 处理错误信息
	for _, err := range errorMessage {
		fmt.Printf("InvalidParameter Error ,field:'%v', current value: '%v', field require: '%v %v'\n",
			err.Field, err.Value, err.Tag, err.Param)
	}
}
```
运行结果：
```text
InvalidParameter Error ,field:'Age', current value: '135', field require: 'lte 130'
InvalidParameter Error ,field:'CreateTime', current value: '2022-01-11 15:25:50', field require: 'datetime 2006-01-02'
```



# use case

### example1

使用常用的tag帮助校验字段。

```go
message CreateUserRequest {
    // @gotags: validate:"required"
    // FirstName
    string FirstName = 1;

    // @gotags: validate:"required"
    // LastName
    string LastName = 2;

    // @gotags: validate:"gte=0,lte=130"
    // Age
    int32 Age = 3;

    // @gotags: validate:"required,email"
    // email
    string Email = 4;

    // @gotags: validate:"datetime=2006-01-02"
    // 时间的格式，参考go原生库的time
    string CreateTime = 6;

    // @gotags: validate:"version"
    // version格式：类似 1.0.1
    string Version = 7;

    // @gotags: validate:"required"
    // friends
    repeated Friend friends = 8;
}

message Friend {
    // @gotags: validate:"required"
    uint32 Id = 1;
    // @gotags: validate:"required"
    string Name = 2;
    // @gotags: validate:"required,number"
    string Group = 3;
}

req := pb.CreateUserRequest{
  FirstName: "Badger",
  LastName:  "Smith",
  Age:       135,
  Email:     "Badger.Smith@gmail.com",
  CreateTime: time.Now().Format("2006-01-02 15:04:05"),
  Version:    "1.0.2.",
  Friends: []*pb.Friend{
    {
      Id:    1,
      Name:  "alice",
      Group: "1",
    },
    {
      Id:    2,
      Name:  "bob",
      Group: "2",
    },
  },
}
errorMessage := ValidateStruct(req)
```

上例中，出现的tag有required、gte、lte、number、version、datetime等。

- 在业务逻辑处理中，常会判断一个字段是否为空，则可以用`required`这个tag来标注此字段。required验证器**验证了该值不是数据类型默认的零值。**对于数字，确保值不为零。 对于字符串，确保值为不是 ""。 对于slice、map、指针、interface、channel和function确保值不为nil。

- 此外，还有比较常用的gte、lte、gt、lt。分别代表着大于等于、小于等于、大于、小于。对于数字，将确保该值限定给定的参数范围。例如gte=10，代表此数字的值应该大于等于10。对于字符串，将确保字符串的utf-8字符个数在参数限定的区间。如果string字段带上gt=5的tag，而值为"alice" 将会产生报错信息，因为字符串的长度应该大于5。对于slice、map和array，它会验证它的实际长度（即调用go的len()方法）是否合法。

- eq代表着等于，对于字符串和数字，eq 将确保该值为等于给定的参数。 对于slice、map和array，将验证它的实际长度（即调用go的len()方法）是否为给定参数。

- ne代表不等于，即保证实际值不等于给定的参数，用法同eq。

- len这个tag可以限定长度，对于数字，将确保**实际的值等于给定的参数**。 对于字符串，它会确保字符串的字符个数等于给定的参数。将验证它的实际长度（即调用go的len()方法）是否等于给定参数。

- number这个tag可以验证string类型字段的值是否为数字。

- version这个tag可以验证类似于 [数字].[数字].[数字]类似结构的字符串。例如"1.23.0"、"1"、"1.0"。

- datetime这个tag会根据提供的日期时间格式验证字符串值是否是有效的日期时间。提供的格式必须与 https://golang.org/pkg/time/ 中记录的官方 Go 时间格式相匹配。例如tag"datetime=2006-01-02",将会匹配"2022-01-11"这样的字符串。需要注意的是layout的所有代表年月日时分秒甚至时区的值都是**互斥不相等**的。

> 比如年份：短年份06，长年份2006，
> 月份：01，Jan，January
> 日：02，2，_2
> 时：15，3，03
> 分：04， 4
> 秒：05， 5

​	这样layout也可以自定义，而且顺序任意，只要符合下列每个区块定义的规则即可。比如 15:04:05-02-01-06，不过没有必要，按容易理解的"2006-01-02 15:04:05"格式来就好。

### example2

自定义验证器。利用go的反射机制，自定义验证器并注册到ti- validate里。 

> 注意: 使用与现有函数相同的标签名称将覆盖现有的

例如有一个服务里，所有的description相关字段都要限制字符串长度为500，那么可以为description相关字段自定义一个验证器，给它起名"desc"，并注册到ti- validate，这样就能在定义字段的时候，加上`// @gotags: validate:"desc"`便可以省去业务逻辑校验代码。

1. 首先定义一个校验函数，入参和出参分别为tiv.FieldLevel和bool

```go

func DescriptionValidator(fl FieldLevel) bool {
  // 获取当前字段 refect.value
	field := fl.Field()

  // 根据当前字段的类型，来做不同的处理
	switch field.Kind() {
	case reflect.String:
		return int64(utf8.RuneCountInString(field.String())) <= 500
	default:
		return false
	}
}
```

2. 注册到ti- validate中，第一个参数为验证器的tag名，第二个参数为校验函数。

```go
RegisterValidator("desc", DescriptionValidator)
```

### example3

给tag注册一个别名。比如

`RegisterAlias("keymax", "max=10")` 那么tag"keymax"便可以等同于"max=10"。



全部validator tag参考：doc.md

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