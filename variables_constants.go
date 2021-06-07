package main

import (
	"fmt"
)

func main(){ 

	// var num1, num2 int
	// num1, num2 = 2, 3
	// num1 := 34 // the same as var result = 9 --this is my choise

	const num = 9 //you cannot reassign it after like num = 12, for instance
	var length, width float64 //declaratio
	length, width = 1.2, 2.4 //assignment

	fmt.Println(length, width)

	var num0 int // == 0
	var num1 int = 2
	var num2 = 3
	var result = num0 + num1 + num2
	fmt.Println(result)

	//WITHOUT VARIABLE TYPE, missing variable type mistake
	element := "hello"
	// so := sign - is for dynamic typization, like in a python, no need to declare type, when we use :=
	//BUT 
	var element34 = "hello34"
	fmt.Println(element, element34)	
}