package main

/*
	Реализовать конкурентную запись данных в map.
*/

import (
	"fmt"
	"math/rand"
	"sync"
)

type ConcurrentMap[T any] struct {
	sync.RWMutex
	data map[string]T
}

func NewConcurrentMap[T any]() *ConcurrentMap[T] {
	return &ConcurrentMap[T]{
		data: make(map[string]T),
	}
}

func (m *ConcurrentMap[T]) Get(key string) (T, bool) {
	m.RLock()
	defer m.RUnlock()
	v, ok := m.data[key]
	return v, ok
}

func (m *ConcurrentMap[T]) Set(key string, val T) {
	m.Lock()
	defer m.Unlock()
	m.data[key] = val
}

func (m *ConcurrentMap[T]) Has(key string) bool {
	m.RLock()
	defer m.RUnlock()
	_, ok := m.data[key]
	return ok
}

func (m *ConcurrentMap[T]) Delete(key string) {
	m.Lock()
	defer m.Unlock()
	delete(m.data, key)
}

func main() {
	m := NewConcurrentMap[int]()
	wg := sync.WaitGroup{}
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(m *ConcurrentMap[int], index int) {
			key := fmt.Sprintf("test%v", index)
			val := rand.Intn(10)
			m.Set(key, val)
			fmt.Println("goroutine", index, "set", key, "to", val)
			wg.Done()
		}(m, i)
	}
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(m *ConcurrentMap[int], index int) {
			key := fmt.Sprintf("test%v", index)
			val, ok := m.Get(key)
			if !ok {
				fmt.Println(key, "not found")
			}
			fmt.Println("Key", key, "value:", val)
			wg.Done()
		}(m, i)
	}
	wg.Wait()

}
