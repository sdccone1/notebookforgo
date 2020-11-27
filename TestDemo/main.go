package main

/**
@author:DavidMa
@Content: 展示demo包下的MyConst.go以及AnonymousFunction.go的结果
@Date: 2020-11-27 13:32
*/
import "./demo"

func main() {
	f := demo.Squares()
	//观察结果可以得知，局部变量x的生命周期并不取决于其作用域，而是取决于这个变量是否被引用，此时将squares()赋值给f，此后不论调用多少次f()，此时都相当于调用的是同一个squares()
	//而不是每调用一次f()都相当于重新调用1次squares()
	//所以局部变量x会一直被f所引用，所以才会有一下的输出结果
	for i := 0; i < 4; i++ {
		println(f())
	}
	demo.MyConst()
}
