package main

import "sync"

type mutexCounter struct {
	counter int64
	mutex   *sync.RWMutex
}

func newMutexCounter() *mutexCounter {
	return &mutexCounter{
		counter: 0,
		mutex:   &sync.RWMutex{},
	}
}

func (m *mutexCounter) Add() {
	m.mutex.Lock()
	defer m.mutex.Unlock()
	m.counter++
}

func (m *mutexCounter) Get() int64 {
	m.mutex.RLock()
	defer m.mutex.RUnlock()
	return m.counter
}
