package main

import "fmt"

func main() {
	// slice1:=make([]int, 5, 20)
	// var test []int = make([]int, 3)
	// test = slice1[:]
	// //copy(test, slice1)
	// slice2:=append(slice1, 88)
	// fmt.Println(slice1, slice2)
	// fmt.Printf("%p %p\n", slice1, slice2)
	// fmt.Println(len(slice1), len(slice2))
	// slice2 = append(slice2, 77)

	// fmt.Println(len(slice1), len(slice2))
	// fmt.Printf("%p %p\n", test, slice1)
	// fmt.Println(test)
	array:=[]int{0,1,2}
	slice:=make([]int, 3)
	slice[0] = array[0]
	fmt.Printf("%p %p\n", slice, array)
}