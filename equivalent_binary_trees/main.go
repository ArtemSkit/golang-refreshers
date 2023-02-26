package main

import (
	"fmt"

	"golang.org/x/tour/tree"
)

//Exercise: Equivalent Binary Trees

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	if t == nil {
		return
	}
	Walk(t.Left, ch)
	ch <-t.Value
	Walk(t.Right, ch)
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) (bool, error) {
	ch1:=make(chan int)
	ch2:=make(chan int)
	readAll:=make(chan bool)

	go func() {
		Walk(t1, ch1)
		close(ch1)
	}()
	go func() {
		Walk(t2, ch2)
		close(ch2)
	}()

	go func() {
		for {
			if num1, ok1 := <- ch1; ok1 {
				if num2, ok2 := <- ch2; ok2 {
					if num1 == num2 {
						continue
					} else {
						readAll <- false
						break
					}
				} else {
					readAll <- false
					break
				}
			} else {
				if _, ok2 := <- ch2; ok2 {
					readAll <- false
					break
				} else {
					readAll <- true
					break
				}
			}
		}
		close(readAll)
	}()

	if res, ok := <- readAll; ok {
		return res, nil
	} else {
		return false, fmt.Errorf("channel that is used to signal the result was closed before the result was read")
	}
}

func main() {
	if res, err := Same(tree.New(100), tree.New(100)); err == nil {
		fmt.Println(res)
	} else {
		fmt.Println(err)
	}
}
