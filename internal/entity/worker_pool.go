package entity

import (
	"fmt"
	"sync"
)

type WorkerPool struct {
	Files       []ImageFile
	concurrency int
	filesChan   chan ImageFile
	wg          sync.WaitGroup
}

func NewWorkerPool(files []ImageFile, concurrency int) *WorkerPool {
	return &WorkerPool{
		Files:       files,
		concurrency: concurrency,
		filesChan:   make(chan ImageFile, concurrency),
		wg:          sync.WaitGroup{},
	}
}

func (wp *WorkerPool) Start() {
	fmt.Println("Starting worker pool")

	// Add the number of workers to the WaitGroup
	wp.wg.Add(len(wp.Files))

	// Start the workers
	for i := 0; i < wp.concurrency; i++ {
		fmt.Println("Starting worker: ", i+1)

		go wp.worker()
	}

	// Add the files to the channel
	for _, file := range wp.Files {
		wp.filesChan <- file
	}

	// Close the channel
	close(wp.filesChan)

	// Wait for all workers to finish
	wp.wg.Wait()

	fmt.Println("All workers have finished")
}

func (wp *WorkerPool) worker() {
	for file := range wp.filesChan {
		// Process the file
		file.Process()

		// Mark the worker as done
		wp.wg.Done()
	}
}
