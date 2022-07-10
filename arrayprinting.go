package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello, world!")

	elements := []string{"Irina", "Igor", "Victor"}

	fmt.Println(elements) //[Irina Igor Victor]

	fmt.Println(elements[0]) //Irina

	fmt.Println()

	for i := 0; i <= len(elements)-1; i++ {
		fmt.Println(elements[i])
	}

}
