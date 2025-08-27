package main

import (
	"flag"
	"fmt"
)

// Swapping values of two variables using XOR
func main() {
	fstPtr := flag.Int("fst", 0, "the value of the first number will be swapped with the value of the second number")
	sndPtr := flag.Int("snd", 0, "the value of the second number will be swapped with the value of the first number")
	flag.Parse()

	fst := *fstPtr
	snd := *sndPtr

	fmt.Println("Before swap\nfirst number:", fst, "second number:", snd, "\n")

	fst = fst ^ snd
	snd = fst ^ snd
	fst = fst ^ snd

	fmt.Println("After swap\nfirst number:", fst, "second number:", snd)
}
