package main

import "fmt"

//there is only one syntax for for all kinds of cycles in go: only for
func main(){
//------------------------------------------
	for i := 0; i < 8; i++ {
		fmt.Println("Antonio, save your soul", i)
	}
//------------------------------------------
	for {
		fmt.Println("Antonio, save your soul")
	}//ctrl+c to exit
}



