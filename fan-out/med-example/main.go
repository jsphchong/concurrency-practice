package main

import (
	"fmt"
	"time"
)

// Borrowed from example of Medium, improved formatting 

func main() {
	source := Generate("for")
	process := GetProcessor()

	for i := 0; i < 12; i++ {
		process.PostJob(source)
	}
}

func Generate(data string) <-chan string {
	channel := make(chan string)
	go func() {
		for {
			channel <- data
			time.Sleep(1000)
		}
	}()

	return channel
}

func GetProcessor() *Processor {
	p := &Processor{
		JobChannel: make(chan string),
		Workers:    make([]*Worker, 5),
		Done:       make(chan *Worker),
	}

	for i := 0; i < 5; i++ {
		w := &Worker{
			Name: fmt.Sprintf("<Worker - %d>", i),
		}
		p.Workers[i] = w
	}
	p.StartProcess()
	return p
}
