package _1_string

import (
	"bytes"
	"fmt"
	"strings"
)

func StringJoint() {
	str1, str2 := "go", "lang"
	fmt.Println(str1 + str2)
	fmt.Println(fmt.Sprintf("%s%s", str1, str2))
	fmt.Println(strings.Join([]string{str1, str2}, ""))

	var bt bytes.Buffer
	bt.WriteString(str1)
	bt.WriteString(str2)
	fmt.Println(bt.String())

	var build strings.Builder
	build.WriteString(str1)
	build.WriteString(str2)
	fmt.Println(build.String())
}
