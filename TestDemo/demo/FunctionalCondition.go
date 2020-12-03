package demo

import "math"

/*
@Author:David Ma
@Content:将函数当做参数进行方法传参
@Date:2020-12-03 10:46
*/

// step1 :先定义一个对应的函数类型，方便将函数作为一种类型进行传参
type MyFunctionType func(int) bool

// step2:定义测试函数 IsOdd(int)
func IsOdd(num int) bool {
	if num%2 == 0 {
		return false
	}
	return true
}

func IsSS(num int) bool {
	sqrtOfNUm := math.Sqrt(float64(num))
	for i := 2; float64(i) <= sqrtOfNUm; i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}
