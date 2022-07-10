package main

import (
	"fmt"
	"math/rand"
)

func selectWaiterAndChef() {

	allChefs := []string{"Jack", "Bob", "Mark"}
	allWaiters := []string{"A", "B", "C"}

	randomC := rand.Intn(len(allChefs))
	randomW := rand.Intn(len(allWaiters))
	cpick := allChefs[randomC]
	wpick := allWaiters[randomW]

	return cpick, wpick
}

// Main function
func main() {

	orderChan := make(chan chefOrder)
	cookChan := make(chan chefCooked)

	chef, waiter := selectWaiterAndChef()

	fmt.Println(chef)
	fmt.Println(waiter)
}
