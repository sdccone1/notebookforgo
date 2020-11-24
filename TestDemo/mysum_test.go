package main

/**
@author：DavidMa
@Content:为demo.Mysum()编写测试用例，注意这里的数组和切片的定义方式以及区别，特别是切片，很类似Java中的ArrayList
@Date : 2020-11-24 21:09
*/
import "testing"
import "./demo"

func TestMysum(t *testing.T) {

	assertCorrectMessage := func(t *testing.T, got, want int) {
		t.Helper()
		if got != want {
			t.Errorf("got: '%d' want: '%d'", got, want)
		}
	}
	t.Run("test1 for demo.MysumForArray()", func(t *testing.T) {
		data := [5]int{1, 2, 3, 4, 5}
		got := demo.MysumForArray(data)
		want := 15
		assertCorrectMessage(t, got, want)
	})
	//下面展示了定义slices切片的两种方式
	t.Run("test1 for demo.MysumForSlices()", func(t *testing.T) {
		dataSlices := []int{1, 2, 3, 4, 5}
		got := demo.MysumForSlices(dataSlices)
		want := 20
		assertCorrectMessage(t, got, want)
	})
	t.Run("test2 for demo.MysumForSlices()", func(t *testing.T) {
		dataSlices := make([]int, 5)
		i := 1
		for idx, _ := range dataSlices {
			dataSlices[idx] = i
			i++
		}
		got := demo.MysumForSlices(dataSlices)
		want := 20
		assertCorrectMessage(t, got, want)
	})
}
