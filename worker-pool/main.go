package main

import (
	"fmt"
	"sync"
)

type Job struct {
	Id   int
	Data string
}

type Result struct {
	Id   int
	Data string
}

func main() {
	const jobs = 20000
	const workers = 10

	job := make(chan Job, jobs)
	result := make(chan Result)

	wg := sync.WaitGroup{}

	go func() {
		for i := range result {
			fmt.Println("Result:", i)
		}
	}()

	for i := 0; i < workers; i++ {
		wg.Add(1)
		go worker(i, job, result, &wg)
	}

	for i := 0; i < jobs; i++ {
		job <- Job{Id: i, Data: fmt.Sprintf("Job %d", i)}
	}
	close(job)

	wg.Wait()
	close(result)
}

func worker(workerId int, jobs <-chan Job, results chan<- Result, wg *sync.WaitGroup) {
	defer wg.Done()
	for job := range jobs {
		results <- Result{Id: job.Id, Data: fmt.Sprintf("Processed %s by Worker %d", job.Data, workerId)}
	}
}
