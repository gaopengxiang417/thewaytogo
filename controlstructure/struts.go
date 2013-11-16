package main

import (
	"fmt"
)

func main() {
	t := new(T)
	fmt.Println(t)
	fmt.Println(t.a)

	var s T
	s.a = 44
	s.b = "china"
	fmt.Println(s)

	//测试struct的使用
	var ss = new(S)
	ss.i1 = 43
	ss.f1 = 98.3232
	ss.s1 = "first line"

	fmt.Printf("the int value %d\n", ss.i1)
	fmt.Printf("the float value %f\n", ss.f1)
	fmt.Printf("the string value %s\n", ss.s1)
	fmt.Println(ss)
	fmt.Printf("%v", ss)

	//第二种初始化结构体的方法，这个方法和new关键字出来的对象一样的
	//他们都是指向对象的指针
	sss := &S{23, 54.22, "second"}
	fmt.Printf("the int value %d\n", sss.i1)
	fmt.Printf("the value %v\n", sss)

	//这个就是就是对象而不是指针
	var sst S = S{54, 65.22, "third"}
	fmt.Printf("the int value %d\n", sst.i1)
	fmt.Printf("the value %v\n", sst)

	//初始化的几种方式
	inter1 := Interval{3, 5}
	inter2 := Interval{start: 43, end: 565}
	inter3 := Interval{end: 54545}

	fmt.Printf("%v\n", inter1)
	fmt.Printf("%v\n", inter2)
	fmt.Printf("%v\n", inter3)
}

type Interval struct {
	start int
	end   int
}

type S struct {
	i1 int
	f1 float32
	s1 string
}

type T struct {
	a int32
	b string
}
