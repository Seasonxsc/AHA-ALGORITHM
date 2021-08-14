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
	sort.Sort(students)
	fmt.Println(students)
}

// 方法1
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

// 方法二
func (s Students) Len() int {
	return len(s)
}

func (s Students) Less(i, j int) bool {
	return s[i].score < s[j].score
}

func (s Students) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
