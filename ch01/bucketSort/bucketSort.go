package main

import "fmt"

func main() {
	arr := [10]int{8, 100, 50, 22, 15, 6, 1, 1000, 999, 0}
	bucketSort(arr[:])
	fmt.Println(arr)

}

func bucketSort(arr []int) {
	max := arr[0]
	for _, v := range arr {
		if v > max {
			max = v
		}
	}

	bucket := make([]int, max+1)

	for _, v := range arr {
		bucket[v]++
	}

	indx := 0
	for i, _ := range bucket {
		if bucket[i] != 0 {
			for j := 0; j < bucket[j]; j++ {
				arr[indx] = i
				indx++
			}
		}
	}
}
