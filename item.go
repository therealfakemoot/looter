package looter

import (
	"fmt"
)

type Item struct {
	Type      string
	BaseValue float64
	Quality   string
	Material  string
	Table     LootGenerator
}

func (i Item) CashValue() CashValue {
	cv := CashValue{}
	// base values are in GP, but it'll be easier to work with copper pieces as the unit
	baseValueInCoppers := i.Table.Items[i.Type] * 100
	baseValueInCoppers *= i.Table.Materials[i.Material].Mult
	baseValueInCoppers *= i.Table.Quality[i.Quality].Mult

	cv.CP = int(baseValueInCoppers)
	return ReduceCoins(cv)
}

func (i Item) String() string {
	return fmt.Sprintf("%s %s %s (%s)", i.Quality, i.Material, i.Type, i.CashValue())
}
