//just basic function
package main //need name main

import "fmt"

func main() { //need the same as name of the package
	for i := 0; i < 3; i++ {
		fmt.Println(i)
		defer fmt.Println(i, "defer") //means pass it to the end of the execution plan and and execute in after you execute all your nested code
		// by the way defer works by LIFO principle
		// this is like postpone
		fmt.Println(i)
	}
	fmt.Println()
	fmt.Println("End of official cycle and start of deffered execution. That was proposed above.")
	fmt.Println()
}
