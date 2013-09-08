package main

import (
	"bytes"
	"fmt"
	"sort"
	"unicode/utf8"
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

	//slice
	slice_test()

	////函数中传递slice
	var origin_array = [10]int{2, 3, 4, 6}
	fmt.Println(sum_slice(origin_array[:]))

	create_slice()
	fmt.Println()
	//
	slice_funcion()

	//
	concat_buffer()

	for_range_test()

	var sumslice = make([]float64, 10)
	fmt.Println(sum_for_float(sumslice))
	var sumarraysecond = []float64{23.4}
	fmt.Println(sum_for_float(sumarraysecond))

	reslice()

	var firstslice = []int{2, 4, 5}
	var secondslce = []int{5, 6, 7}
	var result_slice = copy_slice(firstslice, secondslce)
	for _, value := range result_slice {
		fmt.Printf("%d\t", value)
	}
	fmt.Println()
	fmt.Println(len(result_slice))
	fmt.Println(cap(result_slice))

	relarge_slice()

	var input = []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	ed := filter_slice(input, slice_iseven)
	for _, value := range ed {
		fmt.Printf("result is %d\n", value)
	}

	var a_slice = []int{2, 3, 4, 56}
	var b_slice = []int{6, 7, 8}
	insert_slice(a_slice, b_slice, 4)
	for _, value := range a_slice {
		fmt.Printf("result %d\t", value)
	}

	var llsds = []int{1, 2, 3, 4, 5, 6, 7, 8, 89, 9}
	remove_slice(llsds, 4, 7)

	//转换string为byte数组
	var str_slice = "love 中国 this is"
	c := []byte(str_slice)
	fmt.Println(len(c))
	for index, value := range c {
		fmt.Println(index, ":", value)
	}
	fmt.Println(utf8.RuneCountInString(str_slice))

	//截取一个字符串的子集,不过他的数字指定的是指节的位置
	sub_string := str_slice[7:19]
	fmt.Println(sub_string)

	string_replce()

	slice_sort()
}

//排序的算法
func slice_sort() {
	var int_slice = []int{2, 3, 543, 56, 767, 4, 54, 6, 7, 8}
	sort.Ints(int_slice)
	fmt.Println(sort.IntsAreSorted(int_slice))
	for _, value := range int_slice {
		fmt.Printf("%d\t", value)
	}
	//search
	var index = sort.SearchInts(int_slice, 767)
	if index == len(int_slice) {
		fmt.Println("not search")
	}
	fmt.Println(index)
}

//string的替换,string,替换必须转换为字节数组,替换,然后再转换回来
func string_replce() {
	var str = "this is a 中国"
	bytes := []byte(str)
	bytes[0] = 'H'
	s := string(bytes)
	fmt.Println(s)

}
func remove_slice(a []int, start, end int) {
	if start > end || start < 0 || end > cap(a) {
		// do nothing
		return
	}
	//首先应该拷贝一个数组出来,防止越界
	first := a[:start]
	second := a[end:]

	var result = make([]int, len(first)+len(second))
	var count int
	for index, value := range first {
		result[index] = value
		count++
	}
	for index, value := range second {
		result[index+count] = value
	}

	for _, value := range result {
		fmt.Printf("%d\t", value)
	}
}

//insert 一个slice到另一个的指定位置
func insert_slice(a []int, b []int, index int) int {
	if index > cap(a) || index < 0 {
		return 0
	}
	alength := len(a)
	acapcity := cap(a)
	blength := len(b)

	if alength+blength > acapcity {
		var temp = make([]int, (alength+blength)*2)
		//拷贝
		for i, value := range a {
			temp[i] = value
		}
		a = temp
	}
	//b拷贝到a
	if index >= alength {
		for index, value := range b {
			a[alength+index] = value
		}
	} else {
		ttt := a[index:]
		for i, value := range b {
			a[index+i] = value
		}
		for i, value := range ttt {
			fmt.Println("*****", blength, "***", i, "***", index, "***")
			a[blength+index] = value
		}
	}
	for _, value := range a {
		fmt.Printf(" %d\t", value)
	}
	return blength
}

//校验函数
func filter_slice(a []int, fn func(int) bool) []int {
	var result = make([]int, len(a))
	for index, value := range a {
		if fn(value) {
			result[index] = value
		}
	}
	return result
}

//是否是偶数
func slice_iseven(a int) bool {
	if a%2 == 0 {
		return true
	}
	return false
}

