//https://golang.org/pkg/fmt/
package main

import (
	"fmt"
	"math"
)

func main(){
	var num float64 = 9 //can not write :=
	var result = math.Sqrt(num)
	var result2 = math.Ceil(num)
	var intResult = math.Floor(result)
	var intResult2 = math.Floor(result2)
	fmt.Printf("The output is %.2f and %.2g", intResult, intResult2)
}