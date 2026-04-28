package worker

import (
	"errors"
	"sync"
)

type Job func()

type WorkerPool struct {
	jobQueue chan Job
	wg       sync.WaitGroup
	once     sync.Once
}

func NewWorkerPool(size int) *WorkerPool {
	wp := &WorkerPool{
		jobQueue: make(chan Job, 100),
	}

	for i := 0; i < size; i++ {
		go wp.startWorker(i)
	}
	return wp
}

func (wp *WorkerPool) startWorker(id int) {
	for job := range wp.jobQueue {
		if job == nil {
			wp.wg.Done()
			continue
		}

		func() {
			defer func() {
				if r := recover(); r != nil {
					println("Worker panic recovered")
				}
				wp.wg.Done()
			}()
			job()
		}()
	}
}

func (wp *WorkerPool) Submit(job Job) error {
	if job == nil {
		return errors.New("job cannot be nil")
	}

	wp.wg.Add(1)
	wp.jobQueue <- job
	return nil
}

func (wp *WorkerPool) Wait() {
	wp.wg.Wait()
}

func (wp *WorkerPool) Stop() {
	wp.once.Do(func() {
		close(wp.jobQueue)
	})
}
