#### 如何使用多个验证器

一个字段上的多个验证器将按照定义的顺序进行处理。例子：将检查 max 然后 min

```go
type Test struct {
  Field `validate:"max=10,min=1"`
}
```

错误的验证器定义将不会成功。例子：

```go
type Test struct {
  Field `validate:"min=10,max=0"`
}
```

逗号 (",") 是验证器tag的默认分隔符。如果你想在参数中包含一个逗号（即 excludesall=,），需要使用 UTF-8 十六进制表示 0x2C，在代码中替换为逗号，所以上面会变成 excludesall=0x2C。

```go
type Test struct {
  Field `validate:"excludesall=,"`    // 不要包含逗号
  Field `validate:"excludesall=0x2C"` // 正确用法
}
```

管道（"|"）是"或"验证标签分隔符。如果你想在参数中包含一个管道，即 excludesall=|，需要使用 UTF-8 十六进制表示 0x7C，所以上面会变成 excludesall=0x7C

```go
  type Test struct {
      Field `validate:"excludesall=|"`    // BAD! Do not include a a pipe!
      Field `validate:"excludesall=0x7C"` // GOOD! Use the UTF-8 hex representation.
   }
```

#### 内置验证器

##### 跳过操作符

告诉验证器跳过这个结构体字段 （用法： -）

##### 或操作符

多个tag直接的关系为 **或**。 
(用法: omitempty,rgb|rgba)

##### Omit Empty

允许条件验证，如果字段未设置值，则其他验证器不会生效（例如 min 或 max），如果设置了值，则其他验证器将生效。

##### Dive

这告诉验证器深入到slice、array或map中并验证slice、array或map的级别以及后面的验证tag。
还支持多维嵌套，您希望的每个级别都会需要另一个潜水标签。 潜水有一些子标签，'keys' & 'endkeys'，请看
下面的 Keys & EndKeys 部分。

```
Example #1

   [][]string with validation tag "gt=0,dive,len=1,dive,required"
   // gt=0 will be applied to []
   // len=1 will be applied to []string
   // required will be applied to string

Example #2

   [][]string with validation tag "gt=0,dive,dive,required"
   // gt=0 将被作用于 []
   // []string will be spared validation
   // required will be applied to string
```

##### Keys & EndKeys

这些将在dive标签之后直接一起使用并告诉验证器'keys' 和 'endkeys' 之间的任何东西都适用于map的key，而不是
Value； 将其视为“dive”标签，但用于映射键而不是值。还支持多维嵌套，您希望验证的每个级别都会需要另一个 'keys' 和 'endkeys' 标签。 这些标签仅对map有效。

   Usage: dive,keys,othertagvalidation(s),endkeys,valuevalidationtags

```
Example #1

   map[string]string with validation tag "gt=0,dive,keys,eg=1|eq=2,endkeys,required"
   // gt=0 will be applied to the map itself
   // eg=1|eq=2 will be applied to the map keys
   // required will be applied to map values

Example #2

   map[[2]string]string with validation tag "gt=0,dive,keys,dive,eq=1|eq=2,endkeys,required"
   // gt=0 will be applied to the map itself
   // eg=1|eq=2 will be applied to each array element in the the map keys
   // required will be applied to map values
```

##### required

验证该值不是数据类型默认的零值。对于数字，确保值不为零。 对于字符串，确保值为不是""。 对于切片、映射、指针、接口、通道和函数确保值不为nil。

##### isdefault

验证该字段的值是默认值。对于数字，确保值为零。 对于字符串，确保值为""。 对于切片、映射、指针、接口、通道和函数确保值是nil。跟required相反.

##### len

对于数字，长度将确保该值等于给定的参数。 对于字符串，将检查字符串长度是否等于给定参数。 对于切片、数组和映射，将验证它们项的个数（即len()）等于给定的参数。

使用举例：`validate:"len=10"`

##### max

对于数字， max 将确保该值是小于或等于给定的参数。 对于字符串，它检查字符串长度小于等于给定参数。 为了
切片、数组和映射，将验证项的个数（即len()）等于给定的参数。

使用举例：`validate:"max=10"`

##### min

用法同max。

##### equals

用法同max，验证是否相等。

使用举例：`validate:"eq=10"`

##### not equal

用法跟equal一样，验证是否不相等。

使用举例：`validate:"ne=10"`

##### One Of

对于字符串、整数，oneof 将确保该值是参数中的值之一。 该参数应该是由**空格**分隔的值列表。 值可以是字符串或数字。 要匹配带有空格的字符串，请将目标字符串包含在单引号之间。

