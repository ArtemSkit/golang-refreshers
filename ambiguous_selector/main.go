package main

import "fmt"

type A struct {
	n, m int
}

type B struct {
	n string
}

type collectror struct {
	A
	B
}

func (v A) run() {
	fmt.Println("Type A method \"run\"", v.n)
}

func (v A) runA() {
	fmt.Println("Type A method \"runA\"", v.m)
}

func (v B) run() {
	fmt.Println("Type B method \"run\"", v.n)
}

func (v collectror) run() {
	// Ambiguous type field selector
	fmt.Println("Type collector method \"run\"", v.A.n, v.m, v.B.n)
}

func main() {
	u:=collectror{A{5, 10}, B{"Hello"}}

	// Ambiguous type method selector
	u.run()
	u.runA()
	u.A.run()
	u.B.run()
}