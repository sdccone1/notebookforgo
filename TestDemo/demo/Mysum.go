package demo

func Mysum(data []int) int {
	var totalSum int = 0
	for i := 0; i < len(data); i++ {
		totalSum += data[i]
	}
	return totalSum
}
