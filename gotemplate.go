package main

//首先应该进行导入
import (
	"fmt"
)

//然后应该定义常量,全局变量
const c = "c"

var v int = 4

type T struct{}
type TT int

//然后定义初始化的一些方法
func init() {
	//init
}

//然后定义main函数
func main() {
	var a int
	Func1()
	fmt.Println(a)

//类型的转化
	x := 4
	y := int(x)
	fmt.Println(y)

	var xy TT = 44
	yy := int(xy)
	fmt.Println(xy)
	fmt.Println(yy)

	//常量的定义
	const abc int = 46
	fmt.Println(abc)

	//常量不能改变他的值
	// abc = 77
	// fmt.Println(abc)

	//常量只能用于基本类型不能用于自定义类型
	// const tt T = 44

	//常量可以用于无线的数字
	const Ln2 = 0.2222222222222222222222222
		const Log2E = 1/Ln2
		const Billion = 1e9
		const hardEight = (1 << 100) >> 97

		fmt.Println(Ln2)
		fmt.Println(Log2E)
		fmt.Println(Billion)
		fmt.Println(hardEight)
		//常量的批量定义
		const (
			Monday , Tuesday,Wennesday = 1 , 2 , 3
		)
		fmt.Println(Monday,Tuesday,Wennesday)

		//枚举的类型
		const (
			a1 = 0 
			b = 1 
			c = 2
		)
		fmt.Println(a1)
		fmt.Println(b)
		fmt.Println(c)

		const (
			x1 = iota
			x2
			x3
			x4
		)
		fmt.Println(x1,x2,x3,x4)
}


//最后定义方法
func (t T) Method1() {
	//todo
}
func Func1() {
	//todo
}
