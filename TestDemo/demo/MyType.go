package demo

/*
@Author:David Ma
@Content:go中的面向对象编程(oop)的实现，以及类型别名和组合在使用时的区别
@Date:202-12-01 15:41
*/
import (
	"fmt"
)

type Person struct {
	Name string
}

// 类型别名
type Student2 Person

//类型组合
type Student1 struct {
	Person
}

// go中oop的实现，在函数名前加上一个声明(Type *Type)表明该方法是属于哪个对象(Type)
func (p *Person) WakeUp() {
	fmt.Printf("%s waking up\n", p.Name)
}

func (p *Person) Eat() {
	fmt.Printf("%s eating\n", p.Name)
}

func (p *Person) Sleep() {
	fmt.Printf("%s sleeping\n", p.Name)
}

func (s *Student1) Study() {
	fmt.Printf("%s learning\n", s.Name)
}

func (s *Student2) Study() {
	fmt.Printf("%s learning\n", s.Name)
}

func (s *Student1) Student1Daily() {
	s.WakeUp()
	s.Eat()
	s.Study()
	s.Sleep()
}

//虽然说类型别名没有创建新类型，只是换了个名字起了个小名，但是在OOP中，无法利用类型别名来实现代码复用，只能利用组合来实现代码复用
//func (s *Student2) Student2Daily() {
//	s.WakeUp() // compile error:Unresolved reference 'WakeUp'
//	s.Eat()  // compile error:Unresolved reference 'WakeUp'
//	s.Study()
//	s.Sleep() // compile error:Unresolved reference 'WakeUp'
//}
