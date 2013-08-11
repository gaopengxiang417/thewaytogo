package main

import (
	"fmt"
)

func main() {
	labe:
	for i := 0; i < 20; i++ {
		for j := 0 ; j < 20 ; j++ {
			if i == j {
				continue labe
			}
			fmt.Println("i=" , i ,",j=" , j)
		}
	}
}

