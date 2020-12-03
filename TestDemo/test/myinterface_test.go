package test

import (
	"../demo"
	"testing"
)

/*
@Author:David Ma
@Content:接口相关定义
@Date:2020-12-03 15:24
*/

func TestMyInterface(t *testing.T) {
	t.Helper()
	t.Run("", func(t *testing.T) {
		cat := demo.Cat{Name: "Tom"}
		dog := demo.Dog{Name: "Jerry"}
		// 注意这里要传一个指针类型
		//否则compile error:Cannot use 'cat' (type Cat) as type EatService Type does not implement 'EatService' as 'Eat' method has a pointer receiver
		demo.MyEat(&cat)
		demo.MySleep(&cat)
		//如果对象实现了接口的某个方法，则其必须实现该接口的其他方法，否则报错：
		//Cannot use '&dog' (type *Dog) as type EatService Type does not implement 'EatService' as some methods are missing: Sleep()
		demo.MyEat(&dog)
		demo.MySleep(&dog)
	})
}
