package main

import (
	"fmt"
	"math/rand"
	"sync"
)

func main() {
	sum := 0

	c1 := make(chan int)
	c2 := make(chan int)

	xc := []chan int{c1, c2}

	for _, v := range xc {
		go place(v)
	}

	fanIn := make(chan int)

	go funneling(fanIn, c1, c2)
	for v := range fanIn {
		sum += v
	}

	fmt.Println("Program ending, final sum:", sum)
}

func place(c chan<-int){
	for i := 0; i < 5; i++ {
		c <- rand.Intn(10) + 1 
	}
	close(c)
}

func funnel(inChan <-chan  int, outChan chan<-int){
	for v := range inChan {
			outChan <- v
	}
}

func funneling(outChan chan<-int, xc ... <-chan int){
	var wg sync.WaitGroup

	for _, v := range xc {
		wg.Add(1)
		go func(inChan <-chan int){
			funnel(inChan, outChan)
			wg.Done()
		}(v) 
	}

	wg.Wait()
	close(outChan)
}    