package test

/*
@Author:David Ma
@Date:2020-12-07
@Content:reflect_demo
*/
import (
	"../demo"
	"fmt"
	"reflect"
	"testing"
)

func TestReflect(t *testing.T) {
	u := new(demo.User)
	u.SetPhone("12345678901")
	u.SetUserName("Bob")
	u.SetPassword("11111")
	u.SetAddress("中国", "苏州", "平江")

	movie := demo.Movie{
		"123242342",
		2016,
		true,
		[]string{},
	}

	t.Run("", func(t *testing.T) {
		t.Helper()
		/*
		   Type相关操作:Type相关主要用于获取一个(接口)变量的类型信息，而不是实例信息(实例信息通过Value来获取)！！！
		*/

		//1、获取变量的Type
		typeOfu := reflect.TypeOf(u)
		fmt.Printf("u's type = %v\n", typeOfu) //u's type = *demo.User
		//
		type2Movie := reflect.TypeOf(movie)
		fmt.Printf("movie's type = %v\n", type2Movie) //movie's type = demo.Movie

		//2、获取变量的基础类型(也就是底层类型)；
		//注意这里要分两种情况：
		//case1：如果变量是指针类型比如u，则必须再次通过Elem()方法获取其type(指针类型)的基础类型(底层类型)demo.user，这样我们才可以获取demo.user这个类型的所有字段，因为字段定义在demp.user这个结构体内而不是*demo.user这个指针类型内
		baseType2User := typeOfu.Elem()
		fmt.Printf("baseTypeOf u = %+v\n", baseType2User) //baseTypeOf u = demo.User
		//case2:如果变量不是指针类型，比如像movie变量。则可以直接通过TypeOf()函数获取其基础类型即可，不能再使用Type.Elem()函数再取获取其基础变量，会触发painc exception,因为此时Type已经是基础类型了。

		//3、获取变量对应类型的字段信息(这里以结构体举例，且注意只有获取结构体变量的基础类型(如果变量是指针类型的话我们需要先利用Type.Elem()获得其基础类型后我们才能遍历它的字段!!!))
		fmt.Printf("变量u对应的类型共有 %d 个字段\n", baseType2User.NumField())
		//case1:获取所有字段
		for i := 0; i < baseType2User.NumField(); i++ {
			field := baseType2User.Field(i)
			fmt.Printf("这是变量u对应的类型的第 %d 个字段: %+v\n", i+1, field) //这是变量u对应的类型的第 1 个成员 {Name:phone PkgPath:_/E_/notebookforgo/TestDemo/demo Type:string Tag: Offset:0 Index:[0] Anonymous:false}：
			// 获取struct的tag信息
			lName, nameIsExisted := field.Tag.Lookup("lowerName")
			tagType, typeIsexisted := field.Tag.Lookup("type")
			if nameIsExisted && typeIsexisted {
				fmt.Printf("这是变量u对应的类型的字段 %s 的标签tag的field为: %+v, type为 %+v \n", field.Name, lName, tagType)
			}

			if field.Anonymous { //如果是匿名结构体成员，继续遍历输出它的所有字段
				fmt.Printf("第 %d 个field 是匿名的：%v,其type为：%+v\n", i+1, field.Anonymous, field.Type)
				for j := 0; j < field.Type.NumField(); j++ { //注意这里不能在使用Type.Elem().NumField,因为field.Type已经是filed对应的基础类型了，其没有更底层的类型了，所以不能再调用Elem()来获取你底层(基础)类型
					fmt.Printf("这是匿名变量%s对应的类型的第 %d 个字段: %+v：\n", field.Name, j+1, field.Type.Field(j))
				}
			}
		}
		//case2：根据字段名来获取(下例表明，使用反射一样可以访问非导出成员)
		fieldByName, ok := baseType2User.FieldByName("phone")
		if ok {
			fmt.Printf("获取成功，变量u的 %s 字段的信息为： %v \n", "phone", fieldByName)
		} else {
			fmt.Println("获取失败，变量u没有该字段")
		}

		//4.获取变量对应类型(一般是结构体类型)的方法,注意这里，一样区分基础类型和指针类型
		//case1:此时因为baseType2User是一个demo.User类型，所以它所拥有的方法是那些方法的接收参数为基础类型：demo.User的而不是指针类型：*demo.User的
		fmt.Printf("变量u对应的实例中方法的接收参数(receiver)为基础类型共有 %d 个方法\n", baseType2User.NumMethod())
		for i := 0; i < baseType2User.NumMethod(); i++ {
			method := baseType2User.Method(i)
			fmt.Printf("这是变量u对应的类型的第 %d 个方法: %+v：\n", i+1, method)

		}
		fmt.Printf("变量u对变量u对应的实例中方法的接收参数(receiver)为指针类型共有 %d 个方法\n", typeOfu.NumMethod())
		for i := 0; i < typeOfu.NumMethod(); i++ {
			method := typeOfu.Method(i)
			fmt.Printf("这是变量u对应的类型的第 %d 个方法: %+v：\n", i+1, method)

		}
		// case2:根据指定方法名来获取，
		methodByName, ok := typeOfu.MethodByName("SetName")
		if ok {
			fmt.Printf("获取成功，变量u的 %s 方法的信息为： %v \n", "SetName", methodByName)
		} else {
			fmt.Println("获取失败，变量u没有该方法")
		}

	})

	t.Run("", func(t *testing.T) {
		/*
			Vaule：用于实现对象实例的读写，且因为在go中(接口)变量传递时是值拷贝，所以如果想要实现利用反射来修改变量的值，则需要传入的是该变量的值，如果只想实现读的话，可以进行值传递！！！
		*/
		// case1:获取并修改基础类型变量的值
		i := 10
		s := "abc"

		//获取值：则可以直接进行值传递
		valueOfI := reflect.ValueOf(i)
		valueOfS := reflect.ValueOf(s)
		fmt.Printf("变量i的值是 %+v , canSet?：%v \n", valueOfI, valueOfI.CanSet()) //变量i的值是 10 , canSet：false
		fmt.Printf("变量s的值是 %+v , canSet?：%v \n", valueOfS, valueOfS.CanSet())

		//设置值的话，必须进行指针传递
		ptrOfI := reflect.ValueOf(&i) //进行地址传递，但这里虽然知道该变量的地址了，但因为接口变量存储的指针本身是无法进行寻址操作的，所以不能直接进行设置操作！！！，所以还得调用Elem()来获取指针所指向的对象
		ptrOfS := reflect.ValueOf(&s)
		valueOfI2 := ptrOfI.Elem()
		valueOfS2 := ptrOfS.Elem()
		fmt.Printf("变量i的值是 %+v , canSet?：%v \n", valueOfI2, valueOfI2.CanSet()) //变量i的值是 10 , canSet：true
		fmt.Printf("变量s的值是 %+v , canSet?：%v \n", valueOfS2, valueOfS2.CanSet())

		if valueOfI2.CanSet() {
			valueOfI2.SetInt(100)
			fmt.Printf("修改后变量i的值是 %+v \n", valueOfI2)
		}
		if valueOfS2.CanSet() {
			//valueOfS2.SetInt(100)//panic: reflect: call of reflect.Value.SetInt on string Value
			valueOfS2.SetString("def")
			fmt.Printf("修改后变量i的值是 %+v \n", valueOfS)
		}

		// case2:获取并修改复合数据类型变量的值
		slices := []int{1, 2, 3, 4}
		//跟上面一样，即使传递的是地址，也必须通过Elem()获取对应的对象，才能进行值设置
		valueOfSlices := reflect.ValueOf(&slices)
		fmt.Printf("变量slices的值是 %+v, canSet?：%v \n", valueOfSlices, valueOfSlices.CanSet())
		m := map[string]int{"Bob": 10}
		valueOfMap := reflect.ValueOf(&m).Elem()
		fmt.Printf("变量m的值是 %+v, canSet?：%v \n", valueOfMap, valueOfMap.CanSet())
		if valueOfMap.CanSet() {
			valueOfMap.SetMapIndex(reflect.ValueOf("scott"), reflect.ValueOf(30))
			fmt.Printf("修改后的map为 %+v \n", valueOfMap)
		}

		// case3；访问结构体实例的成员
		vauleOfUser := reflect.ValueOf(u)
		vauleOfMovie := reflect.ValueOf(movie)
		fmt.Printf("vauleOfUser: %+v ,type:%+v, canSet? %v \n", vauleOfUser, vauleOfUser.Type(), vauleOfUser.CanSet())     //实际上获得的只是一个地址，vauleOfUser: &{phone:12345678901 username:Bob password:11111 address1:{country:中国 city:苏州 street:平江} Address:{country:中国 city:苏州 street:平江}}
		fmt.Printf("vauleOfMovie: %+v ,type:%+v, canSet? %v \n", vauleOfMovie, vauleOfMovie.Type(), vauleOfMovie.CanSet()) //获取的是所指向对象的一个拷贝，所以无法直接被修改：vauleOfMovie: {Title:123242342 Year:2016 Color:true Actors:[]} ,canSet? false
		fmt.Printf("vauleOfMovie的 %s 字段值为 %v\n", "Title", vauleOfMovie.FieldByName("Title"))

		//case4；修改结构体实例的成员，跟上面一样，必须借助Elem()来获取指针原本所指向的对象，才能进行修改操作
		instanceOfUser := vauleOfUser.Elem()
		fmt.Printf("instanceOfUser: %+v canSet? %v \n", instanceOfUser, instanceOfUser.CanSet()) //已经获取到了u这个变量所指向的User实例：instanceOfUser: {phone:12345678901 username:Bob password:11111 address1:{country:中国 city:苏州 street:平江} Address:{country:中国 city:苏州 street:平江}}
		if instanceOfUser.FieldByName("phone").CanSet() {                                        // 无法修改非导出成员！！！
			instanceOfUser.FieldByName("phone").SetString("123456")
			fmt.Printf("vauleOfUser: %+v \n", vauleOfUser)
		} else {
			fmt.Printf("instanceOfUser 的 %s 字段无法被设置\n", "phone")
		}
		instanceOfMovie := reflect.ValueOf(&movie).Elem()
		fmt.Printf("instanceOfMovie: %+v canSet? %v \n", instanceOfMovie, instanceOfMovie.CanSet()) //已经获取到了u这个变量所指向的User实例：instanceOfUser: {phone:12345678901 username:Bob password:11111 address1:{country:中国 city:苏州 street:平江} Address:{country:中国 city:苏州 street:平江}}

		if instanceOfMovie.FieldByName("Title").CanSet() { //导出成员可以被修改！！！
			instanceOfMovie.FieldByName("Title").SetString("Tile01")
			fmt.Printf("修改后的instanceOfMovie: %+v \n", instanceOfMovie)
		} else {
			fmt.Printf("instanceOfMovie 的 %s 字段无法被设置\n", "Title")
		}
		//调用方法,当方法不需要传参时，可以传入一个空的切片，或者直接传入nil
		//method := instanceOfUser.MethodByName("GetUserName")
		//returnList := method.Call([]reflect.Value{}) // 这里触发一个panic exception: reflect: call of reflect.Value.Call on zero Value [recovered],因为instanceOfUser是demo.user类型，它只能访问那些被声明为(demo.User)的导出方法
		method := instanceOfUser.MethodByName("GetName")
		returnList := method.Call(nil)
		fmt.Printf("返回值列表为%+v\n", returnList) // 返回值列表为[Bob]

		method2 := vauleOfUser.MethodByName("GetUserName") // vauleOfUser是 *demo.User类型，所以它可以访问那些接收参数(receiver)为基础类型：demo.user以及指针类型：*demo.user的导出方法
		returnList2 := method2.Call(nil)
		fmt.Printf("返回值列表为%+v\n", returnList2) // 返回值列表为[Bob]

	})
}
