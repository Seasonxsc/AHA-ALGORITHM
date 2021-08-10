package main

import "fmt"

func main() {
	arr := [10]int{6, 1, 2, 7, 9, 3, 4, 5, 10, 8}
	quickSort(arr[:], 0, len(arr)-1)
	fmt.Println(arr)
}

func quickSort(arr []int, left, right int) {
	if left > right {
		return
	}

	pivot := arr[left]
	i := left
	j := right

	for i != j {
		for arr[j] >= pivot && i < j {
			j--
		}

		for arr[i] <= pivot && i < j {
			i++
		}

		if i < j {
			arr[i], arr[j] = arr[j], arr[i]
		}
	}

	arr[left] = arr[i]
	arr[i] = pivot
	quickSort(arr, left, i-1)
	quickSort(arr, i+1, right)
}
