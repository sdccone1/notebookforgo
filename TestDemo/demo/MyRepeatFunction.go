package demo

/**
@author:DavidMa
@Content:go中循环只有一个for，接下来用go实现strings包的repeat()
@Date:2020-11-24
*/
func MyRepeat(s string, repeatedCounts int) string {
	// go中有两种声明变量的方式
	// case1 通过var标识符
	var result string
	/* case2 使用 := 此时它相当于 var name string = "",但只能在方法体内使用 := 方式来声明并初始化一个变量而不能在一个方法体外使用 ：= 的方式来初始化一个变量，
	且要区分开：= 和 = 的区别前者是声明一个局部变量，而后者仅是一个赋值操作
	*/
	//result := ""
	for i := 1; i <= repeatedCounts; i++ {
		result = result + s
	}
	return result
}
