package demo

/**
@author:DavidMa
@Content: iota微小常量生成器的用法，以及如何定义多个同类型常量
@Date: 2020-11-27 13:32
*/
import "fmt"

func MyConst() {
	const (
		a = iota //在第一个声明的常量所在的行，iota将会被置为0，然后在每一个有常量声明的行加一。
		b
		c
		d
		e
	)
	fmt.Printf("a = '%d' b = '%d' c = '%d' d = '%d' e = '%d' \n", a, b, c, d, e)
}
