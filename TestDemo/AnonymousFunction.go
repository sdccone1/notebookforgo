package main

/**
@author:David Ma
@Content:
匿名函数
@Date:2020-11-18 21:07
*/
func squares() func() int { //squares()返回一个匿名函数这个匿名函数又返回1个int值
	var x int
	return func() int {
		x++
		return x * x
	}
}
func main() {
	f := squares()
	//观察结果可以得知，局部变量x的生命周期并不取决于其作用域，而是取决于这个变量是否被引用，此时将squares()赋值给f，此后不论调用多少次f()，此时都相当于调用的是同一个squares()
	//而不是每调用一次f()都相当于重新调用1次squares()
	//所以局部变量x会一直被f所引用，所以才会有一下的输出结果
	for i := 0; i < 4; i++ {
		println(f())
	}
}
