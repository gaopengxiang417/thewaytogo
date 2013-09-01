package main

import (
	"fmt"
	"io"
	"log"
	"runtime"
	"strings"
	"time"
)

var num = 10
var num2, num3 int

const LIMIT = 41

var fibs [LIMIT]uint64

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

	//我们不关心第二个返回值
	first, _, third := threereturn()
	fmt.Printf("the first : %d , the third : %f\n", first, third)

	min, max := minmax(43, 44)
	fmt.Printf("the min :%d , and the max :%d\n", min, max)

	tmp := 0
	c := &tmp
	multiply(44, 3, *c)

	fmt.Println("the result value :", *c)

	//首先传入的是一个不定常的多个数据
	minx := minTeset(3, 23, 4, 56, 7, 0)
	fmt.Printf("the min value is : %d\n", minx)
	//然后传入的是一个数组
	arr := []int{3, 43, 43, 6, 78, 32, 2}
	miny := minTeset(arr...)
	fmt.Printf("the min array is :%d\n", miny)

	printEachLine(2, 323, 43, 55)
	printEachLine(arr...)

	//defer 的使用
	deferFirst()
	deferLoation()
	deferorder()
	dbOperation()

	//通过使用defer来实现追踪的目的
	b()

	returndefer("go")
	buildinfunction()

	for i := 0; i < 10; i++ {
		fmt.Printf("%d\t", fabonacci(i))
	}

	fabonacciTest(10)

	fmt.Println()

	fmt.Printf("%d is odd :%v\t", 45, isEven(45))
	fmt.Printf("%d is even :%v\t", 44, isEven(44))

	fmt.Println()
	//计算内存办的fabonicci函数
	test_fib()

	for i := 0; i < 10; i++ {
		index, value := fabonaccitwo(i)
		fmt.Printf("%d : %d\t", index, value)
	}

	printtre(1)

	for i := 0; i < 10; i++ {
		fmt.Printf("the factorial %d:%d\n", i, factorial(i))
	}

	//函数的传递使用
	AddTest(4, addtest)

	//闭包的使用
	//将函数赋值给一个变量,通过变量来调用
	ff := func(a, b int) int { return a + b }
	fmt.Printf("%d \n", ff(3, 4))

	//直接调用
	func(a, b int) int {
		fmt.Println(a + b)
		return a + b
	}(4, 5)
	fmt.Println()

	calltest()

	fv := func() { fmt.Println("helloworld") }
	fmt.Printf("type is : %T,value is %v", fv, fv)

	//首先获取一个函数
	onetest := addOne()
	fmt.Printf("the value is %d , type is %T\n", onetest(4), addOne)
	twotest := addTwo(2)
	fmt.Printf("the value is %d,type is %T\n", twotest(4), twotest)

	//
	suffixone := addsuffix("first")
	fmt.Printf("%p\n", suffixone(".suffix"))

	//调试方式
	where := func() {
		pc, file, line, _ := runtime.Caller(1)
		log.Printf("%d , %s:%d", pc, file, line)
	}
	where()

	//
	where()

	log.SetFlags(log.Llongfile)
	log.Print("")

}
func test_fib() {
	start := time.Now()
	fmt.Println(start)
	for i := 0; i < LIMIT; i++ {
		result := fabonicci_memory(i)
		fmt.Printf("fibnaicci(%d) is :%d\n", i, result)
	}
	end := time.Now()
	eclipse := end.Sub(start)
	fmt.Printf("the eclipse time is %v\n", eclipse.Nanoseconds())
}

//对于faboniacci函数的计算现在开始从前往后计算,并且将计算结果放到内存中,节省时间
func fabonicci_memory(i int) (res uint64) {
	if fibs[i] != 0 {
		res = fibs[i]
		return
	}
	if i <= 1 {
		res = 1
	} else {
		res = fabonicci_memory(i-1) + fabonicci_memory(i-2)
	}
	fibs[i] = res
	return
}

//给文件添加后缀的方法
func addsuffix(filename string) func(s string) string {
	return func(s string) string {
		if !strings.HasSuffix(filename, s) {
			return filename + s
		}
		return filename
	}
}

//函数返回函数的使用,也就是closure
func addOne() func(a int) int {
	return func(a int) int {
		return a + 3
	}
}

func addTwo(a int) func(b int) int {
	return func(b int) int {
		return a + b
	}
}

//匿名函数的使用
func calltest() {
	for i := 0; i < 10; i++ {
		g := func(i int) { fmt.Printf("%d", i) }
		g(i)
		fmt.Printf(" -g is the type of %T , has the value %v\n", g, g)
	}
}

