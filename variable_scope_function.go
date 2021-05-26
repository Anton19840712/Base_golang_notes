package main

import "fmt"

var a = 9 // package level variable, thats way it is seen from demo() scope
var s = 90
func demo()(element int){
	//var a = 9 // function level variable
	s := 32
	a := 75
	fmt.Println(s, a)
	return s
}

func main(){
	sm := demo()
	fmt.Println(sm, s)
}
