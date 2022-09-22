package main

import (
	"flag"
	"log"
	"math/rand"
	"os"
	"time"

	wr "github.com/mroth/weightedrand"
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

	/*
		targetValue := genloot.CashValue{
			PP: pp,
			GP: gp,
			EP: ep,
			SP: sp,
			CP: cp,
		}
	*/

	lootTableFile, err := os.Open(table)
	if err != nil {
		log.Fatalf("error opening loot table file: %s\n", err)
	}

	g := genloot.LootGenerator{}
	err = g.LoadReader(lootTableFile)
	if err != nil {
		log.Fatalf("error parsing TOML: %s\n", err)
	}

	weightedMaterials := make([]wr.Choice, 0)
	for k, v := range g.Materials {
		weightedMaterials = append(weightedMaterials, wr.Choice{
			Item:   k,
			Weight: uint(v.Weight),
		})
	}
	materialChooser, err := wr.NewChooser(weightedMaterials...)

	weightedQualities := make([]wr.Choice, 0)
	for k, v := range g.Quality {
		weightedQualities = append(weightedQualities, wr.Choice{
			Item:   k,
			Weight: uint(v.Weight),
		})
	}
	qualityChooser, err := wr.NewChooser(weightedQualities...)

	i := genloot.Item{}
	// pick an item type, then material, then quality
	itemTypes := make([]string, 0)
	for k := range g.Items {
		itemTypes = append(itemTypes, k)
	}

	i.Table = g
	i.Type = itemTypes[rand.Intn(len(itemTypes))]
	i.BaseValue = g.Items[itemTypes[rand.Intn(len(itemTypes))]]
	i.Material = materialChooser.Pick().(string)
	i.Quality = qualityChooser.Pick().(string)
	log.Println(i)
}
