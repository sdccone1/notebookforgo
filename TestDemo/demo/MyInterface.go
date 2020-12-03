package demo

import "fmt"

/*
@Author:David Ma
@Content:接口相关定义
@Date:2020-12-03 15:24
*/

type EatService interface {
	Eat()
	Sleep()
}

type Cat struct {
	Name string
}

func (cat *Cat) Eat() {
	fmt.Printf("%s eatting \n", cat.Name)
}

func (cat *Cat) Sleep() {
	fmt.Printf("%s sleeping\n", cat.Name)
}

type Dog struct {
	Name string
}

func (dog *Dog) Sleep() {
	fmt.Printf("%s sleeping\n", dog.Name)
}

func (dog *Dog) Eat() {
	fmt.Printf("%s eatting \n", dog.Name)
}

func MyEat(es EatService) {
	es.Eat()
}
func MySleep(es EatService) {
	es.Sleep()
}
