package main

import (
	"fmt"
	"strings"
)

func main(){
	broken := "G# r#cs!"
	fixed := strings.Replace(broken, "#", "o", 0) //цифра показывает, с какого индекса будет происходить замена, можно написать и -2.
	//но можно написать иначе
	fixedElse := strings.NewReplacer("#","o")	
	fmt.Println(fixed)
	fmt.Println(fixedElse.Replace(broken))
}