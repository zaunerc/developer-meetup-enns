package main

import (
	"fmt"
	"math"
	"sync"
)

func worker(tasksCh <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		task, ok := <-tasksCh
		if !ok {
			return
		}

		sqrt := math.Sqrt(float64(task))
		fmt.Printf("input: %d square root: %f\n", task, sqrt)
	}
}

func pool(wg *sync.WaitGroup, workers int, tasks int) {
	tasksCh := make(chan int)

	fmt.Printf("Spawning %d go routines...\n", workers)
	for i := 0; i < workers; i++ {
		go worker(tasksCh, wg)
	}
	fmt.Println("Finished spawning go routines.")

	//fmt.Println("Waiting before continuing...")
	//time.Sleep(10 * time.Second)
	//fmt.Println("Finished waiting.")

	// By default, sends and receives block until the other side is
	// ready. This allows goroutines to synchronize without explicit
	// locks or condition variables.

	fmt.Printf("Sending %d work packages to channel...\n", tasks)
	for i := 0; i < tasks; i++ {
		tasksCh <- i
	}
	fmt.Println("Finished sending work packages to channel.")

	close(tasksCh)
}

func main() {

	const workerCount = 1000000
	const taskCount = 20

	var wg sync.WaitGroup
	wg.Add(workerCount)
	go pool(&wg, workerCount, taskCount)
	wg.Wait()

	fmt.Println("All workers have finished.")
}
