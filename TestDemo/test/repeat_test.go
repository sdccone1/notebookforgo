package test

/**
@author:David Ma
@Content:对demo.MyRepeat()编写测试函数
@Date:2020-11-24 20:47
*/
import (
	"../demo"
	"strings"
	"testing"
)

func TestMyRepeat(t *testing.T) {
	assertCorrectMessage := func(t *testing.T, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got: '%q' want : '%q'", got, want)
		}
	}
	t.Run("test1 for MyRepeat()", func(t *testing.T) {
		const char = "a"
		got := demo.MyRepeat(char, 5)
		want := strings.Repeat(char, 5)
		assertCorrectMessage(t, got, want)
	})
	t.Run("test2 for MyRepeat()", func(t *testing.T) {
		const char = "b"
		got := demo.MyRepeat(char, 6)
		want := strings.Repeat(char, 6)
		assertCorrectMessage(t, got, want)
	})
}
