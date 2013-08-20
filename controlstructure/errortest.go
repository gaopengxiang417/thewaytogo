package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {
	var str = "ABC"

	a, err := strconv.Atoi(str)

	if err != nil {
		fmt.Println("error exception", err, a)
		// os.Exit(1)
	} else {
		fmt.Println("normal branches....")
	}

	file, err1 := os.Open("G:\\java_magnize\\go\\sublime.txt")
	if err1 != nil {
		fmt.Println("error ....")
		// os.Exit(1)
	}

	var bs []byte = make([]byte, 2048)
	count, err2 := file.Read(bs)
	if count == 0 || err2 != nil {
		fmt.Println("exit....")
	}

	fmt.Println(string(bs))

	b, err3 := strconv.Atoi(str)
	if err3 != nil {
		fmt.Println("error.............", b)
	}

	//chmod
	if err := file.Chmod(0755); err != nil {
		fmt.Println("can not change mod")
	}

	//
	if ok := boolreturn(); ok {
		fmt.Println("enter the block")
	}

	result, ok := mysqar(43.5)
	fmt.Println(result, ok)

	fmt.Println(myAtoi("44"))
}

func myAtoi(str string) int {
	strr, _ := strconv.Atoi(str)
	return strr
}

func mysqar(f float64) (float64, bool) {
	if f < 0 {
		return 0, false
	}
	return math.Sqrt(f), true
}

func boolreturn() bool {
	return true
}
