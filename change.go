package genloot

import (
	"fmt"
)

type CashValue struct {
	CP int
	SP int
	EP int
	GP int
	PP int
}

func (cv CashValue) UnitValue() int {

	return cv.CP + (cv.SP * 10) + (cv.EP * 50) + (cv.GP * 100) + (cv.PP * 1000)
}

func (l CashValue) String() string {
	return fmt.Sprintf("%dPP,%dGP,%dEP,%dSP,%dCP", l.PP, l.GP, l.EP, l.SP, l.CP)
}
func ReduceCoins(cv CashValue) CashValue {
	reduced := CashValue{}
	total := 0
	total += cv.CP
	total += cv.SP * 10
	total += cv.EP * 50
	total += cv.GP * 100
	total += cv.PP * 1000

	reduced.PP = total / 1000
	total -= reduced.PP * 1000

	reduced.GP = total / 100
	total -= reduced.GP * 100

	reduced.EP = total / 50
	total -= reduced.EP * 50

	reduced.SP = total / 10
	total -= reduced.SP * 10

	reduced.CP = total
	return reduced
}
