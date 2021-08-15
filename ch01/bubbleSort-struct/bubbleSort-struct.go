package main

import (
	"fmt"
	"sort"
)

type Student struct {
	name  string
	score int
}

type Students []Student

func main() {
	students := Students{{"huhu", 5}, {"haha", 3}, {"xixi", 5}, {"hengheng", 2}, {"gaoshou", 8}}
	// 方法一
	// studentsSort(students)
	// 方法二
	sort.Slice(students, func(i, j int) bool {
		return students[i].score < students[j].score
	})
	fmt.Println(students)
}

func studentsSort(students Students) {
	length := len(students)
	if students == nil || length < 2 {
		return
	}

	for e := length - 1; e >= 1; e-- {
		for i := 0; i < e; i++ {
			if students[i+1].score > students[i].score {
				students[i], students[i+1] = students[i+1], students[i]
			}
		}
	}
}
