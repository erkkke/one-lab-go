package workerpool

import (
	"context"
	"errors"
	"fmt"
	"sync"
)

var ErrorCountOfErrExceeded = errors.New("error, count of errors exceeded")

func Execute(tasks []func() error, E int) error {
	var wg sync.WaitGroup
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	numJobs := len(tasks)
	numWorkers := 5
	jobs := make(chan func() error, numJobs)
	results := make(chan error, numJobs)
	errorCnt := 0

	for i := 1; i <= numWorkers; i++ {
		wg.Add(1)
		go worker(ctx, &wg, i, jobs, results)
	}

	for _, job := range tasks {
		jobs <- job
	}
	close(jobs)

	for i := 1; i <= numJobs; i++ {
		if <-results != nil {
			errorCnt++
		}
		if errorCnt >= E {
			cancel()
			return ErrorCountOfErrExceeded
		}
	}

	wg.Wait()
	return nil
}

func worker(ctx context.Context, wg *sync.WaitGroup, id int, jobs <-chan func() error, results chan<- error) {
	defer wg.Done()

	for {
		select {
		case <-ctx.Done():
			return
		case res, ok := <-jobs:
			if !ok {
				return
			}
			results <- res()
			fmt.Printf("#%v worker finished job\n", id)
		}
	}
}
