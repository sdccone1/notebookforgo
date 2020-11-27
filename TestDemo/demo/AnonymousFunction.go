package demo

/**
@author:David Ma
@Content:
匿名函数
@Date:2020-11-18 21:07
*/
func Squares() func() int { //squares()返回一个匿名函数这个匿名函数又返回1个int值
	var x int
	return func() int {
		x++
		return x * x
	}
}
