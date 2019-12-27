package main

import (
	"fmt"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())

}

func main() {
	pRobot1 := NewRobot1()
	for {
		if !pRobot1.Bet(0) {
			fmt.Println(pRobot1.nTotalBet, pRobot1.nMoney, pRobot1.nMax)
			break
		}
		pRobot1.Calc(open())

	}
	main()
}

func main2() {

	nSystem := Cash(0) // 系统的金额
	nBig := 0
	nSmall := 0
	// 倍投的机器人AI
	pRobot1 := NewRobot1()
	// pRobot2 := NewRobot2()

	// 循环开奖中
	for {

		if nBig < nSmall {
			// 只有在在概率偏差时才下注
			if !pRobot1.Bet(0) {
				break
			}
		} else {
			if !pRobot1.Bet(1) {
				break
			}
		}

		nOpen := open()
		nSystem -= pRobot1.Calc(nOpen)

		switch nOpen {
		case 0: // 小
			nBig++
		case 1: // 大
			nSmall++
		}
	}
	fmt.Println(nBig, nSmall, pRobot1.nMax)
	time.Sleep(10 * time.Millisecond)
	main()
}

// 开奖
func open() int {
	n := 20
	r := rand.Intn(n + 1)

	if r == n {
		return 2
	}

	r = r & 1
	if r == 0 {
		return 0
	}
	return 1
}
