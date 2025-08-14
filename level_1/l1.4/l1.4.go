package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

func worker(ctx context.Context, ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case val, ok := <-ch:
			if !ok {
				return
			}
			fmt.Println(val)
		case <-ctx.Done():
			return
		}
	}
}

func write(ctx context.Context, ch chan int) {
	defer close(ch)
	for i := 0; ; i++ {
		select {
		case <-ctx.Done():
			fmt.Println("context done, stopping write")
			return
		case ch <- i:
			fmt.Println("wrote", i)
		}
	}
}

func main() {
	workersPtr := flag.Int("n", 5, "number of worker goroutines")
	flag.Parse()
	workers := *workersPtr

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	ch := make(chan int)
	var wg sync.WaitGroup
	wg.Add(workers)

	for range workers {
		go worker(ctx, ch, &wg)
	}
	go write(ctx, ch)

	<-ctx.Done()
	fmt.Println("received shutdown signal, exiting...")
	wg.Wait()
	fmt.Println("all workers finished")
}

// Ctrl+C обрабатывается через signal.NotifyContext: SIGINT/SIGTERM отменяют контекст (ctx.Done()).
// Writer по отмене выходит и закрывает канал; worker'ы завершаются либо по ctx.Done(), либо по закрытию канала; main ждёт wg.
