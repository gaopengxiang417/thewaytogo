package main

import (
	"fmt"
	"time"
)

const NUM = 100000000

func main() {
	var start = time.Now()
	fmt.Println("start....", start)
	var count = 0
label1:
	for i := 1; ; i = i + 3 {
	label2:
		for j := 0; j <= 1; j++ {
			// var tmp = 2*(i+j) - 1
			var tmp = (3*i+j)<<1 - 1
			if tmp > NUM {
				break label1
			}
			for k := 2; k*k < NUM; k++ {
				if tmp%k == 0 {
					if j == 0 {
						continue label2
					} else {
						continue label1
					}
				}
			}
			count++
		}
	}

	fmt.Println("end....", time.Since(start).Nanoseconds()/1000000, "ms")
	fmt.Println("count:", count)
}
