package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	runtime.GOMAXPROCS(1)
	ch := make(chan int)
	go func() {
		ch <- 1
	}()
	//time.Sleep(time.Second)
	go func() {
		select {
		case num := <-ch:
			fmt.Println(num)
		default:
			fmt.Println("default")
		}
	}()

	time.Sleep(2 * time.Second)
}
