package main

/*
	if else 的学习
*/
import (
	"fmt"
	"runtime"
)

func main() {
	//其实 if (1 == 1){}也是支持的,只是一般fmt格式化的时候会把小括号去掉
	if 1 == 1 {
		fmt.Println("enter if ")
	}

	var boo = true

	if boo {
		fmt.Println("the value is true:", boo)
	} else {
		fmt.Println("the value is false:", boo)
	}

	if !(1 == 2) {
		fmt.Println("! 的使用")
	}

	reuslt := returndiff(4)
	fmt.Println("result : ", reuslt)

	//检查字符串是否为空
	fmt.Printf("is null string :%v\n", isEmpty(""))
	fmt.Printf("not null string :%v\n", isEmpty("china"))

	//
	fmt.Println(osname())

	//获取一个数的绝对值
	fmt.Printf("%d \n", abstest(-4))

	//返回最大的一个数
	fmt.Println(greator(4, 90))

	//if的另一种简单的形式
	if i := 6 ; 1 == 1{
		fmt.Println(i)
	}

	var ii int = 50
	var jj int

	if ii < 0{
		fmt.Println("the data is little than 0")
	}else if ii < 50{
		fmt.Println("the data is little than 50")
	}else{
		fmt.Println("the data is greator than 50")
	}

	if jj = 5 ; jj > 10{
		fmt.Println("jj is greator than 10")
	}else{
		fmt.Println("jj is little than 10")
	}

}

//比较连个数的大小
func greator(a, b int) int {
	if a > b {
		return a
	}
	return b
}

//获取一个数的绝对值
func abstest(i int) int {
	if i > 0 {
		return i
	}
	return -i
}

//获取操作系统名称
func osname() string {
	var prompt = "enter a digit e.g. 3" + " or %s to quic"
	var os = runtime.GOOS
	if os == "windows" {
		prompt = fmt.Sprintf(prompt, "ctrl + z")
	} else {
		prompt = fmt.Sprintf(prompt, "ctrl + c")
	}

	return prompt
}

//判断字符串视为为空
func isEmpty(str string) bool {
	if str == "" || len(str) == 0 {
		return true
	}
	return false
}
func returndiff(i int) int {
	if i == 3 {
		return 8
	} else {
		return 4
	}
}
