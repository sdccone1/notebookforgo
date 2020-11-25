package demo

func MysumForArray(data [5]int) (sum int) { //MysumForArray()的参数必须是长度为5的int型数组
	var totalSum int = 0
	for i := 0; i < len(data); i++ {
		totalSum += data[i]
	}
	return totalSum
}
func MysumForSlices(data []int) (sum int) { //此时MysumForSlices的参数就是一个slices切片了
	var totalSum int = 0
	//range 关键字可以修饰一个集合(可迭代的对象)其返回两个值：(index , value),且可以 利用range 关键字来遍历修改一个集合

	//先令原切片data的每个元素自增1
	for idx, _ := range data {
		data[idx]++
	}
	//再求和
	/*
		每次循环迭代，range产生一对值；索引以及在该索引处的元素值。这个例子不需要索引，但range的语法要求, 要处理元素, 必须处理索引。
		一种思路是把索引赋值给一个临时变量, 如temp, 然后忽略它的值，但Go语言不允许使用无用的局部变量（local variables），因为这会导致编译错误。
		Go语言中这种情况的解决方法是用空标识符（blank identifier），即_（也就是下划线）。空标识符可用于任何语法需要变量名但程序逻辑不需要的时候, 例如, 在循环里，丢弃不需要的循环索引, 保留元素值
	*/
	for _, num := range data { // 使用空白标识符(下划线)来忽略原始值
		totalSum += num
	}
	return totalSum
}