//函数的传递
func addtest(a, b int) {
	fmt.Printf("the sum of %d and %d is %d\n", a, b, a+b)
}

func AddTest(num int, f func(int, int)) {
	addtest(num, 5)
}

//使用递归来解决阶乘的问题
func factorial(num int) int {
	if num <= 1 {
		return 1
	} else {
		return num * factorial(num-1)
	}
}

//使用递归打印从10~1的数字
func printtre(num int) {
	if num > 10 {
		return
	}
	defer fmt.Printf("%d\t", num)
	printtre(num + 1)
}

//计算fabonacci函数的返回两个值,
func fabonaccitwo(a int) (index, num int) {
	index = a
	if a <= 0 {
		num = 1
	} else {
		_, value := fabonaccitwo(a - 1)
		_, value1 := fabonaccitwo(a - 2)
		num = value + value1
	}
	return
}
func isEven(num int) bool {
	if num == 0 {
		return true
	}
	return isOdd(relConv(num) - 1)
}

func isOdd(num int) bool {
	if num == 0 {
		return false
	}
	return isEven(relConv(num) - 1)
}

//不同函数的相互递归调用
func relConv(num int) int {
	if num < 0 {
		return -num
	}
	return num
}

//给定一个数字,直接数据其所有的fabonacci序列
func fabonacciTest(a int) int {
	var total int = 0
	if a <= 1 {
		fmt.Printf("%d\t", 1)
		total = 1
	} else {
		temp := fabonacciTest(a-1) + fabonacciTest(a-2)
		fmt.Printf("%d\t", temp)
		total += temp
	}

	return total
}

//递归函数的使用
func fabonacci(a int) (re int) {
	if a <= 1 {
		re = 1
	} else {
		re = fabonacci(a-1) + fabonacci(a-2)
	}
	return
}

func buildinfunction() {
	fmt.Println("len :", len("闪电似的"))
	fmt.Println("cap :", cap([]int{3, 43, 5}))
}

//其实defer可以获取到赶回的值
func returndefer(s string) (a int, e error) {
	defer func() {
		log.Printf("func1(%q) = %d , %v\n", s, a, e)
	}()
	return 5, io.EOF
}
func b() {
	startTrace("b")
	defer endTrace("b")
	fmt.Println("in b")
	a()
}

func a() {
	startTrace("a")
	defer endTrace("a")
	fmt.Println("in a")
}

//可以使用defer关键字来实现追踪的目的
func startTrace(s string) {
	fmt.Println("start enter:", s)
}

func endTrace(s string) {
	fmt.Println("exit :", s)
}

//模拟数据库的操作，打开使用和关闭连接
func dbOperation() {
	connecttoDB()
	defer closeConnection()

	fmt.Println("do some db operations....")
	fmt.Println("oh ,my god ,something error occur...")
	fmt.Println("redeay to stop the db....")
}
func connecttoDB() {
	fmt.Println("connect to db....", "ok")
}
func closeConnection() {
	fmt.Println("ok ,finally close the connection...")
}

//如果一个代码块里面有多个defer语句，那么这些defer语句的执行是类似与java的FILO
func deferorder() {
	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
	}
}

//defer的位置问题
func deferLoation() {
	i := 0
	defer fmt.Println("current defer value :", i)
	i++
	fmt.Println("before defer value :", i)
}

//defer语句的使用，类似与java的finally语句
func deferFirst() {
	fmt.Println("defer first line is start.....")
	defer defermiddle()
	fmt.Println("defer last line is called....")
}

func defermiddle() {
	fmt.Println("defer middle is called....")
}
func printEachLine(a ...int) {
	if len(a) == 0 {
		fmt.Println("empty line....")
		return
	}

	for _, value := range a {
		fmt.Println("the value :", value)
	}

}

//函数内部调用函数对于可变参数的传递有两种
func variableFirst(a ...int) {
	variableSecond(a...)
	variableThird(a)
}
func variableSecond(b ...int) {

}
func variableThird(a []int) {

}

//求可变参数列表的最小值
func minTeset(a ...int) int {
	if len(a) == 0 {
		return 0
	}
	min := a[0]
	for _, value := range a {
		if value < min {
			min = value
		}
	}
	return min
}

//修改引用类型的值，避免对象的拷贝
func multiTest(a, b int, c *int) {
	*c = a * b
}

//计算最小值和最大值的函数
func minmax(a, b int) (min, max int) {
	if a < b {
		min = a
		max = b
	} else {
		min = b
		max = a
	}
	return
}

//说明的是如果不关心返回值可以用_来接受
func threereturn() (int, int, float32) {
	return 3, 5, 5.4
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
