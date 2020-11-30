package test

import (
	"../demo"
	"fmt"
	"testing"
)

func TestMyStruct(t *testing.T) {

	t.Run("如何定义并初始化一个struct", func(t *testing.T) {
		// case1：使用 new()函数
		user := new(demo.User)
		demo.SetPhone(user, "12345678901")
		demo.SetUserName(user, "Bob")
		demo.SetPassword(user, "11111")
		demo.SetAddress(user, "中国", "苏州", "平江")

		//case2:还可以使用字面常量的方式来定义，只不过由于User的成员不是导出的，所以此时由于不在一个包下，无法访问到其成员，且匿名成员无法通过这种字面常量的方式来定义
		//u2 := demo.User{
		//	phone : "12345678901"
		//	...
		//}

		fmt.Printf("name = '%s' phone = '%s' password = '%s' address1 = '%s' address2 = '%s' \n",
			demo.GetUserName(user), demo.GetPhone(user), demo.GetPassWord(user), demo.GetAddress(user)[0], demo.GetAddress(user)[1])
	})
}
