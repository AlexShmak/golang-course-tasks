package main

import "fmt"

func deleteSliceElByIndex(i int, s []int) []int {
	copy(s[i:], s[i+1:])
	n := len(s)
	return s[:n-1]
}

func main() {
	testSlice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(testSlice)

	testSlice = deleteSliceElByIndex(5, testSlice) // delete number at index 5 (number 6)
	fmt.Println(testSlice)
}
