package main

import "fmt"

// SortSlice sorts a slice of integers without using the sort package
func SortSlice(slice []int) {
	for i := 0; i < len(slice)-1; i++ {
		for j := 0; j < len(slice)-i-1; j++ {
			if slice[j] > slice[j+1] {
				slice[j], slice[j+1] = slice[j+1], slice[j]
			}
		}
	}
}

// IncrementOdd increments integers by one in odd positions
func IncrementOdd(slice []int) {
	for i := 1; i < len(slice); i += 2 {
		slice[i]++
	}
}

// PrintSlice prints the slice
func PrintSlice(slice []int) {
	fmt.Println(slice)
}

// ReverseSlice reverses the slice
func ReverseSlice(slice []int) {
	for i := 0; i < len(slice)/2; i++ {
		j := len(slice) - i - 1
		slice[i], slice[j] = slice[j], slice[i]
	}
}

// appendFunc appends processing functions to a destination function
func appendFunc(dst func([]int), src ...func([]int)) func([]int) {
	return func(slice []int) {
		dst(slice) // Call the destination function first
		for _, fn := range src {
			fn(slice) // Call each additional function
		}
	}
}

func main() {
	// Test the functions
	numbers := []int{5, 3, 9, 1, 7}
	fmt.Println("Original slice:")
	PrintSlice(numbers)

	combinedFunc := appendFunc(SortSlice, IncrementOdd, ReverseSlice)
	combinedFunc(numbers)
	fmt.Println("Combined functions applied:")
	PrintSlice(numbers)
}
