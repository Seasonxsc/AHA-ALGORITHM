package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

var (
	xiaoHeng []string // 0
	xiaoHa   []string // 1
	poker    []string
	book     map[string]int
	desktop  []string
)

func main() {
	rand.Seed(time.Now().UnixNano())
	initPoker()
	var n int // 发牌数
	fmt.Printf("请选择发几张牌: ")
	fmt.Scanln(&n)
	deal(n)

	if winner := play(rand.Intn(2)); winner == 0 {
		fmt.Println("小哼获胜！")
	} else {
		fmt.Println("小哈获胜！")
	}
}

func initPoker() {
	for i := 0; i < 2; i++ {
		poker = append(poker, "joker")
	}

	for i := 0; i < 4; i++ {
		poker = append(poker, "ace")
		poker = append(poker, "jack")
		poker = append(poker, "queen")
		poker = append(poker, "king")
	}

	for i := 2; i <= 10; i++ {
		p := strconv.Itoa(i)
		for j := 0; j < 4; j++ {
			poker = append(poker, p)
		}
	}
	rand.Shuffle(len(poker), func(i, j int) { poker[i], poker[j] = poker[j], poker[i] })
}

func deal(n int) {
	for i := 0; i < n; i++ {
		xiaoHa = append(xiaoHa, poker[len(poker)-1])
		poker = poker[:len(poker)-1]
		xiaoHeng = append(xiaoHeng, poker[len(poker)-1])
		poker = poker[:len(poker)-1]
	}
	book = make(map[string]int)
}

func play(player int) int {
	count := 1
	if player == 0 {
		fmt.Println("小哼先出牌")
		for len(xiaoHeng) != 0 && len(xiaoHa) != 0 {
			fmt.Printf("第%d轮:\n", count)
			fmt.Printf("小哼当前手中的牌是: %v\n", xiaoHeng)
			fmt.Printf("小哈当前手中的牌是: %v\n", xiaoHa)
			// 小哼出牌
			fmt.Printf("小哼打出: ")
			turn(&xiaoHeng)
			// 小哈出牌
			fmt.Printf("小哈打出: ")
			turn(&xiaoHa)
			fmt.Printf("当前桌面上的牌: %v\n", desktop)

			count++
		}
	} else {
		fmt.Println("小哈先出牌")
		for len(xiaoHeng) != 0 && len(xiaoHa) != 0 {
			fmt.Printf("第%d轮:\n", count)
			fmt.Printf("小哈当前手中的牌是: %v\n", xiaoHa)
			fmt.Printf("小哼当前手中的牌是: %v\n", xiaoHeng)
			// 小哈出牌
			fmt.Printf("小哈打出: ")
			turn(&xiaoHa)
			// 小哼出牌
			fmt.Printf("小哼打出: ")
			turn(&xiaoHeng)
			fmt.Printf("当前桌面上的牌: %v\n", desktop)

			count++
		}
	}

	if len(xiaoHeng) == 0 {
		return 1
	} else {
		return 0
	}
}

func turn(player *[]string) {
	face := (*player)[len(*player)-1]
	fmt.Println(face)
	*player = (*player)[:len(*player)-1]
	desktop = append(desktop, face)
	check(face, player)
}

func check(face string, player *[]string) {
	if book[face] != 0 {
		for i, _ := range desktop {
			if desktop[i] == face {
				collector := []string{}
				collector = append(collector, desktop[i:]...)
				for j, _ := range collector {
					if j == 0 {
						continue
					}
					book[collector[j]]--
				}
				collector = append(collector, *player...)
				*player, desktop = collector, desktop[:i]
				break
			}
		}
	} else {
		book[face]++
	}
}
