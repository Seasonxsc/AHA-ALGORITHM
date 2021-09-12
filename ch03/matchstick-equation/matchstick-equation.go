package main

import "fmt"

func main() {
	c, m, sum := 0, 0, 0
	fmt.Printf("请输入火柴棍的根数: ")
	fmt.Scanln(&m)

	for a := 0; a <= 11111; a++ {
		for b := 0; b <= 11111; b++ {
			c = a + b

			if needNum(a)+needNum(b)+needNum(c) == m-4 {
				fmt.Printf("%d + %d = %d\n", a, b, c)
				sum++
			}

		}
	}

	fmt.Printf("一共可以拼出%d个不同的等式", sum)
}

func needNum(n int) int {
	num := 0
	f := [10]int{6, 2, 5, 5, 4, 5, 6, 3, 7, 6}

	for n/10 != 0 {
		num += f[n%10]
		n /= 10
	}

	num += f[n]

	return num
}
