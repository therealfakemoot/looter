package main

import (
	"log"
)

func reduceCoins(cv CashValue) CashValue {
	reduced := CashValue{}
	total := 0
	total += cv.CP
	total += cv.SP * 10
	total += cv.EP * 50
	total += cv.GP * 100
	total += cv.PP * 1000

	log.Printf("Raw Total: %d", total)
	reduced.PP = total / 1000
	total -= reduced.PP * 1000
	log.Printf("Post-Platinum Total: %d", total)

	reduced.GP = total / 100
	total -= reduced.GP * 100
	log.Printf("Post-Gold Total: %d", total)

	reduced.EP = total / 50
	total -= reduced.EP * 50
	log.Printf("Post-Electrum Total: %d", total)

	reduced.SP = total / 10
	total -= reduced.SP * 10
	log.Printf("Post-Silver Total: %d", total)

	reduced.CP = total
	return reduced
}
