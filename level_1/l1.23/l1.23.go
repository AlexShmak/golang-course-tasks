package main

import "fmt"

func deleteSliceElByIndex[T any](i int, s []T) []T {
	n := len(s)
	if i < 0 || i >= n {
		return s
	}
	return append(s[:i], s[i+1:]...)
}

func main() {
	intSlice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(intSlice, len(intSlice))

	intSlice = deleteSliceElByIndex(5, intSlice) // delete number at index 5 (number 6)
	fmt.Println(intSlice, len(intSlice))

	stringSlice := []string{"foo", "bar", "foobar"}
	fmt.Println(stringSlice, len(stringSlice))

	stringSlice = deleteSliceElByIndex(2, stringSlice) // delete string at index 5 ("foobar")
	fmt.Println(stringSlice, len(stringSlice))
}