使用举例：

    oneof=red green
    oneof='red green' 'blue yellow'
    oneof=5 7 9

##### Greater Than

用法类似min，确保实际值大于给定的参数。

使用举例：`validate:"gt=10"`

##### Greater Than or Equal

用法同min，确保实际值大于等于给定的参数。

使用举例：`validate:"gte=10"`

##### Less Than

用法类似Greater Than。

使用举例：`validate:"lt=10"`

##### Less Than or Equal

用法同max，确保实际值小于等于给定的参数。

使用举例：`validate:"lte=10"`

##### Unique

对于slice和array，将确保没有重复的值。
对于map，将确保没有重复的value。
对于slice结构体，unique 将确保通过参数指定的结构字段中没有重复值。

使用举例：

```
// 对于slice和array和map:
Usage: unique

// 对于slice结构体
Usage: unique=field
```

##### Alpha Only

验证字符串值仅包含 ASCII 字母字符

 使用举例：`validate:"alpha"`

##### Alphanumeric

这验证字符串值仅包含 ASCII 字母数字字符

使用举例：`validate:"alphanum"`

##### Boolean

验证了一个字符串值可以成功地用 strconv.ParseBool 解析成布尔值

使用举例：`validate:"boolean"`

##### Number

这验证字符串值仅包含数字值(整数或浮点数都行)。

使用举例：`validate:"number"`

##### Lowercase String

这验证字符串值仅包含小写字符. 空字符串不是有效的小写字符串。

 使用举例：`validate:"lowercase"` 

##### Uppercase String

这验证字符串值仅包含大写字符. 空字符串不是有效的大写字符串。

 使用举例：`validate:"uppercase"` 

##### E-mail String

这将验证字符串值是否包含是有效的电子邮件。

 使用举例：`validate:"email"` 

##### JSON String

这将验证字符串值是否包含是有效的json。

使用举例：`validate:"json"` 

##### File path

这将验证字符串值是否包含有效的文件路径，并且该文件存在于机器上。
这是使用 os.Stat 完成的，它是一个独立于平台的函数。

使用举例：`validate:"file"` 

##### URL String

这将验证字符串值是否包含有效的 url
这将接受 golang 请求 uri 接受的任何 url，但必须包含例如 http:// 或 rtmp://

使用举例：`validate:"url"` 

URI String

这将验证字符串值是否包含有效的 uri（golang 请求 uri 接受的任何 uri）

使用举例：`validate:"uri"` 

##### Contains

这将验证字符串的值是否包含给定参数的值（即给定参数是子字符串）。

使用举例：`validate:"contains=a"` 

##### Contains Any

这将验证字符串值是否包含子字符串值中的任何 Unicode。

使用举例：`validate: "containsany=!@#?"`

##### Excludes

这将验证字符串值不包含给定参数的值（即给定参数不是子字符串）。是Contains的反向用法。

 使用举例：`validate:"excludes=a"` 

##### Excludes All

这将验证字符串值不包含子字符串值中的任何 Unicode。

 使用举例：`validate:"excludesall=!@#"` 

##### Starts With

这验证字符串值是否以给定的字符串值开头

 使用举例：`validate:"startswith=hello`"

##### Ends With

这验证字符串值是否以给定的字符串值结尾

 使用举例：`validate:"endswith=goodbye`"

##### Does Not Start With

跟Starts With相反。

使用举例：`validate:"startsnotwith=hello`"

##### Does Not End With

跟Ends With相反

使用举例：`validate:" endsnotwith=goodbye`"

##### IP

这将验证字符串值是否包含有效的 IP 地址。

使用举例：`validate:"ip"`

##### IPv4

这将验证字符串值是否包含有效的 IPv4 地址。

使用举例：`validate:"ipv4"`

##### IPv6

这将验证字符串值是否包含有效的 IPv6 地址。

使用举例：`validate:"ipv6"`

##### Directory

这将验证字符串值是否包含有效目录，并且它存在于机器上。
这是使用 os.Stat 完成的，它是一个独立于平台的函数。

使用举例：`validate:"dir"`

##### HostPort

这将验证字符串值是否包含有效的 DNS 主机名和端口。

使用举例：`validate:"hostname_port"`

##### Datetime

这会根据提供的日期时间格式验证字符串值是否是有效的日期时间。
提供的格式必须与 https://golang.org/pkg/time/ 中记录的官方 Go 时间格式布局相匹配

使用举例：`validate:"datetime=2006-01-02"`