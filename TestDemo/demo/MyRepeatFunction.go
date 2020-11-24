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
	// case2 使用 := 此时它相当于 var name string = ""
	//result := ""
	for i := 1; i <= repeatedCounts; i++ {
		result = result + s
	}
	return result
}
