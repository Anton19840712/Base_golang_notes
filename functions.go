package main

import "fmt"

//func add (x, y int) int {
func add (x int, y int) int { //pay attention: variable, than type
	result := x + y
	return result
}

//interesting syntax
func calc(x, y int)(int, int){ //first types in, second types out
	res1 := x + y //remembe about syntax like :=
	res2 := x - y
	return res1, res2 // here you try to aim in a signature
}

func main(){
	num1 := 56
	num2 := 35

	fmt.Println(num1 + num2)

	result := add(num1, num2)
	fmt.Println(result)

	calc1, calc2 := calc(num1, num2)
	fmt.Println(calc1, calc2)

	result1, result2 := calcadd(num1, num2)
	fmt.Println(result1, result2)
}

//not interesting syntax
func calcadd(x, y int)(res1, res2 int){
	res1 = x + y
	res2 = x - y
	return
}