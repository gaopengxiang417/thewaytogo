package main

import (
	"fmt"
)

func main() {
	fmt.Println("hello world")

	//声明一个数组
	var first [5]int
	fmt.Println("first length :", len(first))
	//打印
	for index, value := range first {
		fmt.Println("first[", index, "]=", value)
	}

	//打印
	forandprint()
	//打印string数组
	forstringarray()
	arraycopyvalue()

	//
	var a [4]int
	valuearraypass(a)
	referencearraypass(&a)
}

//函数中数组的值传递和引用传递
//引用传递
func referencearraypass(a *[4]int) {
	fmt.Println(a)
}

//值传递
func valuearraypass(a [4]int) {
	fmt.Println(a)
}

//数组是值类型,所以将一个数组赋值给另一个,会拷贝一个
func arraycopyvalue() {
	var array1 [5]int
	array2 := array1

	array2[0] = 4
	array1[0] = 5
	for _, value := range array1 {
		fmt.Printf("%d \t", value)
	}

	fmt.Println()

	for _, value2 := range array2 {
		fmt.Printf("%d \t", value2)
	}
}

//循环string数组
func forstringarray() {
	strs := [...]string{"a", "b", "c", "d"}
	for index, value := range strs {
		fmt.Println("index:", index, "value", value)
	}
}

//打印for循环数组
func forandprint() {
	//首先初始化一个
	var first [6]int
	for i := 0; i < len(first); i++ {
		first[i] = i * 2
	}
	//打印
	for i := 0; i < len(first); i++ {
		fmt.Printf("the index %d , value is %d\n", i, first[i])
	}
}
