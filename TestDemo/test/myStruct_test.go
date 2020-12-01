package test

/*
@Author:David Ma
@Content:struct结构体相关操作
@Date:2020-12-01 15:03
*/
import (
	"../demo"
	"encoding/json"
	"fmt"
	"log"
	"testing"
)

func TestMyStruct(t *testing.T) {

	// case1：使用 new()函数
	t.Run("如何定义并初始化一个struct之使用new()函数来初始化一个struct", func(t *testing.T) {

		user := new(demo.User)
		demo.SetPhone(user, "12345678901")
		demo.SetUserName(user, "Bob")
		demo.SetPassword(user, "11111")
		demo.SetAddress(user, "中国", "苏州", "平江")
		add1, add2 := demo.GetAddress(user)
		fmt.Printf("name = '%s' phone = '%s' password = '%s' address1 = '%s' address2 = '%s' \n",
			demo.GetUserName(user), demo.GetPhone(user), demo.GetPassWord(user), add1, add2)
		//case2:还可以使用字面常量的方式来定义，
		type Animal struct {
			Age int
		}

		type Person struct {
			Animal
			Name string
		}

		type Student struct {
			Person
			ClassName string
		}

	})

	//case2:还可以使用字面常量的方式来定义，
	t.Run("如何定义并初始化一个struct之使用字面常量的方式来初始化一个struct", func(t *testing.T) {

		type Animal struct {
			Age int
		}

		type Person struct {
			Animal
			Name string
		}

		type Student struct {
			Person
			ClassName string
		}
		s := Student{
			Person{
				Animal: Animal{ //注意这里匿名成名如何使用字面常量的方式来初始化
					Age: 18,
				},
				Name: "Bob",
			},
			"Class01",
		}
		fmt.Printf("%+v \n", s)

	})

	// 只有导出的结构体成员才会被编码，这也就是我们为什么选择用大写字母开头的成员名称
	t.Run("定义一个Movie类型(其成员均可导出，方便编解码操作)的切片，并展示将其编码为json以及反编码", func(t *testing.T) {
		data := []demo.Movie{
			{Title: "Casablanca", Year: 1942, Color: false,
				Actors: []string{"Humphrey Bogart", "Ingrid Bergman"}},
			{Title: "Cool Hand Luke", Year: 1967, Color: true,
				Actors: []string{"Paul Newman"}},
			{Title: "I Can Fly", Year: 1968, Color: true,
				Actors: []string{"Steve McQueen", "Jacqueline BBB"}},
		}
		//marshal
		msg, err := json.MarshalIndent(data, "", " ")
		if err != nil {
			log.Fatalf("Json marshal faliled '%s'", err)
		}
		fmt.Printf("JSON串为：%s\n", msg)

		//unmarshal(下面只解码Title和Actors信息)
		var receivedTitles []struct {
			Title  string
			Actors []string
		}
		//json.Unmarshal(@args1:json串，@argus2:一个用于接收解码后的数据的结构)，这里特别注意arus2必须为指针类型，因为要对其做修改，而go又是值传递
		err2 := json.Unmarshal(msg, &receivedTitles)
		if err2 != nil {
			log.Fatalf("Json unmarshal faliled '%s'\n", err2)
		}
		fmt.Printf("the Title of Movies is:'%s' \n", receivedTitles)
	})
}
