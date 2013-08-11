package main

import (
	"fmt"
	"math"
	"time"
)

const NUM = 10000000

func main() {

	var start = time.Now()
	fmt.Println("start....", start)

	var boos = make([]bool, NUM+1)

	var sqar = int(math.Sqrt(float64(NUM)))

	var count = 1
	for factor := 3; factor <= sqar; factor = factor + 2 {
		if boos[factor] {
			continue
		}
		for i := factor * factor; i < NUM+1; i = i + 2*factor {
			if !boos[i] {
				count++
				boos[i] = true
			}

		}
	}
	count = count + NUM/2 - 1 //这里加上所有的偶数并且减去(2)这个特殊情况

	fmt.Println("count:", NUM-count)

	fmt.Println("time:", time.Since(start).Nanoseconds()/1000000, "ms")

}
