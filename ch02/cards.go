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
	deal()
	winner := play(rand.Intn(2))

	if winner == 0 {
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

func deal() {
	for len(poker) != 0 {
		xiaoHa = append(xiaoHa, poker[len(poker)-1])
		poker = poker[:len(poker)-1]
		xiaoHeng = append(xiaoHeng, poker[len(poker)-1])
		poker = poker[:len(poker)-1]
	}
	book = make(map[string]int)
}

func play(player int) int {
	if player == 0 {
		for len(xiaoHeng) != 0 && len(xiaoHa) != 0 {
			// 小哼出牌
			turn(&xiaoHeng)
			// 小哈出牌
			turn(&xiaoHa)
		}
	} else {
		for len(xiaoHeng) != 0 && len(xiaoHa) != 0 {
			// 小哈出牌
			turn(&xiaoHa)
			// 小哼出牌
			turn(&xiaoHeng)
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
	*player = (*player)[:len(*player)-1]
	desktop = append(desktop, face)
	check(face, player)
}

func check(face string, player *[]string) {
	if book[face] != 0 {
		for i, _ := range desktop {
			if desktop[i] == face {
				collection := []string{}
				collection = append(collection, desktop[i:]...)
				for j, _ := range collection {
					if j == 0 {
						continue
					}
					book[collection[j]]--
				}
				collection = append(collection, *player...)
				*player = collection
				desktop = desktop[:i]
				break
			}
		}
	} else {
		book[face]++
	}
}
