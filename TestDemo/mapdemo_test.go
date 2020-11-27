package main

/**
@author:David Ma
@Content:测试demo包下Equals()，并展示map的若干种定义方法
@Date:2020-11-27 15:33
*/
import (
	"./demo"
	"fmt"
	"testing"
)

func TestEquals(t *testing.T) {
	assertCorrectMessage := func(t *testing.T, map1, map2 map[string]int) {
		t.Helper()
		fmt.Println(demo.Equals(map1, map2))
	}
	//下面展示了创建map的几种方法
	map1 := map[string]int{
		"Scott": 23,
		"Bob":   16,
		"Kitty": 18,
	}

	map2 := make(map[string]int)
	map2["Scott"] = 23
	map2["Bob"] = 16
	map2["Kitty"] = 18

	var map3 = make(map[string]int)
	map3["Scott"] = 23
	map3["Bob"] = 16
	map3["Kitty"] = 18

	t.Run("the first testing for demo.Equals()", func(t *testing.T) {
		assertCorrectMessage(t, map1, map2)
	})
	t.Run("the second testing for demo.Equals()", func(t *testing.T) {

		assertCorrectMessage(t, map1, map3)
	})
	t.Run("the third testing for demo.Equals()", func(t *testing.T) {

		assertCorrectMessage(t, map2, map3)
	})
	t.Run("the fourth testing for demo.Equals()", func(t *testing.T) {
		delete(map3, "Scott")
		assertCorrectMessage(t, map2, map3)
	})
	t.Run("the fifth testing for demo.Equals()", func(t *testing.T) {
		map3["Scott"] = 23
		assertCorrectMessage(t, map2, map3)
	})
}
