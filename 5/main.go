package main

import (
	"flag"
	"fmt"
	"math/rand"
	"os"
	"sync"
	"time"
)

/*
	Разработать программу, которая будет последовательно отправлять значения в канал,
	а с другой стороны канала — читать. По истечению N секунд программа должна завершаться.
*/

// читает данные из канала и выводит их
func read(wg *sync.WaitGroup, in <-chan int) {
	defer wg.Done()
	for v := range in {
		fmt.Printf("got %d\n", v)
	}
}

func main() {
	wg := &sync.WaitGroup{}
	var duration time.Duration
	input := make(chan int, 1)

	fl := flag.NewFlagSet("duration", flag.ContinueOnError)
	fl.DurationVar(&duration, "d", 10*time.Second, "number of seconds to run")

	if err := fl.Parse(os.Args[1:]); err != nil {
		panic(err)
	}

	wg.Add(1)
	go read(wg, input)

	// отправляет в канал случайные числа,
	// по истечении определенного времени завершается
	go func() {
		ticker := time.NewTicker(time.Second)
		defer ticker.Stop()
		timer := time.NewTimer(duration)
		defer timer.Stop()

		for {
			select {
			case <-ticker.C:
				n := rand.Int()
				input <- n
				fmt.Println("Sending:", n)
			case <-timer.C:
				fmt.Println("Timeout, exiting")
				close(input)
				return
			}
		}
	}()

	wg.Wait()
}
