package main

import "fmt"
type Student struct {
	rollno int
	name string
	marks int
}

func main() { //need the same as name of the package
	var s1 Student = Student{101, "Navin", 55}
	fmt.Println("Student1:", s1)
	fmt.Println("Student1:", s1.name)
	var s2 Student = Student{rollno: 101, marks: 55}
	fmt.Println("Student2:", s2)
	fmt.Println("Student2:", s2.marks)
	fmt.Println("Student2:", s2.name)
}