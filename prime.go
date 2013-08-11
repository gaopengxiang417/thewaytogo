package main

import (
	"fmt"
	"math"
	"time"
)

func main() {
	//首先定义最大的范围
	var max int = 10000000
	var count int 

	var start = time.Now()
	fmt.Println("start ...." , start)
	labe:
	for i := 2 ; i < max ; i++ {
		if i % 2 == 0{
			continue
		}
		sqat := int(math.Sqrt(float64(i)))
		for j := 3 ; j <sqat ; j = j + 2{
			if i % j == 0{
				continue labe
			}
		}
		count++
	}

	fmt.Println("end ...." , time.Now())

	fmt.Println("total :" , count + 1)

	fmt.Println("total :" , (time.Since(start).Nanoseconds() / 1000000))
}
