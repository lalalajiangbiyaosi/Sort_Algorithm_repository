package main

import (
	"fmt"
	"sync"
	"time"
)

var workCount int = 4
var wg sync.WaitGroup

type Job struct {
}

var cancelChan chan struct{} = make(chan struct{})
var JobChan chan Job = make(chan Job, 1000)

func process(job Job) {
	fmt.Println("Processing ")
	time.Sleep(time.Second)
}
func worker(jobChan <-chan Job, cancelChan <-chan struct{}) {
	defer wg.Done()
	for {
		select {
		case <-cancelChan:
			return
		case job := <-jobChan:
			process(job)
		}
	}
}

func main() {
	for i := 0; i < 10; i++ {
		JobChan <- Job{}
	}
	for i := 0; i < workCount; i++ {
		wg.Add(1)
		go worker(JobChan, cancelChan)
	}

	time.Sleep(time.Second)
	for {
		
	}
	// wg.Wait()
}
