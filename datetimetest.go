package main

/*
	对于time和date的学习和使用

*/
import (
	"fmt"
	"time"
)

var week time.Duration

func main() {
	//获取当前的时间
	var now = time.Now()
	//打印
	fmt.Printf("%02d-%02d-%04d\n", now.Month(), now.Day(), now.Year())
	//格式化日期
	fmt.Println(now.Format(time.ANSIC))

	now = time.Now().UTC()
	fmt.Println(now)
	fmt.Println(time.Now())

	//加一个指定的日期
	//注意这里必须是纳秒级别的
	week = 60 * 60 * 24 * 7 * 1e9
	afterweek := now.Add(week)
	fmt.Println(afterweek)

	//不同的格式化形式
	//Mon, 02 Jan 2006 15:04:05 -0700
	//必须是日志日期进行格式化
	fmt.Println(now.Format(time.RFC822))
	fmt.Println(now.Format(time.ANSIC))
	fmt.Println(now.Format("02 Jan 2006 15:04:05"))
	fmt.Println("before:", now, "after:", now.Format("20060102"))

	//指针的逻辑 %p用来格式化指针
	var i1 = 5
	fmt.Printf("the integer %d,and it memory address is : %p\n", i1, &i1)

	var i *int
	fmt.Printf("i value is %p, type is %T,value is %v\n", i, i, i)

	i = &i1
	fmt.Printf("location is :%v\n", i)
	fmt.Println(*i)

	//改变指针的指向的值的时候,其余的指针指向同一个值的也会改变
	var str = "this is origin string"
	var pointer *string = &str
	fmt.Printf("origin string is %s\n", str)
	fmt.Printf("location is %p\n", pointer)

	//改变值
	*pointer = "after changed value"
	fmt.Printf("pointer after changed value is %s\n", *pointer)
	fmt.Printf("str after changed value is %s\n", str)

	//指针的嵌套使用
	var intt int = 4
	var pointer1 *int = &intt
	var pointer2 *int = pointer1
	var pointer3 *int = pointer2

	fmt.Printf("origin value is : %d\n", intt)
	fmt.Printf("pointer1 address is %p, value is %d\n", pointer1, *pointer1)
	fmt.Printf("pointer2 address is %p, value is %d\n", pointer2, *pointer2)
	fmt.Printf("pointer3 address is %p, value is %d\n", pointer3, *pointer3)
}
