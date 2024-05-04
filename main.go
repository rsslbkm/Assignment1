package main

import "fmt"

func SortSlice(slice []int) {
	for i := 0; i < len(slice)-1; i++ {
		for j := 0; j < len(slice)-i-1; j++ {
			if slice[j] > slice[j+1] {
				slice[j], slice[j+1] = slice[j+1], slice[j]
			}
		}
	}
}

func IncrementOdd(slice []int) {
	for i := 1; i < len(slice); i += 2 {
		slice[i]++
	}
}

func PrintSlice(slice []int) {
	fmt.Println(slice)
}

func ReverseSlice(slice []int) {
	for i := 0; i < len(slice)/2; i++ {
		j := len(slice) - i - 1
		slice[i], slice[j] = slice[j], slice[i]
	}
}

func appendFunc(dst func([]int), src ...func([]int)) func([]int) {
	return func(slice []int) {
		dst(slice)
		for _, fn := range src {
			fn(slice)
		}
	}
}

func main() {
	numbers := []int{9, 3, 5, 7, 1}
	fmt.Println("Original slice:")
	PrintSlice(numbers)

	combinedFunc := appendFunc(SortSlice, IncrementOdd, ReverseSlice)
	combinedFunc(numbers)
	fmt.Println("Combined functions applied:")
	PrintSlice(numbers)
}
