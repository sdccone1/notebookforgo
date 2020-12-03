package main

import (
	"./demo"
	"bytes"
	"fmt"
	//"sync"
)

func main() {
	f := demo.Squares()
	fmt.Printf("f = '%v'\n", f)
	//观察结果可以得知，局部变量x的生命周期并不取决于其作用域，而是取决于这个变量是否被引用，此时将squares()赋值给f，此后不论调用多少次f()，此时都相当于调用的是同一个squares()
	//而不是每调用一次f()都相当于重新调用1次squares()
	//所以局部变量x会一直被f所引用，所以才会有一下的输出结果

	for i := 0; i < 4; i++ {
		println(f())
	}
	demo.MyConst()
	a := 1
	b := 2
	fmt.Println(^a)
	fmt.Println(a ^ b)
	s1 := "hello"
	s2 := "world"
	c := s1 + string(1)
	// 借助字节缓冲区byteBuffer来实现类似Java中StringBuilder的效果，因为在go中string底层是由byte数组实现的
	StringBuilder := bytes.Buffer{}
	StringBuilder.WriteString(s1)
	StringBuilder.WriteString(" ")
	StringBuilder.WriteString(s2)
	fmt.Println(StringBuilder.String())
	e := 'a'
	fmt.Println(e)
	fmt.Printf("%c \n", c)
	//ConcurrentHashMap := sync.Map{}

}
