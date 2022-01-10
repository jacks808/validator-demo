package test

import (
	"fmt"
	"strings"
	"testing"
)

func Test_unit(t *testing.T) {

	s := strings.SplitN("a,b,c,d,e,f", ",", 2)
	for _, s2 := range s {
		fmt.Println(s2)
	}

}
