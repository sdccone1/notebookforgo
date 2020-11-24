package main

/**
@author：DavidMa
@Content:为demo.Mysum()编写测试用例，注意这里的数组定义方式
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
	t.Run("test1 for demo.Mysum()", func(t *testing.T) {
		data1 := []int{1, 2, 3, 4, 5}
		got := demo.Mysum(data1)
		want := 15
		assertCorrectMessage(t, got, want)
	})
	t.Run("test2 for demo.Mysum()", func(t *testing.T) {
		data2 := []int{1, 2, 3, 4, 5, 6}
		got := demo.Mysum(data2)
		want := 15
		assertCorrectMessage(t, got, want)
	})
}
