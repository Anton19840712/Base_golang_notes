//just basic function
package main //need name main

import "fmt"

func main() { //need the same as name of the package
	num := 43
	if num == 1{
		fmt.Println(num)
	} else if num ==43 {
		fmt.Println(num)
	}else{
		fmt.Println("N/A")
	}

	//or
	num2 := 3

	switch num2 {
	case 1:
		fmt.Println(num)
	case 43:
		fmt.Println(num)
	default:
		fmt.Println("N/A")
	}
}