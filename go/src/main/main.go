package main

import(
	"fmt"
)

var s = "Hello World"

func main(){
	fmt.Printf("print 's' variable from outer block %v\n", s)
	b:=true
	if b{
		fmt.Printf("printing 'b' varible from outer block %v\n",b)
		i:=42
		if b!=false{
			fmt.Printf("printing 'i' variable from outer block %v\n",i)
		}
	}
}