package main

import (
	"fmt"
	"math"
	"time"
)

func main() {
	fmt.Println("start ...", time.Now())
	ch := make(chan int)

	const max = 10000000
	const max2 = 1000000
	go processPrimeCount(1, 1000000, ch)
	go processPrimeCount(1000001, 2000000, ch)
	go processPrimeCount(2000001, 3000000, ch)
	go processPrimeCount(3000001, 4000000, ch)
	go processPrimeCount(4000001, 5000000, ch)
	go processPrimeCount(5000001, 6000000, ch)
	go processPrimeCount(6000001, 7000000, ch)
	go processPrimeCount(7000001, 8000000, ch)
	go processPrimeCount(8000001, 9000000, ch)
	go processPrimeCount(9000001, 10000000, ch)

    var count = 0
    for i := 0; i < 10; i++ {
    	count += <- ch
    }
	fmt.Println(count)
	fmt.Println("end...", time.Now())

}

//进行计算质数的逻辑
func processPrimeCount(min int, max int, out chan int) {
	var count = 0
labe:
	for i := min; i < max; i++ {
		if i%2 == 0 {
			continue
		}
		//计算其平方根
		sqart := int(math.Sqrt(float64(i)))

		for j := 3; j < sqart; j = j + 2 {
			if i%j == 0 {
				continue labe
			}
		}
		count++
	}
	out <- count
}
