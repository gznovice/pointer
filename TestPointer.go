package main

import "fmt"

import "TestPointer/other"

func main(){
	testPointer1()
	testPointer2()
	other.TestPointer3()
}

func testPointer2(){
	base := 0
	x := base
	y := base
	px, py, px2 := &x, &y, &x
	
	fmt.Println(x == y)
	fmt.Println(px == py)
	fmt.Println(px == px2)
	
}