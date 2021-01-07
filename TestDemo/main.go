package main

/*
@Author:David Ma
@Content: 多对多的生产者消费者模型
@Date:2020-12-21
*/

import (
	"./util"
	"fmt"
	"math/rand"
	"runtime"
	"strconv"
	"strings"
	"sync"
	"time"
)

var cond = sync.NewCond(&sync.Mutex{}) //创建一个全局的条件变量
const BUFFERSIZE = 10

func GoID() int { //go中没有对外暴露获取goroutineID的api
	var buf [64]byte
	n := runtime.Stack(buf[:], false)
	idField := strings.Fields(strings.TrimPrefix(string(buf[:n]), "goroutine "))[0]
	id, err := strconv.Atoi(idField)
	if err != nil {
		panic(fmt.Sprintf("cannot get goroutine id: %v", err))
	}
	return id
}

func Producer(ch chan<- int) {
	for i := 0; i <= 5; i++ {
		//因为条件变量一定是与互斥锁绑定使用，go里面干脆直接将这两个绑定在一块了
		cond.L.Lock() //注意这里，一定不能把lock()操作放在for循环外，否则报错fatal error: sync: unlock of unlocked mutex，
		// 原因也很明显，当lock操作放在for循环外时，Producer协程只会在第一次访问时执行lock()，(这明显不合理，每个协程应该在对临界资源channel做每次操作时都要加锁!!!),而我们却在for循环体内执行了多次unlock操作
		for len(ch) >= BUFFERSIZE { //使用for循环代替if来进行条件判定，原因为上下文切换
			cond.Wait()
		}
		data := rand.Int()
		ch <- data
		fmt.Printf("Producer '%d'生产了一个数据 '%d'\n", GoID(), data)
		cond.L.Unlock()
		cond.Signal() //发送一个信号唤醒被cond所阻塞的Consumer协程，相当于Java的notify()
		time.Sleep(time.Second * 3)
	}

}

func Consumer(ch <-chan int) {
	for i := 0; i < 15; i++ {
		cond.L.Lock()
		for len(ch) <= 0 { //使用for循环代替if来进行条件判定，原因为上下文切换
			cond.Wait()
		}
		fmt.Printf("Consumer '%d'消费了一个数据 '%d'\n", GoID(), <-ch)
		cond.L.Unlock()
		cond.Signal() //发送一个信号唤醒被cond所阻塞的Producer协程，相当于Java的notify()
		time.Sleep(time.Second)
	}
}

func main() {
	ipAddr := util.IPV4()
	fmt.Printf("ipAddr:%s\n", ipAddr)
	ip := strings.Replace(ipAddr, ".", "", -1)
	fmt.Printf("after replace ipAddr:%s\n", ip)
	ipInt, _ := strconv.ParseInt(ip, 10, 64)
	fmt.Printf(" ipAddrOfInt:%d\n", ipInt)
	//idSlices := make([]int64, 0, 4000000)
	//myChannel := make(chan int, 1000)
	//locker := sync.Mutex{}
	//fmt.Println("start", time.Now().Format("13:04:05"))
	//
	//for i:=1;i<=1000;i++{
	//	go func() {
	//		    machineID := 192168000000+i
	//		    for j:=1;j<=4000;j++{
	//				demo.SetMachineId(int64(machineID))
	//				id := demo.GetSnowflakeId()
	//				locker.Lock()
	//				idSlices = append(idSlices, id)
	//				locker.Unlock()
	//			}
	//		    myChannel <- i
	//	}()
	//}
	//for i:=1;i<=1000;i++{
	//	<- myChannel
	//}
	//fmt.Println("end", time.Now().Format("13:04:05"))
	//fmt.Printf("去重前的ID数量：%d\n",len(idSlices))
	//result := demo.Duplicate(idSlices)
	//fmt.Printf("去重后的ID数量：%d\n",len(result))
	//interfaces, err := net.Interfaces()
	//if err != nil {
	//	panic(err)
	//}
	//inter := interfaces[0]//例如笔记本电脑可以连接无线网和有线连接，这就是两个不同的MAC地址,所以一般取第一个MAC地址作为唯一标识。
	//macAddr := inter.HardwareAddr.String() //获取本机MAC地址
	//fmt.Printf("origin mac:%v\n",macAddr)
	//datetime := "2015-01-01 00:00:00"  //待转化为时间戳的字符串
	//
	////日期转化为时间戳
	//timeLayout := "2006-01-02 15:04:05"  //转化所需模板
	//loc, _ := time.LoadLocation("Local")    //获取时区
	//tmp, _ := time.ParseInLocation(timeLayout, datetime, loc)
	//timestamp := tmp.Unix()    //转化为时间戳 类型是int64
	//fmt.Println(timestamp)
	//
	////时间戳转化为日期
	//datetime = time.Unix(timestamp, 0).Format(timeLayout)
	//fmt.Println(datetime)
	//movie := new(demo.Movie)
	//fmt.Printf("movie:%+v movie==nil?%+v\n",movie,movie==nil)
	//var m2 *demo.Movie
	//fmt.Printf("m2:%+v m2==nil?%+v\n", m2, m2==nil)
	//rand.Seed(time.Now().UnixNano())
	//ch := make(chan int, BUFFERSIZE)
	//for i:=1; i<=5; i++{
	//	go Producer(ch)
	//}
	//for i:=1; i<=2; i++{
	//	go Consumer(ch)
	//}
	//time.Sleep(time.Second * 30)
}
