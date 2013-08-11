package main

import (
	"fmt"
	"math"
	"time"
)

const NUM = 10000000

func main() {

	var start = time.Now()
	fmt.Println("start", start)

	// var boos [NUM]bool
	boos := make([]bool, NUM)

	/*for i := 0; i < NUM; i++ {
		boos[i] = true
	}*/

	boos[0] = false
	boos[1] = true

	/*//爱拉托逊斯筛选法
	for i := 2; i < NUM; i++ {
		for j := i * 2; j < NUM; j = j + i {
			boos[j] = false
		}
	}*/

	//改进的爱拉托逊斯筛选法
	sqar := int(math.Sqrt(float64(NUM)))
	for i := 2; i < sqar; i++ {
		if !boos[i] {
			for j := i * i; j < NUM; j = j + i {
				boos[j] = false
			}
		}
	}

	var count = 1
	for i := 2; i < NUM; i++ {
		if boos[i] {
			count++
		}
	}

	var end = time.Now()
	fmt.Println("end....", end)

	fmt.Println("count:", count)
	fmt.Println("total time : ", time.Since(start).Nanoseconds()/1000000, "ms")
}
