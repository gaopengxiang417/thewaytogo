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
}

//show bitwise
func bitwisetest() {
	for i := 0; i < 10; i++ {
		fmt.Print("the bit repreesion is %b\n" , i)
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
