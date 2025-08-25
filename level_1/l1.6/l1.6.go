package main

import (
	"context"
	"fmt"
	"runtime"
	"sync"
	"time"
)

func stopByCondition() {
	fmt.Println("= stop by condition demo =")
	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer wg.Done()
		for i := range 3 {
			fmt.Println("goroutine: working, step", i+1)
			time.Sleep(400 * time.Millisecond)
		}
		fmt.Println("goroutine: done\n")
	}()

	wg.Wait()
}

func stopWithNotificationChannel() {
	fmt.Println("= stop with notification channel demo =")
	done := make(chan bool)
	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer wg.Done()
		for {
			select {
			case <-done:
				fmt.Println("goroutine (channel): received signal, stopping work...\n")
				return
			default:
				fmt.Println("goroutine (channel): working...")
				time.Sleep(500 * time.Millisecond)
			}
		}
	}()

	time.Sleep(2 * time.Second)

	fmt.Println("sending stop signal to goroutine...")
	done <- true

	wg.Wait()
}

func stopWithContext() {
	fmt.Println("= stop with context demo =")
	ctx, cancel := context.WithCancel(context.Background())
	var wg sync.WaitGroup
	wg.Add(1)

	go func(ctx context.Context) {
		defer wg.Done()
		for {
			select {
			case <-ctx.Done():
				fmt.Println("goroutine (context): context canceled, stopping work.\n")
				return
			default:
				fmt.Println("goroutine (context): working...")
				time.Sleep(500 * time.Millisecond)
			}
		}
	}(ctx)

	time.Sleep(2 * time.Second)

	fmt.Println("canceling context...")
	cancel()

	wg.Wait()
}

func stopWithGoexit() {
	fmt.Println("= stop with Goexit demo =")
	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer func() {
			fmt.Println("goroutine (Goexit): deferred call executed.")
			wg.Done()
		}()

		fmt.Println("goroutine (Goexit): starting work...")
		time.Sleep(1 * time.Second)

		fmt.Println("goroutine (Goexit): calling runtime.Goexit().")
		runtime.Goexit()

		fmt.Println("goroutine (Goexit): this line is unreachable.")
	}()

	wg.Wait()
	fmt.Println("goroutine was stopped with Goexit.\n")
}

func stopByClosingChannel() {
	fmt.Println("= stop by closing channel demo =")
	dataChan := make(chan int)
	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer wg.Done()
		for item := range dataChan {
			fmt.Printf("received: %d\n", item)
			time.Sleep(500 * time.Millisecond)
		}
		fmt.Println("channel closed, goroutine exiting.")
	}()

	fmt.Println("sending data to channel...")
	dataChan <- 1
	dataChan <- 2
	dataChan <- 3
	close(dataChan)
	fmt.Println("all data sent, channel closed.")

	wg.Wait()
}

func main() {
	stopByCondition()
	stopWithNotificationChannel()
	stopWithContext()
	stopWithGoexit()
	stopByClosingChannel()
}
