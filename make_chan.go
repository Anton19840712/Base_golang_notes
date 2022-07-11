package main

import "fmt"

func main() {

	messages := make(chan string) // что здесь дает make?? make, типа, означает забодяжить что-то.

	go func() { messages <- "ping" }() // я так понимаю, словом go вызывается go рутина. Вызывается функция messages. Ключевое слово go является ключом к реализации параллелизма.

	msg := <-messages

	fmt.Println(msg)
}
