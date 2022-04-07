package other

import "fmt"


func TestPointer3(){
	base := 4
	double(&base)
	
	fmt.Println(base)
}

func double(x *int){
	*x += *x
}