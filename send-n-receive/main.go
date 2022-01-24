package main

import "fmt"

func main() {
	c := make(chan int)

	go func() { // Places 42 on the channel within go routine so doens't block 
		c <- 42
	}()

	fmt.Println(<-c) // Channel receives input 

	c2 := make(chan int, 1) // can also use buffer
	c2 <- 42
	fmt.Println(<-c2)
}
