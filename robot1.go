package main

import "fmt"

// Cash 金额
type Cash int64

func (c Cash) String() string {
	return fmt.Sprintf("%.2f", float64(c)/100.0)
}

// TRobot1 机器人1  倍投机器人
type TRobot1 struct {
	nMoney  Cash
	nTakein Cash
	nBet    Cash
	nBetPos int  // 下注位置
	bBet    bool // 是否下注
	bBroken bool // 是否破产

	nTotalBet Cash
	nMax      Cash // 最大的钱数
}

// Bet 下注
func (self *TRobot1) Bet(nBetPos int) bool {
	if self.bBroken {
		return false
	}

	if self.bBet {
		return false
	}

	if self.nMoney <= 100 {
		// fmt.Println("\n破产")
		// fmt.Println("最大拥有金额是", self.nMax)
		self.bBroken = true
		// 破产了
		return false
	}

	self.nBetPos = nBetPos

	// 钱不够翻倍的时候
	if self.nMoney < self.nBet {
		self.nBet = self.nMoney
	}

	if self.nBet > 100000 {
		self.nBet = 100000 // 限制
	}

	self.bBet = true
	self.nMoney -= self.nBet    // 扣下注额度
	self.nTotalBet += self.nBet // 总下注
	return true
}

// 计算结果
func (self *TRobot1) Calc(nBetPos int) Cash {
	if self.nBetPos == nBetPos {
		return self.Win()
	} else {
		return self.Lose()
	}
}

// Win 胜利
func (self *TRobot1) Win() Cash {
	nWin := Cash(0)
	if self.bBet {
		nWin = self.nBet * 97 / 100
		self.nMoney += (nWin + self.nBet)

		if self.nMoney > self.nMax {
			self.nMax = self.nMoney
		}
		self.nBet = 100 // 下注额恢复到1元
	}

	self.bBet = false

	return nWin
}

// Lose 失败
func (self *TRobot1) Lose() Cash {
	nLose := Cash(0)
	if self.bBet {
		nLose = self.nBet // 输掉下注额度
		self.nBet *= 2    // 进行倍投操作
	}
	self.bBet = false
	return -nLose
}

// Value 盈利值 = 目前的钱减去开始的带入
func (self *TRobot1) Value() Cash {
	return self.nMoney - self.nTakein
}

// NewRobot1 新建机器人1
func NewRobot1() *TRobot1 {
	return &TRobot1{
		nMoney:  1000000,
		nTakein: 1000000,
		nBet:    100,
	}
}
