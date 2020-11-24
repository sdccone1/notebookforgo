package demo

func Mysum(data []int) int { //此时Mysum()的参数必须要是int 类型的数组 且不能指定具体的长度 []int
	var totalSum int = 0
	for i := 0; i < len(data); i++ {
		totalSum += data[i]
	}
	return totalSum
}
