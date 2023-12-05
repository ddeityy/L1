package main

import (
	"fmt"
	"sync"
)

/*
Реализовать структуру-счетчик, которая будет инкрементироваться в конкурентной среде.
По завершению программа должна выводить итоговое значение счетчика.
*/

type Counter interface {
	Add()
	Get() int64
}

var workers = 20
var iterations = 100

func startWorker(wg *sync.WaitGroup, counter Counter) {
	defer wg.Done()
	for i := 0; i < iterations; i++ {
		counter.Add()
	}
}

func main() {
	wg := &sync.WaitGroup{}
	mCounter := newMutexCounter()

	wg.Add(workers)
	for i := 0; i < workers; i++ {
		go startWorker(wg, mCounter)
	}

	wg.Wait()
	fmt.Printf("workers = %d, iterations = %d, counter = %d\n", workers, iterations, mCounter.Get())

	wg = &sync.WaitGroup{}
	aCounter := newAtomicCounter()
	wg.Add(workers)

	for i := 0; i < workers; i++ {
		go startWorker(wg, aCounter)
	}

	wg.Wait()
	fmt.Printf("workers = %d, iterations = %d, counter = %d\n", workers, iterations, aCounter.Get())
}
