package main

import (
	"fmt"
	"runtime"
	"time"
)

func Subscriber () {
	for {
		select {
			case _, ok := <- ch1:
				if !ok {
					fmt.Println("Breaking out of ch1")
					ch1 = nil
					break
				}
				fmt.Println("Running ch1 inner loop")
				for i := 0; i < 10; i++ {
					fmt.Println("Ch1 prints!", i)
					time.Sleep(time.Second)
				}
			case _, ok := <-ch2:
				if !ok {
					fmt.Println("Breaking out of ch1")
					ch1 = nil
					break
				}
				fmt.Println("Running ch2 inner loop")
				for i := 0; i < 10; i++ {
					fmt.Println("Ch2 prints!", i)
					time.Sleep(time.Second)
				}
		}

		if ch1 == nil || ch2 == nil {
			break
		}
	}
	//exitChannel <- true
}
var ch1 chan int = make(chan int)
var ch2 chan int = make(chan int)
//var exitChannel chan bool = make(chan bool)

func main() {
	fmt.Println("GOMAXPROCS = ", runtime.GOMAXPROCS(0))
	fmt.Println("NumCPU = ", runtime.NumCPU())
	go Subscriber()
	go Subscriber()
	ch1 <- 1
	ch2 <- 2
	close(ch1)
	close(ch2)
	time.Sleep(12 * time.Second)
	//<- exitChannel
	//close(exitChannel)
}