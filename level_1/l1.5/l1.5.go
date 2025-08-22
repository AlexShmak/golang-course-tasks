package main

import (
	"context"
	"flag"
	"time"
)

func main() {

	nPtr := flag.Int("n", 5, "number of seconds before timeout")
	flag.Parse()
	n := *nPtr

	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(n)*time.Second)
	defer cancel()

	channel := make(chan int)

	go sender(ctx, channel)

	for {
		select {
		case output := <-channel:
			println(output)
		case <-ctx.Done():
			println("timeout")
			return
		}
	}

}

func sender(ctx context.Context, ch chan<- int) {
	defer close(ch)

	i := 0
	for {
		select {
		case <-ctx.Done():
			return
		case ch <- i:
			i++
		}
	}
}
