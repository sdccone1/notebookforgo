package demo

import "fmt"

/*
@author:David Ma
@content: struct结构体相关
@Date:2020-11-30
*/
var testInt int = 10

type Movie struct {
	Title  string
	Year   int  `json:"released"`
	Color  bool `json:"color,omitempty"`
	Actors []string
}

type Address struct {
	country, city, street string
}

type User struct {
	phone    string
	username string `lowerName:"name" type:"varchar(50)"` // struct的tag信息`k:v`
	password string
	// 测试匿名成员和非匿名成员的区别
	address1 Address
	//在一个struct中最多只允许同时有1个同类型的匿名成员，因为匿名成员隐含的以其类型本身为名
	Address
}

func (user *User) SetName(name string) bool {
	if user == nil {
		return false
	}
	user.username = name
	return true
}

func (User) Test1() {
	fmt.Println("this is func Test1")
}
func (User) Test2() {
	fmt.Println("this is func Test2")
}
func (User) Test3() {
	fmt.Println("this is func Test3")
}
func (user *User) SetUserName(name string) bool {
	if user == nil {
		return false
	}
	user.username = name
	return true
}
func (user *User) SetPhone(phoneNum string) bool {
	if user == nil {
		return false
	}
	user.phone = phoneNum
	return true
}
func (user *User) SetPassword(password string) bool {
	if user == nil {
		return false
	}
	user.password = password
	return true
}
func (user *User) SetAddress(country, city, street string) bool {
	if user == nil {
		return false
	}
	user.address1.country = country
	user.address1.city = city
	user.address1.street = street
	user.country = country
	user.city = city
	user.street = street
	return true
}

func (user *User) GetPhone() (phone string) {
	return user.phone
}

func (user *User) GetUserName() (name string) {
	return user.username
}

func (user User) GetName() (name string) {
	return user.username
}

func (user *User) GetPassWord() (pwd string) {
	return user.password
}

func (user *User) GetAddress() (address1, address2 string) {
	//case1:通过匿名成员的特性来直接获取被嵌套的成员的属性
	//虽然user中的成员均不是导出的(私有的)，但是因为此时同属于用一个包所以非导出成员也可以被访问到
	add := user.country + user.city + "市" + user.street + "街道"
	//case2:传统方式获取被嵌套的成员的属性
	add2 := user.Address.country + user.Address.city + "市" + user.Address.street + "街道"
	return add, add2
}
