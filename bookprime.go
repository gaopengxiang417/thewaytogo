package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("start ..." , time.Now())
	out := stive()
	fmt.Println("end ..." , time.Now())
	for{
		fmt.Println(<-out)
	}
}

//首先生成质素的管道
func generateNumber() chan int {
	ch := make(chan int)
	go func() {
		for i := 2; i < 10000000; i++ {
			ch <- i
		}
	}()
	return ch
}

//过滤的功能
func filterNumber(in chan int, prime int) chan int{
	out := make(chan int)
	go func() {
		for {
			if i := <-in; i%prime != 0 {
				out <- i
			}
		}
	}()
	return out
}

func stive() chan int {
	out := make(chan int)
	go func() {
		ch := generateNumber()
		for {
			prime := <-ch
			out <- prime
			ch = filterNumber(ch, prime)
		}
	}()

	return out
}
