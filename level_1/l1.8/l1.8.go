package main

import (
	"fmt"
)

func setOneOrZeroBit(n int64, i int64) int64 {
	mask := int64(1) << i
	if (n>>i)&1 == 0 {
		return (n | mask)
	} else {
		return (n & (^mask))
	}
}

func main() {
	var n int64
	var i int64

	fmt.Print("Entere the integer and the position (index 0-64) to change the bit: ")
	fmt.Scanf("%v %v", &n, &i)

	fmt.Println(setOneOrZeroBit(n, i))
}
