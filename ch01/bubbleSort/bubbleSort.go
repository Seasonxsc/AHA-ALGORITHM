package main

import "fmt"

func main() {
	arr := [10]int{8, 100, 50, 22, 15, 6, 1, 1000, 999, 0}
	fmt.Println(arr)
}

func bubbleSort(arr []int) {
	length := len(arr)
	if arr == nil || length < 2 {
		return
	}

	for e := length - 1; e >= 1; e-- {
		for i := 0; i < e; i++ {
			if arr[i] > arr[i+1] {
				arr[i], arr[i+1] = arr[i+1], arr[i]
			}
		}
	}
}
