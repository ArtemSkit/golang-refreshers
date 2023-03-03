package main

import (
	"fmt"
	"math/rand"
	"time"
)


func main() {
	ch :=make(chan int)
	closeSignalChannel :=make(chan bool)
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	for i := 0; i < 10; i++ {
		go func (num int) {
			select {
				case _, ok := <-closeSignalChannel:
					if !ok {
						fmt.Println("Returning from iterantion ", num)
						return
					}
				case <-time.After(time.Duration(r1.Intn(2000))*time.Millisecond):
					fmt.Println("Sending from iterantion ", num)
					ch <- 15
			}
		}(i)
	}

	fmt.Println(<-ch)
	close(closeSignalChannel)
	time.Sleep(3*time.Second)
}