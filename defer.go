//just basic function
package main //need name main

import "fmt"

func main() { //need the same as name of the package
	for i:=0;i<3;i++{
		fmt.Println(i)
		defer fmt.Println(i, "defer") //means pass it to the end of execution plan and: 
		// at first print "Hello, world!", and then only for cycle and from the end of stack!!!
		fmt.Println(i)
	}
	fmt.Println("Hello, world!")
}