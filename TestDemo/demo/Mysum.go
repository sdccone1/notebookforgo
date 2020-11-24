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
	for _, num := range data { // 使用空白标识符(下划线)来忽略原始值
		totalSum += num
	}
	return totalSum
}
