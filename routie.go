package main

import (
	"fmt"
	"math"
	"time"
)
const NUM = 10000000

func main() {
	fmt.Println("start...." , time.Now())
	ch := make(chan int)
	go compute(ch)
	fmt.Println("count:" , <-ch)
	fmt.Println("end..." , time.Now())
}

func compute(ch chan int) {
	var boos [NUM]bool
	boos[0] = true
	boos[1] = true
    
    sqrt := int(math.Sqrt(float64(NUM)))

    for i := 2 ; i < sqrt ; i++{
    	// var mod int = NUM / i
        for j := 2 ; (i * j) < NUM ; j++{
        	boos[i * j] = true
        }
    }

    var count = 0
    for _,is := range boos{
        if !is{
        	count++
        }
    }
    ch <- count
}
