package main

/*
Реализовать собственную функцию sleep.
*/

import (
	"fmt"
	"time"
)

func sleep(duration time.Duration) {
	<-time.After(duration)
}

func main() {
	start := time.Now()
	sleep(5 * time.Second)
	fmt.Println(time.Since(start))
}
