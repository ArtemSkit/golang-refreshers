package main

import "fmt"

	type inter interface {
		run()
	}
	type testType struct {
		num int
	}

	func (p *testType) run() {
		p.num++
		fmt.Println("run")
	}
func TestInterfaceAssignment() {
	var x inter = &testType{5}
	y := x
	x.run()
	fmt.Printf("x %p y %p\n", x, y)
	fmt.Println(x, y)
}