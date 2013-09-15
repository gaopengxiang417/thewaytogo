package main

import (
	"container/list"
	"fmt"
	"math"
	"math/big"
	"regexp"
	"strconv"
	"sync"
	"unsafe"
)

func main() {
	double_list()
	unsafe_size_test()
	regexp_test()

	lock_test()

	math_big_test()
}

//math 和 big包的使用
func math_big_test() {
	first := big.NewInt(math.MaxInt64)
	second := first
	third := big.NewInt(18232)
	fouth := big.NewInt(1)

	fmt.Println(fouth.Mul(first, second))
	fmt.Println(fouth.Add(first, third))
	fmt.Println(fouth.Div(first, second))

	//有理数的使用
	rationone := big.NewRat(math.MaxInt64, 232)
	rationtwo := big.NewRat(-1212, math.MaxInt64)
	rationthird := big.NewRat(100, 23)
	rationfouth := big.NewRat(1, 1)

	fmt.Println(rationone)
	fmt.Println(rationtwo)
	fmt.Println(rationthird)
	fmt.Println(rationfouth)

	fmt.Println(rationfouth.Add(rationone, rationtwo))
}

func lock_test() {
	info := new(Info)
	fmt.Println(info.str)
	update(info)
	fmt.Println(info.str)
}

//锁的使用
type Info struct {
	mu  sync.Mutex
	str string
}

func update(info *Info) {
	info.mu.Lock()

	info.str = "wangchen"

	info.mu.Unlock()
}

//使用正则表达式
func regexp_test() {
	var search = "John : 2323.233 walliim:98.3 , china:232.121443"
	pattern := "[0-9]+.[0-9]+"

	fun := func(s string) string {
		v, _ := strconv.ParseFloat(s, 32)
		return strconv.FormatFloat(v*2, 'f', 2, 32)
	}

	if ok, _ := regexp.Match(pattern, []byte(search)); ok {
		fmt.Println("founded")
	}

	regex, err := regexp.Compile(pattern)
	if err != nil {
		fmt.Println("error ocur...")
	}

	result := regex.FindAllString(search, 10)
	fmt.Println(result)

	fmt.Println(regex.ReplaceAllString(search, "####.####"))

	fmt.Println(regex.ReplaceAllStringFunc(search, fun))
}

//判断描述类型所占的字节多少
func unsafe_size_test() {
	i := 4
	fmt.Println(unsafe.Sizeof(i))
	var f float32
	fmt.Println(unsafe.Sizeof(f))
}

//对双向列表的添加和获取数据
func double_list() {
	ls := list.New()
	ls.PushFront(101)
	ls.PushFront(102)
	ls.PushFront(103)

	fmt.Println("length:", ls.Len())
	for ele := ls.Front(); ele != nil; ele = ele.Next() {
		fmt.Println(ele.Value)
	}

	//从结束到开始
	for ele := ls.Back(); ele != nil; ele = ele.Prev() {
		fmt.Println(ele.Value)
	}
}
