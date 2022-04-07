package main

import "fmt"

func testPointer1(){
	p0 := new(int)	//a pointer points to int value
	fmt.Println(p0)
	fmt.Println(*p0)
}