package main

import (
	"fmt"
)

//здесь правило более простое, чем в c# - если не хватает до следующего целого, то тогда только это
func main() {
	fmt.Println("Hello, world")
	number1 := 9.2
	number2 := 9.5
	number3 := 9.7

	number4 := 10.2
	number5 := 10.5
	number6 := 10.7

	num1 := int(number1)
	num2 := int(number2)
	num3 := int(number3)
	num4 := int(number4)
	num5 := int(number5)
	num6 := int(number6)

	fmt.Println(num1, num2, num3, num4, num5, num6)

}
