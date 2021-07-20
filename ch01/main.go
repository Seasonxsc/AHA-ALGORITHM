package main

import (
	"fmt"
)

type Student struct {
	name  string
	score int
}

type Students []Student

func (s Students) Len() int {
	return len(s)
}

func (s Students) Less(i, j int) bool {
	return s[i].score < s[j].score
}

func (s Students) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func main() {
	arr := [9]int{6, 1, 2, 7, 9, 4, 5, 10, 8}

	// bucketSort(arr[0:])
	bubbleSort(arr[0:])

	// students := Students{{"huhu", 5}, {"haha", 3}, {"xixi", 5}, {"hengheng", 2}, {"gaoshou", 8}}
	// studentsSort(students)
	// sort.Sort(students)
	// quickSort(arr[0:], 0, len(arr)-1)
	// fmt.Println(arr)
}

// 桶排序
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

// 冒泡排序
func bubbleSort(arr []int) {
	len := len(arr)

	for e := len - 1; e > 0; e-- {
		for i := 0; i < e; i++ {
			if arr[i] > arr[i+1] {
				arr[i], arr[i+1] = arr[i+1], arr[i]
			}
		}
	}
	fmt.Println(arr)
}

// 类数组排序01
func studentsSort(students []Student) {
	var tmp Student
	len := len(students)

	for i := 0; i < len-1; i++ {
		for j := 0; j < len-i-1; j++ {
			if students[j+1].score > students[j].score {
				tmp = students[j]
				students[j] = students[j+1]
				students[j+1] = tmp
			}
		}
	}

	for _, i := range students {
		fmt.Printf("%s ", i.name)
	}
}

// 快速排序
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
