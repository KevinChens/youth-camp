package method

import (
	"fmt"
	"testing"
)

func TestUser_ModifyInfo1(t *testing.T) {
	u := User{Name: "zhangsan", Age: 19}
	u.ModifyInfo1()
	fmt.Println(u)
}

func TestUser_ModifyInfo2(t *testing.T) {
	u := &User{Name: "zhangsan", Age: 19}
	u.ModifyInfo2()
	fmt.Println(u)
}