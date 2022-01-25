package main

import (
	"fmt"
	"time"
)

type Processor struct {
	JobChannel chan string
	Done       chan *Worker
	Workers    []*Worker
}

func (p *Processor) StartProcess(){
	go func(){
		for {
			select {
			default:
				if len(p.Workers) > 0{
					w := p.Workers[0]
					p.Workers = p.Workers[1:]
					w.ProcessJob(<- p.JobChannel, p.Done)
				}
			
			case w := <-p.Done:
				p.Workers = append(p.Workers, w)
			}
		}
	}()
}

func (p *Processor) PostJob(jobs <-chan string){
	p.JobChannel <- <-jobs
}

type Worker struct {
	Name string
}

func (w *Worker) ProcessJob(data string, done chan *Worker) {
	go func() {
		fmt.Println("Working on data", data, w.Name)
		time.Sleep(3000)
		done <- w
	}()
}