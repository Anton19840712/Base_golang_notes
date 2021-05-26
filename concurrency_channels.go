//chan key word
//transport data of only one data type
//c <- data --push or write data to the channel c
//<- c --read some data from channel c -- you can see there is no data to the right, thats way compiler will interpret it, that it read it, I think.
//var data int
//data = <- c/data := <- c --data coming from the channel c which is of type int can be stored into the variable data of type int.
package main

import "fmt"

func greet(c chan string) { //greet function which accepts a channel c of transport data type string
	//so, here we write not only  f int, or  n string, but we write variable name, 
	//than variable type and dedicate only one data type to it.
	fmt.Println("Hello, " + <-c + "!")
}

func main(){//active goroutine is main goroutine
	// c := make(chan int)
	// fmt.Printf("type of `c` is %T'\n", c)
	// fmt.Printf("value of `c` is %v\n", c)
	fmt.Println("main()started")
	c := make(chan string) //chanel creation using make function... interesting
	go greet(c) // we execute our previously created chan here, using go keyword. 
	// We pass it and go to the c <-"John" and only than to the function greet above
	c <-"John" // why we use it after? //main goroutine becomes blocke// blocking operation
	fmt.Println("main()stopped")
}