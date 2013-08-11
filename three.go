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
	var boos [NUM + 1]bool

	/*for i := 0; i < NUM+1; i++ {
		boos[i] = (i%2 != 0)
	}*/

	/*for i := 0; i < NUM+1; i = i + 2 {
		boos[i] = false
	}*/

	//不需要计算,直接给所有的2的倍数为false,非2的倍数为true,所以这里少了一些判断逻辑
	//从而能减少处理器的气泡
	for i := 1; i < NUM+1; i = i + 2 {
		boos[i] = true
	}

	boos[0] = false
	boos[1] = true

	var sqar = int(math.Sqrt(float64(NUM)))

	for factor := 3; factor <= sqar; factor = factor + 2 {
		for i := factor * factor; i < NUM+1; i = i + 2*factor {
			boos[i] = false
		}
	}

	var count = 1
	for i := 3; i < NUM+1; i = i + 2 {
		if boos[i] {
			count++
		}
	}

	fmt.Println("count:", count)

	fmt.Println("time:", time.Since(start).Nanoseconds()/1000000, "ms")

}