//relarge slice
func relarge_slice() {
	var slice = []int{3, 4, 6}
	var facotr = 8
	var result = make([]int, len(slice)*facotr)

	fmt.Println(len(result))
}

func copy_slice(from []int, to []int) []int {
	var length = len(from)
	var capcity = cap(from)
	if (len(to) + length) > capcity {
		var temp = make([]int, 2*(len(to)+length))
		for index, _ := range from {
			temp[index] = from[index]
		}
		from = temp
	}
	for index, value := range to {
		from[length+index] = value
	}
	return from
}

//reslice
func reslice() {
	var slice = make([]int, 0, 10)
	//填充
	for i := 0; i < cap(slice); i++ {
		slice = slice[0 : i+1]
		slice[i] = i * i
		fmt.Printf("slice index :%d , value is %d,lenth :%d , capcity %d\n", i, slice[i], len(slice), cap(slice))
	}
	fmt.Println()
	//print
	for index, value := range slice {
		fmt.Printf("index %d , value %d\n", index, value)
	}

	var ar = [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	var a = ar[5:7]
	fmt.Println(len(a))
	fmt.Println(cap(a))

	a = a[0:4]
	fmt.Println(len(a))
	fmt.Println(cap(a))

	a = a[0:0]
	fmt.Println(len(a))

}

//slice的计算总和
func sum_for_float(a []float64) (res float64) {
	for _, value := range a {
		res += value
	}
	return
}

//for range for slice
func for_range_test() {
	var arr = make([]int, 5)
	for i := 0; i < len(arr); i++ {
		arr[i] = i * i
	}

	//for range first
	for index, value := range arr {
		fmt.Println("index:", index, ",value:", value)
	}

	//for range second
	for _, value := range arr {
		fmt.Println(value)
	}
	//for range third
	for index := range arr {
		fmt.Println(index)
	}

	//not change for range
	for _, value := range arr {
		value *= 2
	}
	for _, value := range arr {
		fmt.Println(value)
	}

	//can change for range
	for index, value := range arr {
		arr[index] = value * 2
	}
	for _, value := range arr {
		fmt.Println(value)
	}
}

//bytes.Buffer的使用,用来连接多个字符串
func concat_buffer() {
	var str = []string{"this", " is a ", " very inter", "sting story"}
	var buffer bytes.Buffer
	for _, val := range str {
		buffer.WriteString(val)
	}
	fmt.Println(buffer.String())
}

//slice测试
func slice_funcion() {
	s := make([]byte, 5)
	fmt.Println("lenth :", len(s))
	fmt.Println("capcity :", cap(s))

	s = s[2:3]
	fmt.Println("length :", len(s))
	fmt.Println("capcity :", cap(s))

	ss := new([]int)
	fmt.Println(ss)

	s1 := []byte{'a', 'm', 'd', 'k'}
	s2 := s1[1:3]

	s2[1] = 'b'
	for _, value := range s1 {
		fmt.Printf("%c\t", value)
	}
}

//slice的创建
func create_slice() {
	var slice1 = make([]byte, 10)
	slice2 := make([]byte, 10)
	slice3 := make([]byte, 10, 10)
	slice4 := new([10]byte)[:]

	fmt.Printf("the type %T", slice1)
	fmt.Printf("the type %T", slice2)
	fmt.Printf("the type %T", slice3)
	fmt.Printf("the type %T", slice4)
}

//函数中传递slice
func sum_slice(a []int) (res int) {
	for _, value := range a {
		res += value
	}
	return
}

//slice的使用,slice底层是由数组构成,他是没有长度限制
func slice_test() {
	//首先声明一个数组
	var array [10]int
	//声明一个slice
	var slice1 = array[2:7]

	//初始化数组
	for index, _ := range array {
		array[index] = index
	}
	//打印数组
	for index, value := range array {
		fmt.Printf("array[%d]=%d\n", index, value)
	}
	//打印slice
	for index, value := range slice1 {
		fmt.Printf("slice[%d]=%d\n", index, value)
	}

	//长度
	fmt.Printf("the length of array is :%d\n", len(array))
	fmt.Printf("the length of slice is :%d\n", len(slice1))
	fmt.Printf("the capcity of slice is :%d\n", cap(slice1))

	//添加slice的capcity
	slice1 = slice1[0:4]
	fmt.Printf("after slice length:%d\n", len(slice1))
	fmt.Printf("after slice capcity:%d\n", cap(slice1))
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
