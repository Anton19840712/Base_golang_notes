package main

import (
	"fmt"
)

//To create dynamically sized arrays
func main() {

	sliceA := make([]int, 3, 6) //creates array sized to 3  but the highest size could be increased dynamically to 6 (params := type, length, capacity) here make returns a slice - smth similar to list in dot .net,
	//but the max sized level is determined by developer

	fmt.Println(sliceA) //[0 0 0]
	fmt.Println(len(sliceA))
	fmt.Println(cap(sliceA))

	sliceB := make([]int, 6) //creates array sized to 6

	fmt.Println(sliceB) //[0 0 0 0 0 0]

}
