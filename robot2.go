package main

import "fmt"

type TRobot2 struct {
	nMoney  int64
	nTakein int64
	nBet    int64
	bBet    bool // 是否下注
}

// Bet 下注
func (self *TRobot2) Bet() {
	if self.bBet {
		return
	}

	if self.nMoney <= 100 {
		// 破产了
		return
	}

	// 钱不够翻倍的时候
	if self.nMoney < self.nBet {
		self.nBet = self.nMoney
	}

	self.bBet = true
	self.nMoney -= self.nBet // 扣下注额度
}

// Win 胜利
func (self *TRobot2) Win() int64 {
	nWin := int64(0)
	if self.bBet {
		nWin = self.nBet * 97 / 100
		self.nMoney += (nWin + self.nBet)
		self.nBet = 100 // 下注额恢复到1元
	}

	self.bBet = false

	return nWin
}

// Lose 失败
func (self *TRobot2) Lose() int64 {
	nLose := int64(0)
	if self.bBet {
		nLose = self.nBet // 输掉下注额度
		self.nBet *= 2    // 进行倍投操作
	}
	self.bBet = false
	return nLose
}

func (self *TRobot2) String() string {
	return fmt.Sprintf("%.2f", float64(self.nMoney)/100)
}

// Value 盈利值 = 目前的钱减去开始的带入
func (self *TRobot2) Value() int64 {
	return self.nMoney - self.nTakein
}

// NewRobot2 新建机器人1
func NewRobot2() *TRobot2 {
	return &TRobot2{
		nMoney:  1000000,
		nTakein: 1000000,
		nBet:    100,
	}
}
