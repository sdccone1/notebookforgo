package main
import "testing"
import "./demo"
/**
@author:David Ma
@Content:编写测试函数
@Date:2020-11-18 19:33
[Tips]:
1、程序需要在一个名为 xxx_test.go 的文件中编写
2、测试函数的命名必须以单词 Test 开始
3、测试函数只接受一个参数 t *testing.T
*/

func TestHello(t *testing.T) {
    //这样定义方法的方式有点像java的匿名内部类那样，可以在一个方法中去定义实现另一个方法！！
	assertCorrectMessage := func(t *testing.T, got, want string) {
		t.Helper()//t.Helper() 需要告诉测试套件这个方法是辅助函数（helper）。通过这样做，当测试失败时所,报告的错误调用代码的行号是对应的在函数调用中的那个代码所在的位置而不是在辅助函数内部具体发生(打印错误)错误的位置，这里对应的错误代码的行号是26而不再是19了
		if got != want {
			t.Errorf("got '%q' want '%q'", got, want)
		}
	}
	// 下面我们将在我们的测试库中引入另一个工具 —— 子测试(通过t.run(description String,func(t *testing.T){})方法实现)。来使用不同的测试用例来测试我们的API
	t.Run("say something to Lisa", func(t *testing.T) {
			got := demo.Hello("good")
			want := "love you Lisa"
			assertCorrectMessage(t,got,want)
	})
	t.Run("say something to Lisa", func(t *testing.T) {
		got := demo.Hello("bad")
		want := "love you Lisa"
		assertCorrectMessage(t,got,want)
	})
	t.Run("say something to Lisa", func(t *testing.T) {
		got := demo.Hello("")
		want := "love you Lisa"
		assertCorrectMessage(t,got,want)
	})
}
