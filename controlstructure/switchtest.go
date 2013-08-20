package main

import (
	"fmt"
)

func main() {
	switchone()
	switchtwo()
	switchthree()
	fmt.Println(season(5))
}

func season(i int) string {
	var ss string
	switch i {
	case 1:
		ss = "Jar"
	case 2:
		ss = "Feb"
	case 3:
		ss = "Mar"
	case 4:
		ss = "May"
	case 5:
		ss = "Aug"
	case 6:
		ss = "Feb"
	case 7:
		ss = "OCT"
	case 8:
		ss = "ELEV"
	case 9:
		ss = "dec"
	case 10:
		ss = "Nov"
	default:
		ss = "nor correct"
	}
	return ss
}
func switchthree() {
	var num = 6
	switch {
	case num < 0:
		fmt.Println("num is little than 0")
		fallthrough
	case num > 0 && num < 4:
		fmt.Println("num is little than 4")
		fallthrough
	case num >= 4 && num < 8:
		fmt.Println("num is little than 8")
		fallthrough
	default:
		fmt.Println("default brached")
	}
}

func switchtwo() {
	var num int = 8
	switch {
	case num < 0:
		fmt.Println("num is little than 0")
	case num > 0 && num < 5:
		fmt.Println("num is little than 5")
	case num < 10:
		fmt.Println("num is litte than 10")
	default:
		fmt.Println("greator than 10")
	}
}

func switchone() {
	var num = 100
	switch num {
	case 98, 99:
		fmt.Println(98, "or", 99)
	case 100:
		fmt.Println(100)
	default:
		fmt.Println("not find")
	}
}
