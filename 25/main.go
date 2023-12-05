package main

import (
	"fmt"
	"time"
)

/*
Реализовать собственную функцию sleep.
*/

func sleep(duration time.Duration) {
	<-time.After(duration)
}

func main() {
	start := time.Now()
	sleep(5 * time.Second)
	fmt.Println(time.Since(start))
}
