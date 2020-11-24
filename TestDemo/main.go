package main
import "fmt"
import "./demo"
/**
@author:David Ma
@Content:
go的跨模块调用实现：
1、先把.go文件放在相应的package下
2、然后手动import所需要的模块所在的package
3、最后，在调用相应的方法时通过 package_Name.func_Name()的形式来完成调用
@Date:2020-11-18 20:07
*/
func saySomething2(prefix, name string){
	fmt.Println(prefix + " " + name)
}

func main()  {
	//print()与printf()的区别为printf(String format,T data)支持格式化输出
	//fmt.Println("Helllo world")
	//fmt.Printf("%[2]d %4.2[1]f",12.11,2)
	fmt.Println(demo.Hello("bad"))
	//效果上等价于func saySomething2(prefix, name string){....},不同的是使用这种形式，可以在一个方法中去定义实现一个私有方法(该私有方法仅能在其外围方法的作用域内被访问)
	saySomething := func(prefix, name string) {
		fmt.Println(prefix + " " + name)
	}
	saySomething("Hello","Lisa")
	saySomething2("Hello","Lisa")
}

