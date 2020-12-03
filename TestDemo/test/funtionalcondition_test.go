package test

import (
	"../demo"
	"fmt"
	"testing"
)

func TestFD(t *testing.T) {
	t.Helper()
	sf := func(t *testing.T, numList []int, f demo.MyFunctionType) (res []bool) {
		res1 := make([]bool, len(numList), len(numList))
		for i, num := range numList {
			res1[i] = f(num)
			//res = append(res, f(num))  // 这里不用再采用:=的方式来定义res了，因为在返回值列表那里已经定义过了
		}
		return res1
	}
	t.Run("将函数当做参数进行传参", func(t *testing.T) {
		numList := []int{1, 2, 3, 4, 5}
		res := sf(t, numList, demo.IsOdd)
		fmt.Printf("numList中为奇数的结果为： %v \n", res)
		res2 := sf(t, numList, demo.IsSS)
		fmt.Printf("numList中为素数的结果为： %v \n", res2)

	})
}
