package main

import (
	"fmt"
	"strings"
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

	//测试不同的初始化方法在内存中的布局
	//1.
	var p1 Person
	p1.firstname = "p1 firstname"
	p1.lastname = "p1 lastname"
	upPerson(&p1)
	fmt.Printf("the person is %s %s\n", p1.firstname, p1.lastname)

	//2
	p2 := new(Person)
	p2.firstname = "p2 firstname"
	p2.lastname = "p2 lastname"
	upPerson(p2)
	fmt.Printf("the person is %s %s\n", p2.firstname, p2.lastname)

	//3
	p3 := &Person{"p3 firstname", "p3 lastanem"}
	upPerson(p3)
	fmt.Printf("the person is %s %s\n", p3.firstname, p3.lastname)

	//转化
	a := number{53.3434}
	b := nr{42.2323}

	var c = b
	fmt.Println(a, b, c)

	// var d number = b
	var d number = a
	fmt.Println(d)

	var e = number(b)
	fmt.Print(e)
}

type nr number

type number struct {
	f float32
}

func upPerson(person *Person) {
	person.firstname = strings.ToUpper(person.firstname)
	person.lastname = strings.ToUpper(person.lastname)
}

type Person struct {
	firstname string
	lastname  string
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
