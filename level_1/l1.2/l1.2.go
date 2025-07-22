package main

import (
	"fmt"
	"sync"
)

func calculateSquare(input int, wg *sync.WaitGroup) {
	defer wg.Done()
	square := input * input
	fmt.Println(square)

}

func main() {
	input := []int{2, 4, 6, 8, 10}
	var wg sync.WaitGroup

	for _, num := range input {
		wg.Add(1)
		go calculateSquare(num, &wg)
	}
	wg.Wait()
}
