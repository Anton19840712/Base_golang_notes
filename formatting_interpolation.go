//https://golang.org/pkg/fmt/
package main

import (
	"fmt"
	"strings" //writw new one only from a new line, otherwise it does not work
)


func main() { //need the same as name of the package
	//interpolation
var element = "hello"
var element3 = 3
//using %s and %v to print interpolated strings.
//and Printf to print formated string
fmt.Printf("\nThe output is %s and %v", element, element3) 
fmt.Println("\nI am", 26, "years old")
fmt.Println("\nI am", element3, "years old") //I am 3 years old

// Sprint and Print functions work in a similar fashion. 
// Sprint does everything what Print does but does not write 
// the resulting string to the standard output, instead, it returns it.
//s1 := fmt.Sprint("\nI am", element3, "years old") //BUG in function Sprint: I am3years old, need: "I am 3 years old"
s1 := fmt.Sprintln("\nI am", element3, "years old") //But here all OK
fmt.Println(s1)

// struct
s := struct {	name string
	age  int}{"John", 26}
// channel
c := make(chan int)
// map
m := map[string]int{"one": 1,	"two": 2}

fmt.Printf("%v \n", true)           // boolean values //true
fmt.Printf("%v \n", 132)            // signed and unsigned //integers 132
fmt.Printf("%v \n", 198.454)        // floating point numbers //198.454
fmt.Printf("%v \n", 3+7i)           // complex numbers //(3+7i) 
fmt.Printf("%v \n", "Hello World!") // strings  // Hello World!
fmt.Printf("%v \n", []int{1, 2, 3}) // slices and arrays  // [1 2 3]
fmt.Printf("%v \n", m)              // maps // map[one:1 two:2] 
fmt.Printf("%v \n", s)              // structs // &{John 26}
fmt.Printf("%v \n", c)              // channels // 0xc00001e0c0
//boolean
fmt.Printf("%t \n", true)
fmt.Printf("%t \n", false)
//upper, lower cases
fmt.Printf("%x\n", "JKβ") // default lowercase
fmt.Printf("%X\n", "JKβ") // upper-case
fmt.Printf("% x\n", "JKβ") // with space
fmt.Printf("% X\n", "JKβ") // upper-case with space
str3 := "hello! golang"
str4 := "lower CASE I NEED"
fmt.Println(strings.ToUpper(str3))
fmt.Println(strings.ToLower(str4))
}
