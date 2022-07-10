package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main(){
		fmt.Print("Enter a grade: ")
		reader := bufio.NewReader(os.Stdin)
		input, err := reader.ReadString('\n')
		if err!=nil{
			log.Fatal(err)
		}
		input = strings.TrimSpace(input) //странно, почему нельзя :=//

		element, err := strconv.Atoi(input)
		if err!=nil{
			log.Fatal(err)
		} 
		
		if element > 35 {
			fmt.Println("Value is too much")
		}else {
			fmt.Printf("Value %v is suit for me.", input)
		}
}