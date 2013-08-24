package main

import (
	"fmt"
)

func main() {
	forfirst()
	fmt.Println()
	forsecond()
	iterstring()
	fortestone()
	gototest()
	towertest()
	fmt.Println()
	bitwisetest()
	fizzBuzz()
	rectangle_starts()
	fornotEach()
	forrangeTest()
	forlocal()
	fornotnormal()
	//testto
	//four
	//fiv
	//six
	breaktestone()
	continuetest()
	labeltest()
	gotofunctest()
	gotoerror()
	exciseonefor()
	fmt.Println()
	foreach()
}
func foreach() {
	for i := 0; i < 8; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Printf("%d ", i)
	}
}
func exciseonefor() {
	i := 0
	for {
		if i >= 3 {
			break
		}
		fmt.Printf("%d ", i)
		i++
	}
	fmt.Println("\n a statement is just after the loop")
}

//测试goto,label中间不能有变量初始化的信息
func gotoerror() {
	a := 1
	goto TARGET

TARGET:
	b := 2
	b += a
	fmt.Printf("a = %v,b = %v", a, b)
}
func gotofunctest() {
	i := 0
HERE:
	for {
		fmt.Println(i)
		i++
		if i == 8 {
			return
		}
		goto HERE
	}
}
func labeltest() {
LABEL1:
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if j == 5 {
				continue LABEL1
			}
			fmt.Println(i, j)
		}
	}
}

func continuetest() {
	for i := 0; i < 10; i++ {
		if i == 5 {
			continue
		}
		fmt.Print(i, " ")
	}
}

//测试break语句
func breaktestone() {
	for i := 0; i < 3; i++ {
		for j := 0; j < 10; j++ {
			if j > 5 {
				break
			}
			fmt.Print(j)
		}
		fmt.Print(" ")
	}
}

func fornotnormal() {
	s := ""
	for s != "aaaaaa" {
		fmt.Println("s value :", s)
		s += "a"
	}

	for i, j, s := 0, 5, "a"; i < 3 && j < 100 && s != "aaaaaa"; i, j, s = i+1, j+1, s+"a" {
		fmt.Println("value is :", i, j, s)
	}
	//斯循环
	/*for i := 0; i < 3; {
		fmt.Println("value of i :", i)
	}*/
}

//测试局部变量
func forlocal() {
	for i := 0; i < 5; i++ {
		var a int
		fmt.Printf("%d ", a)
		a = 5
	}
}

//测试如果for range的时候是否是拷贝
func forrangeTest() {
	var str string = "hello wolrd忘尘"
	for index, ch := range str {
		fmt.Println("index:", index, ",ch:", ch)
	}

	fmt.Println("**********华丽的分割线**********")
	for index, ch := range str {
		ch = 'w'
		fmt.Printf("%d:%c", index, ch)
	}

	fmt.Println(str)

	for index, ch := range str {
		fmt.Printf("%-2d %d %U '%c' % X \n", index, ch, ch, ch, []byte(string(ch)))
	}
}

//测试for循环充当while的作用
func fornotEach() {
	var a int = 6
	for a > 0 {
		a -= 1
		fmt.Println(a)
	}
}

//打印矩形
func rectangle_starts() {
	var width int = 20
	var height int = 10
	for i := 0; i < height; i++ {
		if i == 0 || (i == height-1) {
			for j := 0; j < width; j++ {
				fmt.Print("*")
			}
			fmt.Println()
		} else {
			fmt.Print("*")
			for j := 0; j < width; j++ {
				fmt.Print(" ")
			}
			fmt.Print("*")
			fmt.Println()
		}
	}
}

func fizzBuzz() {
	for i := 0; i < 100; i++ {
		switch {
		case i%15 == 0:
			fmt.Println("FizzBuzz")
		case i%3 == 0:
			fmt.Println("Fizz")
		case i%5 == 0:
			fmt.Println("Buzz")
		default:
			fmt.Println(i)
		}
	}
}

//show bitwise
func bitwisetest() {
	for i := 0; i < 10; i++ {
		fmt.Print("the bit repreesion is %b\n", i)
	}
}

//print character
func towertest() {
	var str string
	for i := 0; i < 26; i++ {
		str = str + "G"
		fmt.Println(str)
	}
}

//goto
func gototest() {
	for i := 0; i < 20; i++ {
		fmt.Printf("%d \t", i)

		if i == 19 {
			goto label
		}
	}
label:
	fmt.Println("end......")
}

//teest
func fortestone() {
	for i := 0; i < 15; i++ {
		fmt.Printf("%d \t", i)
	}
	fmt.Println()
}

func iterstring() {
	var golang string = "go is a beatiful language"

	fmt.Printf("the length of %s is : %d\n", golang, len(golang))

	for i := 0; i < len(golang); i++ {
		fmt.Printf("the %d character is %c\n", i, golang[i])
	}

	fmt.Println()

	var chinse = "中国人"
	fmt.Printf("the length of chinese is %d\n", len(chinse))

	for i := 0; i < len(chinse); i++ {
		fmt.Printf("the %d character is %c\n", i, chinse[i])
	}
}

func forsecond() {
	i := 0
	for ; i < 5; i++ {
		fmt.Printf("%d \t", i)
	}
}

func forfirst() {
	for i := 0; i < 5; i++ {
		fmt.Printf("%d \t", i)
	}
}
