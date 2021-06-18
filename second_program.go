package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"reflect"
	"strconv"
	"strings"
)

func main() {
	min := 34
	max := 99
	randomValue := rand.Intn((max - min) + min)
	fmt.Println("Hint: ", randomValue)
	fmt.Println(reflect.TypeOf(randomValue))

	for i := 0; i < 2; i++ {
		fmt.Print("Try to guess generated number: ")
		reader := bufio.NewReader(os.Stdin)
		input, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}

		input = strings.TrimSpace(input)
		element, err := strconv.Atoi(input)

		if err != nil {
			log.Fatal(err)
		}
		if element < randomValue {
			attemptsNumber := 1 - i
			fmt.Printf("Value %v is less. \n", element)
			fmt.Printf("You have %v attempts. \n", attemptsNumber)
			checkNumber(attemptsNumber, randomValue)
		}
		if element > randomValue {
			attemptsNumber := 1 - i
			fmt.Printf("Value %v is more. \n", element)
			fmt.Printf("You have %v attempts. \n", attemptsNumber)
			checkNumber(attemptsNumber, randomValue)
		}
		if element == randomValue {
			fmt.Printf("Value %v is equal to randomly generated. \n\n", element)
			break
		}
	}
}
func checkNumber(numberOfIteration int, randomValue int) {
	if numberOfIteration == 0 {
		fmt.Printf("There is no attempts more.\nThe number of iteration is %v You quiz number was %v\n", numberOfIteration, randomValue)
	}
}
