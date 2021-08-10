package main

import "fmt"

func main() {
	arr := [10]int{8, 100, 50, 22, 15, 6, 1, 1000, 999, 0}
	bucketSort(arr[:])
}

func bucketSort(arr []int) {
	maxElement := arr[0]
	for _, i := range arr {
		if i > maxElement {
			maxElement = i
		}
	}

	bucket := make([]int, maxElement+1)

	for _, i := range arr {
		bucket[i]++
	}

	for indx, _ := range bucket {
		if bucket[indx] != 0 {
			for i := 0; i < bucket[indx]; i++ {
				fmt.Printf("%d ", indx)
			}
		}
	}
}
