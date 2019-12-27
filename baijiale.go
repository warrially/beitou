package main

import (
	"fmt"
	"math/rand"
	"time"
)

// TBaiJiaLe 百家乐
type TBaiJiaLe struct {
	cardList []int // 牌堆
	nIndex   int
}

// 洗牌
func (self *TBaiJiaLe) suffle() {
	// fmt.Println("\n洗牌")
	self.cardList = make([]int, 52*8)

	for i := 0; i < len(self.cardList); i++ {
		n := (i%13 + 1) // 得到 1 到 13

		// 10 以上都算0
		if n >= 10 {
			n = 0
		}

		self.cardList[i] = n
	}

	// 准备洗牌
	for i := 0; i < len(self.cardList); i++ {
		r := rand.Intn(len(self.cardList))
		self.cardList[i], self.cardList[r] = self.cardList[r], self.cardList[i]
	}

	self.nIndex = 0
}

// 百家乐开奖
func (self *TBaiJiaLe) run() (int, int) {
	nBanker := 0 // 庄家点数
	nPlayer := 0 // 闲家点数

	nPlayer += self.cardList[self.nIndex]
	self.nIndex++
	nBanker += self.cardList[self.nIndex]
	self.nIndex++

	nPlayer += self.cardList[self.nIndex]
	self.nIndex++
	nBanker += self.cardList[self.nIndex]
	self.nIndex++

	nPlayer %= 10
	nBanker %= 10

	if nBanker >= 8 || nPlayer >= 8 {
		return nBanker, nPlayer // 天牌
	}

	nBu := -1
	if nPlayer <= 5 {
		nBu = self.cardList[self.nIndex]
		self.nIndex++

		nPlayer += nBu // 补牌
	}
	// fmt.Print(" -bu- ", nBu, "                        ")
	switch nBanker {
	case 0, 1, 2:
		nBanker += self.cardList[self.nIndex]
		self.nIndex++
	case 3:
		if !(nBu == 8) {
			nBanker += self.cardList[self.nIndex]
			self.nIndex++
		}
	case 4:
		if !(nBu == 8 || nBu == 9 || nBu == 0 || nBu == 1) {
			nBanker += self.cardList[self.nIndex]
			self.nIndex++
		}
	case 5:
		if !(nBu == 0 || nBu == 1 || nBu == 2 || nBu == 3 || nBu == 8 || nBu == 9) {
			nBanker += self.cardList[self.nIndex]
			self.nIndex++
		}
	case 6:
		if nBu == 6 || nBu == 7 {
			nBanker += self.cardList[self.nIndex]
			self.nIndex++
		}
	}

	// if self.nIndex > len(self.cardList)/2 {
	// 	self.suffle() // 切牌洗
	// }
	return nBanker % 10, nPlayer % 10 // 天牌
}

func main1() {
	rand.Seed(time.Now().UnixNano())

	nBankerWin := 0
	nPlayerWin := 0
	nSame := 0

	nRobot := 0 // 带入10000元
	nTotalBet := 0
	for {
		bjl := &TBaiJiaLe{}
		bjl.suffle()

		for i := 0; i < 48; i++ {
			nBetPos := rand.Intn(2)
			nBetMoney := rand.Intn(20) * 100 // 随机下注20元
			nTotalBet += nBetMoney

			nBanker, nPlayer := bjl.run()
			if nBanker < nPlayer {
				if nBetPos == 0 {
					nRobot += nBetMoney // 闲赢
				} else {
					nRobot -= nBetMoney // 闲输
				}
				nPlayerWin++
			} else if nBanker > nPlayer {
				if nBetPos == 1 {
					nRobot += (nBetMoney * 95 / 100) // 庄赢
				} else {
					nRobot -= nBetMoney // 庄输
				}
				nBankerWin++
			} else {
				nSame++
			}
			// nAll := nBankerWin + nPlayerWin + nSame
			// fmt.Print("\r", nBankerWin*1000000/nAll, " - ", nPlayerWin*1000000/nAll, " - ", nSame*1000000/nAll, "                            ")
		}

		fmt.Println("平台盈利", -nRobot, "总下注额", nTotalBet, "盈利率", -nRobot*10000/nTotalBet)

		time.Sleep(11 * time.Millisecond)
	}

}
