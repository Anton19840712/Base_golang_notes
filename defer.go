//just basic function
package main //need name main

import "fmt"

func main() { //need the same as name of the package
	for i:=0;i<3;i++{
		fmt.Println(i)
		defer fmt.Println(i, "defer") //means pass it to the end of the execution plan and:
		fmt.Println(i)
	}
	fmt.Println("Hello, world!")
}