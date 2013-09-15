/*
	变量的使用和声明
*/
package main

import (
	"fmt"
	"os"
	"time"
)

var china = 44 //全局变量可以声明,但是不使用
//初始化函数是指go在调用的时候首先执行的函数
var abc string

func init() {
	fmt.Println("start...........")
}
func main() {
	//这里声明为指针
	var a, b *int

	fmt.Printf("%T---%v", a, a)
	fmt.Println()
	fmt.Printf("%T---%v", b, b)

	//变量的声明，多个一起声明
	var (
		x string
		y int
		z float32
	)
	fmt.Println()
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)

	//go中不容许不同类型的赋值
	var xx = 5
	fmt.Println(xx)

	// xx = false
	// fmt.Println(xx)

	//测试获取os
	var opersystem string = os.Getenv("GOOS")
	fmt.Printf("the os is %s\n", opersystem)

	var path string = os.Getenv("PATH")
	fmt.Printf("the path is %s\n", path)

	fmt.Println(time.Now())

	var (
		xxx = 15
		yyy = "this is a string test"
		zzz = true
		www = 42.3
	)

	fmt.Printf("%v---%T\n", xxx, xxx)
	fmt.Printf("%v---%T\n", yyy, yyy)
	fmt.Printf("%v---%T\n", zzz, zzz)
	fmt.Printf("%v---%T\n", www, www)

	//如果不想用编译器自动分配的类型,那么应该自己指定
	var mm int64 = 44
	fmt.Printf("%v----%T\n", mm, mm)

	//获取环境变量,这里的变量是在运行时进行赋值的
	home := os.Getenv("HOME")
	user := os.Getenv("USER")
	root := os.Getenv("GOROOT")

	fmt.Println(home)
	fmt.Println(user)
	fmt.Println(root)

	//基本数据类型是值类型，他们直接指向值
	//arrays 和 struct也是值类型
	var ww = 44
	var gg = ww

	fmt.Println("ww:", ww)
	fmt.Println("gg:", gg)

	gg = 45
	fmt.Println("ww:", ww)
	fmt.Println("gg:", gg)

	//slice赋值是引用的赋值
	var first = make([]int, 20)
	var second = first
	fmt.Printf("%v----%T\n", first, first)
	fmt.Printf("%v----%T\n", second, second)

	second[0] = 44
	fmt.Printf("%v----%T\n", first, first)
	fmt.Printf("%v----%T\n", second, second)

	//array的赋值是值的赋值
	var third = [20]int{}
	var four = third
	fmt.Printf("%v----%T\n", third, third)
	fmt.Printf("%v----%T\n", four, four)

	four[0] = 44
	fmt.Printf("%v----%T\n", third, third)
	var result = fmt.Sprintf("%v----%T\n", four, four)
	fmt.Println(result)

	//:=这个操作只能用于函数的内部变量声明,不能用于全局变量的声明
	abc = "G"
	fmt.Println(abc)
	f1()
}
func f1() {
	abc := "o"
	fmt.Println(abc)
	f2()
}

func f2() {
	fmt.Println(abc)
}
