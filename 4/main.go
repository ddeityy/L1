package main

/*
	Реализовать постоянную запись данных в канал (главный поток).
	Реализовать набор из N воркеров, которые читают произвольные данные из канала и выводят в stdout.
	Необходима возможность выбора количества воркеров при старте.
	Программа должна завершаться по нажатию Ctrl+C. Выбрать и обосновать способ завершения работы всех воркеров.
*/

import (
	"context"
	"flag"
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

// Воркер выводит число из канала
func startWorker(ctx context.Context, wg *sync.WaitGroup, id int, in <-chan int) {
	defer wg.Done()
	var number int
	for {
		select {
		case number = <-in:
			fmt.Printf("Worker %d: %d\n", id, number)
		case <-ctx.Done():
			fmt.Printf("Shutting down worker %d\n", id)
			return
		}
	}
}

func main() {
	wg := &sync.WaitGroup{}
	input := make(chan int)

	var delay time.Duration
	var workers int
	fl := flag.NewFlagSet("worker-pool", flag.ContinueOnError)
	fl.DurationVar(&delay, "d", 500*time.Millisecond, "message delay")
	fl.IntVar(&workers, "w", 4, "number of workers")

	if err := fl.Parse(os.Args[1:]); err != nil {
		panic(err)
	}

	wg.Add(workers)

	defer close(input)
	ctx, cancel := context.WithCancel(context.Background())

	// запускаем n воркеров
	for i := 0; i < workers; i++ {
		go startWorker(ctx, wg, i+1, input)
	}

	go func() {
		wg.Add(1)
		defer wg.Done()

		ticker := time.NewTicker(delay)
		defer ticker.Stop()

		for {
			select {
			case <-ticker.C:
				a := rand.Int()
				input <- a
				fmt.Printf("Sending: %d\n", a)
			case <-ctx.Done():
				fmt.Println("Shutting down sender")
				return
			}
		}
	}()

	// по нажатию ctrl+c отправляем сигнал завершения всем горутинам через контекст
	// и ждем их завершения
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)
	<-stop
	fmt.Println("\ngot interrupt signal")
	cancel()
	wg.Wait()

}
