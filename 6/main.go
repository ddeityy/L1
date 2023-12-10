package main

/*
	Реализовать все возможные способы остановки выполнения горутины.
*/

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type goroutine struct {
	wg *sync.WaitGroup
}

// используем проверку закрытия канала через !ok для завершения горутины
func (g *goroutine) channelClosed(ch <-chan int) {
	defer g.wg.Done()
	for {
		v, ok := <-ch
		if !ok {
			fmt.Println("channel closed")
			return
		}
		fmt.Println("Got:", v)
	}
}

// итерация по range завершится сама, как только закроется канал
func (g *goroutine) channelClosedRange(ch <-chan int) {
	defer g.wg.Done()
	for v := range ch {
		fmt.Println("Got:", v)
	}
	fmt.Println("Channel closed, range loop is finished")
}

// используем stop канал для завершения горутины
func (g *goroutine) stopChannel(ch <-chan int, stop <-chan struct{}) {
	defer g.wg.Done()
	for {
		select {
		case v := <-ch:
			fmt.Println("Got:", v)
		case <-stop:
			fmt.Println("Stop channel")
			return
		}
	}
}

// используем сигнал из контекста для завершения горутины
func (g *goroutine) contextWithCancel(ctx context.Context, ch <-chan int) {
	defer g.wg.Done()
	for {
		select {
		case v := <-ch:
			fmt.Println("Got:", v)
		case <-ctx.Done():
			fmt.Println("ctx.Done received")
			return
		}
	}
}

// используем таймаунт контекста для завершения горутины
func (g *goroutine) contextWithTimeout(ctx context.Context, ch <-chan int) {
	defer g.wg.Done()
	for {
		select {
		case <-time.After(10 * time.Second):
			fmt.Println("Got:", <-ch)
		case <-ctx.Done():
			fmt.Println("Timeout")
			return
		}
	}
}

// запись случайных чисел в канал
func (g *goroutine) send(ctx context.Context, ch chan int) {
	defer g.wg.Done()

	ticker := time.NewTicker(time.Second)

	for {
		select {
		case <-ticker.C:
			n := rand.Intn(10)
			ch <- n
			fmt.Println("Sending:", n)
		case <-ctx.Done():
			fmt.Println("Finished sending")
			ticker.Stop()
			return
		}
	}
}

func main() {
	g := goroutine{
		wg: &sync.WaitGroup{},
	}

	// context with timeout
	fmt.Println("\nContext timeout")
	ctx, cancel := context.WithCancel(context.Background())
	timeout, timeoutCancel := context.WithTimeout(context.Background(), 5*time.Second)
	ch := make(chan int, 1)

	g.wg.Add(1)
	go g.send(ctx, ch)
	go g.contextWithTimeout(timeout, ch)

	g.wg.Wait()
	cancel()
	timeoutCancel()

	// context with cancel
	fmt.Println("\nContext cancel")
	ctx, cancel = context.WithCancel(context.Background())
	ch = make(chan int, 1)

	g.wg = &sync.WaitGroup{}
	g.wg.Add(2)
	go g.send(ctx, ch)
	go g.contextWithCancel(ctx, ch)

	time.Sleep(3 * time.Second)
	cancel()
	close(ch)
	g.wg.Wait()

	// separate stop channel
	fmt.Println("\nStop channel")
	ch = make(chan int, 1)
	stop := make(chan struct{})
	ctx, cancel = context.WithCancel(context.Background())

	g.wg = &sync.WaitGroup{}
	g.wg.Add(2)
	go g.send(ctx, ch)
	go g.stopChannel(ch, stop)

	time.Sleep(3 * time.Second)
	stop <- struct{}{}
	cancel()
	close(ch)
	g.wg.Wait()

	// Range over channel
	fmt.Println("\nRange over a channel")
	ch = make(chan int, 1)
	ctx, cancel = context.WithCancel(context.Background())

	g.wg = &sync.WaitGroup{}
	g.wg.Add(2)
	go g.send(ctx, ch)
	go g.channelClosedRange(ch)

	time.Sleep(3 * time.Second)
	close(ch)
	cancel()
	g.wg.Wait()

	// Close channel
	fmt.Println("\nClosing a channel")
	ch = make(chan int, 1)
	ctx, cancel = context.WithCancel(context.Background())

	g.wg = &sync.WaitGroup{}
	g.wg.Add(2)
	go g.send(ctx, ch)
	go g.channelClosed(ch)

	time.Sleep(3 * time.Second)
	close(ch)
	cancel()
	g.wg.Wait()
}
