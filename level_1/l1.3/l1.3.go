package main

import (
	"flag"
	"fmt"
)

func worker(ch chan int) {
	for val := range ch {
		fmt.Println(val)
	}
}

func write(ch chan int) {
	for i := 0; ; i++ {
		ch <- i
		fmt.Println("wrote", i)
	}
}

func main() {
	var workersPtr = flag.Int("n", 5, "number of worker goroutines")
	flag.Parse()
	workers := *workersPtr
	ch := make(chan int)

	go write(ch)

	for range workers {
		go worker(ch)
	}

	select {}
}
