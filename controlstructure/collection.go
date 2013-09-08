package main

import (
	"fmt"
)

var fabs [50]uint64

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

	//计算fab函数
	for i := 0; i < 50; i++ {
		fab_naic_array(i)
	}

	fmt.Println()

	for _, value := range fabs {
		fmt.Printf("%d\t", value)
	}

	fmt.Println()
	array_instance()

	for i := 0; i < 5; i++ {
		listtest(&[]int{i, i * i, i * i * i})
	}

	multi_array()

	//
	var sumarray = [3]float64{33.4, 22.5, 434.21321}
	result := sum_array(&sumarray)
	fmt.Println("result:", result)
}

//数组的引用传递
func sum_array(a *[3]float64) (sum float64) {
	for _, value := range a {
		sum += value
	}
	return
}

//多维数组的使用
func multi_array() {
	var a [10][5]int
	fmt.Println("multi length:", len(a))
	fmt.Println("multi second length:", len(a[0]))

	//输出
	for i := 0; i < len(a); i++ {
		for j := 0; j < len(a[i]); j++ {
			fmt.Printf("a[%d][%d]=%d\n", i, j, a[i][j])
		}
	}

}

//数组的引用传递
func listtest(a *[]int) {
	fmt.Println(len(*a))
	fmt.Println(a)
}

//数组的初始化
func array_instance() {
	var arrindex = [5]int{1, 2, 3, 4, 5}
	var arr_score = [...]int{1, 2, 3, 4, 5}
	var arr_string = [5]string{3: "third", 4: "fouth"}
	var errorarray = []string{"22"}
	fmt.Println("error", len(errorarray))

	fmt.Println("first length", len(arrindex))
	fmt.Println("first capcity", cap(arrindex))

	fmt.Println("second length:", len(arr_score))
	fmt.Println("third length:", len(arr_string))

	fmt.Println("first == second?:", arrindex == arr_score)
}

//计算fabionicc函数的值fun
func fab_naic_array(a int) (ret uint64) {
	if fabs[a] != 0 {
		ret = fabs[a]
		return
	}

	if a <= 1 {
		ret = 1
	} else {
		ret = fab_naic_array(a-1) + fab_naic_array(a-2)
	}
	fabs[a] = ret
	return
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
