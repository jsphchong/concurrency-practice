package main

import "fmt"

func main() {
	c := make(chan int)
	go send(c)

	receive(c)

	close(c)

	fmt.Println("About to exit")
}

// send channel
func send(c chan<- int) {
	c <- 42
}

func receive(c <-chan int) {
	fmt.Println("the value received from the channel:", <- c)
}