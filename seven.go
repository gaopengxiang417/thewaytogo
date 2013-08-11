package main

import (
	"fmt"
	"math"
	"time"
)

const NUM = 10000000

func main() {

	var start = time.Now()

	var small = NUM >> 1
	var boos = make([]bool, small)

	var sqar = int(math.Sqrt(float64(NUM)))

	/*for i := 0; i < small; i++ {
		boos[i] = true
	}*/
	for i := 3; i <= sqar; i = i + 2 {
		var k = (i - 3) >> 1
		if !boos[k] {
			for j := i * i; /*i * i*/ j < NUM; j = j + i<<1 {
				var m = (j - 3) >> 1
				boos[m] = true
			}
		}
	}

	var count = 0
	for i := 0; i < small; i++ {
		if !boos[i] {
			count++
		}
	}

	var time = time.Since(start).Nanoseconds() / 1000000
	fmt.Println("time:", time, "ms")
	fmt.Println(count)

}
