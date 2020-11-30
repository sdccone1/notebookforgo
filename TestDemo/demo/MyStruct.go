package demo

/*
@author:David Ma
@content: struct结构体相关
@Date:2020-11-30
*/

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
	username string
	password string
	// 测试匿名成员和非匿名成员的区别
	address1 Address
	//在一个struct中最多只允许同时有1个同类型的匿名成员，因为匿名成员隐含的以其类型本身为名
	Address
}

func SetUserName(user *User, name string) bool {
	if user == nil {
		return false
	}
	user.username = name
	return true
}
func SetPhone(user *User, phoneNum string) bool {
	if user == nil {
		return false
	}
	user.phone = phoneNum
	return true
}
func SetPassword(user *User, password string) bool {
	if user == nil {
		return false
	}
	user.password = password
	return true
}
func SetAddress(user *User, country, city, street string) bool {
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

func GetPhone(user *User) (phone string) {
	return user.phone
}

func GetUserName(user *User) (name string) {
	return user.username
}

func GetPassWord(user *User) (pwd string) {
	return user.password
}

func GetAddress(user *User) (address []string) {
	//case1:通过匿名成员的特性来直接获取被嵌套的成员的属性
	//虽然user中的成员均不是导出的(私有的)，但是因为此时同属于用一个包所以非导出成员也可以被访问到
	add := user.country + user.city + "市" + user.street + "街道"
	//case2:传统方式获取被嵌套的成员的属性
	add2 := user.Address.country + user.Address.city + "市" + user.Address.street + "街道"
	addList := []string{add, add2}
	return addList
}
