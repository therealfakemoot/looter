package main

import (
	"flag"
	"log"
	"math/rand"
	"os"
	"time"

	"github.com/therealfakemoot/genloot"
)

func main() {

	var (
		pp, gp, ep, sp, cp int
		table              string
		seed               int64
	)

	flag.Int64Var(&seed, "seed", 0, "pRNG seed")
	flag.IntVar(&pp, "platinum", 0, "platinum pieces")
	flag.IntVar(&gp, "gold", 0, "gold pieces")
	flag.IntVar(&ep, "electrum", 0, "electrum pieces")
	flag.IntVar(&sp, "silver", 0, "silver pieces")
	flag.IntVar(&cp, "copper", 0, "copper pieces")

	flag.StringVar(&table, "table", "table.toml", "loot generation table")
	flag.Parse()

	//rand.Seed(seed)
	//rand.Seed(8675309)
	rand.Seed(time.Now().UnixNano())

	targetValue := genloot.CashValue{
		PP: pp,
		GP: gp,
		EP: ep,
		SP: sp,
		CP: cp,
	}

	lootTableFile, err := os.Open(table)
	if err != nil {
		log.Fatalf("error opening loot table file: %s\n", err)
	}

	g, err := genloot.NewLootGenerator(lootTableFile)
	if err != nil {
		log.Fatalf("error parsing TOML: %s\n", err)
	}

	log.Printf("Target loot value: %s\n", genloot.ReduceCoins(targetValue))
	items := g.Fill(targetValue)

	totalValue := genloot.CashValue{}
	for _, item := range items {
		totalValue.CP += item.CashValue().UnitValue()
	}
	log.Printf("Total loot value: %s\n", genloot.ReduceCoins(totalValue))

}
