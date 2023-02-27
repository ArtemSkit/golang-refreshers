package main

import (
	"fmt"
)

type Tester interface {
	ValueReceiver()
	PointerReceiver()
	Both()
}

type TestStruct struct {
	value int
}

func (v TestStruct) ValueReceiver () {
	fmt.Printf("This is value receiver! %+v value at the address %p\n", v, &v)
}

func (p *TestStruct) PointerReceiver () {
	fmt.Printf("This is pointer receiver! %+v value at the address %p\n", p, p)
}

func (v TestStruct) Both () {
	fmt.Println("This is both value receiver!")
}

// Will not compile since Both is already declared
// func (p *TestStruct) Both () {
// 	fmt.Println("This is both pointer receiver!")
// }

func ValueFunction(v TestStruct) {
	fmt.Printf("Value function received value %+v at the address %p\n    ", v, &v)
	v.ValueReceiver()
	fmt.Print("    ")
	v.PointerReceiver()
}

func PointerFunction(p *TestStruct) {
	fmt.Printf("Pointer function received value %+v at the address %p\n    ", p, p)
	(*p).ValueReceiver()
	fmt.Print("    ")
	(*p).PointerReceiver()
}

func ValueInterfaceFunction(v Tester) {
	fmt.Printf("Value interface function received value %+v at the address %p\n    ", v, v)
	v.ValueReceiver()
	fmt.Print("    ")
	v.PointerReceiver()
}

func PointerInterfaceFunction(p *Tester) {
	fmt.Printf("Pointer interface function received value %+v at the address %p\n    ", p, p)
	(*p).ValueReceiver()
	fmt.Print("    ")
	(*p).PointerReceiver()
}



type ValueMethodCaller interface {
	valueMethod()
}

type PointerMethodCaller interface {
	pointerMethod()
}

type TestPointerReceiverImplementationStruct struct {
	value int
}


func (v TestPointerReceiverImplementationStruct) valueMethod () {
	fmt.Printf("This is value receiver! %+v value at the address %p\n", v, &v)
}

func (p *TestPointerReceiverImplementationStruct) pointerMethod () {
	fmt.Printf("This is pointer receiver! %+v value at the address %p\n", p, p)
}

func callValueMethodOnInterface(v ValueMethodCaller) {
	fmt.Printf("callValueMethodOnInterface function received value %+v at the address %p\n    ", v, v)
	v.valueMethod()
}

func callPointerMethodOnInterface(p PointerMethodCaller) {
	fmt.Printf("callPointerMethodOnInterface function received value %+v at the address %p\n    ", p, p)
	p.pointerMethod()
}

func main() {
	value := TestStruct{5}
	pointer := &value

	value.ValueReceiver()
	value.PointerReceiver()
	pointer.ValueReceiver()
	pointer.PointerReceiver()

	ValueFunction(value)
	// cannot use value (variable of type TestStruct) as *TestStruct value in argument to PointerFunction
	//PointerFunction(value)
	// cannot use pointer (variable of type *TestStruct) as TestStruct value in argument to ValueFunction
	//ValueFunction(pointer)
	PointerFunction(pointer)

	// cannot use value (variable of type TestStruct) as Tester value in variable declaration: TestStruct does not implement Tester (method PointerReceiver has pointer receiver)
	// var interf Tester = value
	// var interfPointer *Tester = &value

	var interf Tester = pointer
	var interfPointer *Tester = &interf
	ValueInterfaceFunction(interf)
	// cannot use interf (variable of type Tester) as *Tester value in argument to PointerInterfaceFunction: Tester does not implement *Tester (type *Tester is pointer to interface, not interface)
	// PointerInterfaceFunction(interf)
	//cannot use interfPointer (variable of type *Tester) as Tester value in argument to ValueInterfaceFunction: *Tester does not implement Tester (type *Tester is pointer to interface, not interface)
	// ValueInterfaceFunction(interfPointer)
	PointerInterfaceFunction(interfPointer)



	fmt.Print("\n\n\n")
	var val     TestPointerReceiverImplementationStruct  = TestPointerReceiverImplementationStruct{5}
	var point *TestPointerReceiverImplementationStruct = &val
	fmt.Printf("Original val address is %p\n", &val)
	callValueMethodOnInterface(val)
	callPointerMethodOnInterface(point)
	callValueMethodOnInterface(point)
	// cannot use val (variable of type TestStruct) as PointerMethodCaller value in argument to callPointerMethodOnInterface: TestStruct does not implement PointerMethodCaller (method pointerMethod has pointer receiver)
	// callPointerMethodOnInterface(val)
}