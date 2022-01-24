package main

import "fmt"

func main() {
	cs := make(chan<- int)

	go func() {
		cs <- 42
	}()
	// fmt.Println(<-cs) Invalid operation. 

	fmt.Printf("------\n")
	fmt.Printf("cs\t%T\n", cs)
}