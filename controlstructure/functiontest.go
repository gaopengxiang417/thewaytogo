package main

import (
	"fmt"
)

var num = 10
var num2, num3 int

func main() {

	fmt.Println("in main before calling greeting")
	greeting()
	fmt.Println("in main after calling greeting")

	//函数可以赋值给变量,如果两个变量引用同一个函数,那么他们可以比较
	/*f1 := binOp
	f2 := binOp

	fmt.Println(f1 == f2)*/

	//返回不需要变量名的函数
	fmt.Printf("10 * 10 * 5 = %d\n", multiply(10, 10, 5))

	//下面测试不同两种函数类型的区别
	a := make([]int, 10)
	fmt.Println("before location:", &a)
	firstlocation(&a)
	secondlocation(a)

	//测试返回多个值的两种类型,一种是传入的时候有名称,另一种没有名称
	num2, num3 = num2x3x(num)
	printnum()
	num2, num3 = num2x3x_1(num)
	printnum()
}
func printnum() {
	fmt.Printf("num=%d,num2=%d,num3=%d\n", num, num2, num3)
}

func num2x3x(num int) (int, int) {
	return 2 * num, 3 * num
}

//如果传入的是有参数名称,那么返回的时候不需要主动待参数,但是必须要有return语句
func num2x3x_1(num int) (num2, num3 int) {
	num2 = 2 * num
	num3 = 3 * num
	return
}

//掉一个接收一个地址
func firstlocation(a *[]int) {
	b := a
	fmt.Println("first a location :", &a)
	fmt.Println("first  b location:", &b)
}

func secondlocation(a []int) {
	b := &a
	fmt.Println("second a location:", a)
	fmt.Println("second b location:", &b)
}

func multiply(a, b, c int) int {
	return a * b * c
}

//返回值的问题,如果出现了return，那么多个分支必须有return
/*func (st *List) Notreturn() int {
	v := 0
	for ix := len(st) - 1; ix >= 0; ix-- {
		if v := st[ix]; v != 0 {
			st[ix] = 0
			return v
		}
	}
}*/

//定义一个没有函数体的函数
// type binOp func(int, int) int

//go中是不可以进行方法重载的
//所以下面的会得到编译器错误
/*func first(a int) {

}
*/
/*func first(a string) {
}*/

//一般的函数调用
func greeting() {
	fmt.Println("hello world")
}
