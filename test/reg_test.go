package test

import (
	"fmt"
	"strings"
	"testing"
)

func TestReg(t *testing.T) {
	//match, err := regexp.MatchString("^[0-9]([0-9]+\\.)+[0-9]+", "x.1.3.214.x")
	//fmt.Println(match, err)
	//if err == nil {
	//	if match {
	//		fmt.Println("匹配成功")
	//	} else {
	//		fmt.Println("匹配失败")
	//	}
	//}
	s := "123.2350.42"
	split := strings.Split(s, ".")
	for _, str := range split {
		if len(str) == 0 {
			fmt.Println("匹配失败")
			return
		}
		for i := 0; i < len(str); i++ {
			if str[i] < '0' || str[i] > '9' {
				fmt.Println("匹配失败")
				return
			}
		}
	}
	fmt.Println("匹配成功")
}
