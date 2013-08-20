package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
	"unicode"
	"unicode/utf8"
)

func main() {
	var s = "ab"
	var ss = "cd"

	fmt.Println(s + ss)

	//bool类型的使用
	var b bool = true
	fmt.Println(b)

	//go 不支持不同类型的比较，说白了就是不支持隐士的转化
	// var a int = 44
	// var c int64 = 66
	// fmt.Println(a == c)
	//go对逻辑运算的支持
	fmt.Println(true && false)
	fmt.Println(true || false)
	//go对逻辑运算也是支持的一样，从左右到，支持的是最左匹配原则
	//所以这里的后面的方法是不会执行的
	if true || isbool() {
		fmt.Println("true success")
	}
	//go对%t的用处
	fmt.Printf("%t bool", true)
	//int uint类型是依赖于平台的,但是int8,int6等是固定的

	//八进制的表示
	var octal = 077
	//十六进制的表示
	var hexi = 0xff
	fmt.Println("octal:", octal)
	fmt.Println("hexi:", hexi)

	//go 不支持隐士的赋值
	var first int32 = 55
	var second int64

	// second = first
	second = int64(first)
	fmt.Println(second)

	//%d 表示整数 , %x用来表示十六进制,%g表示浮点,%f也是浮点，%e科学计数法,
	//m.ng 表示保留整型部分是m位，小数部分是n为
	result, _ := intconver(55)
	fmt.Println(result)

	result2, _ := floatconver(99.8)
	fmt.Println(result2)

	//复数的类型
	var com = 4 + 8i
	fmt.Printf("%v is %T\n", com, com)
	//返回复数对应的实数和虚数部分
	fmt.Printf("%v is %T\n", real(com), real(com))
	fmt.Printf("%v is %T\n", imag(com), imag(com))

	//比特的运算
	//比特运算只能用于整型
	var s1 = 1
	var s2 = 3
	fmt.Printf("%b.....%v\n", s1&s2, s1&s2) //and
	fmt.Printf("%b.....%v\n", s1|s2, s1|s2) //or
	fmt.Printf("%b.....%v\n", s1^s2, s1^s2) //xor
	fmt.Printf("%b......%v\n", ^s1, ^s1)

	//移位的操作
	var s3 = s1 << 8
	fmt.Printf("%v......%T\n", s3, s3)

	var s4 = 8
	s4 >>= 2
	fmt.Printf("%v....%T\n", s4, s4)

	fmt.Printf("%v...%T\n", KB, KB)

	fmt.Printf("%v.....%T\n", Active, Active)
	fmt.Printf("%v.....%T\n", Send, Send)

	//逻辑运算符
	fmt.Printf("%v...%T\n", 2 == 1, 2 == 1)

	//算数运算
	//+-*/对于整型和浮点型都是支持的
	//下面的例子说明了只要有一个是浮点类型,就结果为浮点类型,也就是最大边界原则
	fmt.Printf("%v...%T\n", 8/3, 8/3)
	fmt.Printf("%v...%T\n", 8.0/3, 8.0/3)
	fmt.Printf("%v....%T\n", 8/3.0, 8/3.0)

	//+= -= *= /=
	var s6 = 90
	s6 += 90
	fmt.Printf("%v.....%T\n", s6, s6)
	//++ -- 只能用于操作数的后面,不能用于前面,并且支持整型和浮点型
	//只能用于statement ,不能用于表达式
	var s7 = 94.3
	s7++
	fmt.Printf("%v....%T\n", s7, s7)

	//随机数字的使用
	//首先打印随机的整数
	for i := 0; i < 10; i++ {
		fmt.Printf("%d / ", rand.Int())
	}

	//限定最大的上限
	for i := 0; i < 5; i++ {
		fmt.Printf("%d / ", rand.Intn(8))
	}

	fmt.Println()

	//生成浮点型的随机数
	rand.Seed(int64(time.Now().Nanosecond()))

	for i := 0; i < 10; i++ {
		fmt.Printf("%2.2f / ", rand.Float64()*100)
	}
	fmt.Println()

	//别名的定义,是不是相当于java的继承呢?
	var ff FF = 4
	var ft FF = 5
	dd := ff + ft
	fmt.Printf("%v.....%T\n", dd, dd)

	var str STR = "ss"
	fmt.Printf("%s.....%T\n", str, str)

	//字符的使用
	var ch int = '\u0041'
	var ch1 int = '\u03b2'
	var ch2 int = '\U00023674'
	fmt.Printf("%d,%d,%d \n", ch, ch1, ch2)
	fmt.Printf("%c,%c,%c\n", ch, ch1, ch2)
	fmt.Printf("%X,%X,%X\n", ch, ch1, ch2)
	fmt.Printf("%U,%U,%U\n", ch, ch1, ch2)

	//判断字符的有效性
	fmt.Println(unicode.IsDigit(int32(ch)))
	fmt.Println(unicode.IsLetter(int32(ch)))
	fmt.Println(unicode.IsSpace(int32(ch)))

	//字符串,go中额字符串用的是UTF8进行编码的
	//所以导致一个字符需要的字节范围是1~4
	//string有两种定义方式
	//1. 通过""双引号来进行定义,但是里面转义会执行
	var str1 string = "hello \t world \n welcome to \u0002 china"
	fmt.Printf("%s....%T", str1, str1)

	fmt.Println()
	//2. 通过``来进行定义,里面的东西不会转义
	var str2 string = `hello \t world \n l love you \u0004 china
			very much \U00000001`
	fmt.Printf("%s.....%T\n", str2, str2)

	//string 的比较是== <= >= != > < 比较的是字节
	//不像java比较的内存地址,因为go中的string是值类型,不是引用类型
	var array []int = []int{3, 5, 6, 7}

	fmt.Println("int address:", &array[2])

	//字符串也可以用下标来获取他的字节,但是一般用于全部都是ascii码
	var str5 = "china i love you "
	fmt.Printf("%c......%T", str5[3], str5[3])
	// fmt.Printf("%d......%T", &str5[2], &str5[2]) // 字符串的字符的不能地址
	fmt.Printf("%d......%T\n", &str5, &str5)

	str6 := "begining of the string" +
		"second part of string"

	fmt.Printf("%s....%T\n", str6, str6)

	str7 := "hello"
	str7 += "world"
	fmt.Printf("%s.....%T\n", str7, str7)

	//+来连接string不是最高效的方式,应该用strings.Join 或者更高效用byte-buffer

	var str8 = "asSASA ddd dsjkdsjs dk"
	bytes,size := stringsutil(str8)
	fmt.Println("bytes and zie :" , bytes,size)

	var str9 = "asSASA ddd dsjkdsjsこん dk"
	bytes2,size2 := stringsutil(str9)
	fmt.Println("bytes and zies2:" ,bytes2,size2)
}

//计算string的字节数和字符数
func stringsutil(str string) (int, int) {
	var bytes = len(str)
	var char = utf8.RuneCountInString(str)
	return bytes, char
}

//自定义一个字符串来进行继承
type STR string

type FF int

type BigFlat int

const (
	Active BigFlat = 1 << iota
	Send
	Recice
)

type ByteSize float64

//通过移位来定义一些常量
const (
	_           = iota
	KB ByteSize = 1 << (10 * iota)
	MB
	GB
	TB
	PB
	EB
	ZB
	YB
)

//float转化为int
func floatconver(f float64) (int, error) {
	if f >= math.MinInt32 && f <= math.MaxInt32 {
		in, decimal := math.Modf(f)
		if decimal > 0.5 {
			in++
		}
		return int(in), nil
	}
	return 0, fmt.Errorf("%f out of range", f)
}

//将int类型转化为uint类型
func intconver(x int) (uint16, error) {
	if 0 < x && x < math.MaxUint16 {
		return uint16(x), nil
	}
	return 0, fmt.Errorf("%d is out range of uint16", x)
}

//判断是否分支能够进入
func isbool() bool {
	fmt.Println("enter the is bool method")
	return false
}
