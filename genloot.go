package main

import (
	"flag"
	"fmt"
	"github.com/BurntSushi/toml"
	"io"
	"log"
	"math/rand"
	"os"
	"time"

	wr "github.com/mroth/weightedrand"
)

type WeightedEntry struct {
	Mult   float64
	Weight int
}

type LootGenerator struct {
	Materials map[string]WeightedEntry `toml:"material"`
	Quality   map[string]WeightedEntry `toml:"quality"`
	Items     map[string]float64       `toml:"items"`
}

func (gc *LootGenerator) LoadReader(r io.Reader) error {
	e := toml.NewDecoder(r)
	_, err := e.Decode(gc)
	return err
}

type CashValue struct {
	CP int
	SP int
	EP int
	GP int
	PP int
}

func (cv CashValue) UnitValue() int {

	return 0
}

func (l CashValue) String() string {
	return fmt.Sprintf("%dPP,%dGP,%dEP,%dSP,%dCP", l.PP, l.GP, l.EP, l.SP, l.CP)
}

type Item struct {
	Type      string
	BaseValue float64
	Quality   string
	Material  string
}

func (i Item) CashValue(gen LootGenerator) CashValue {
	cv := CashValue{}
	// base values are in GP, but it'll be easier to work with copper pieces as the unit
	baseValueInCoppers := gen.Items[i.Type] * 100
	baseValueInCoppers *= gen.Materials[i.Material].Mult
	baseValueInCoppers *= gen.Quality[i.Quality].Mult

	return cv
}

func (i Item) String() string {
	return fmt.Sprintf("%s %s %s", i.Quality, i.Material, i.Type)
}

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
		l := CashValue{
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

	g := LootGenerator{}
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

	i := Item{}
	// pick an item type, then material, then quality
	itemTypes := make([]string, 0)
	for k := range g.Items {
		itemTypes = append(itemTypes, k)
	}

	i.Type = itemTypes[rand.Intn(len(itemTypes))]
	i.BaseValue = g.Items[itemTypes[rand.Intn(len(itemTypes))]]
	i.Material = materialChooser.Pick().(string)
	i.Quality = qualityChooser.Pick().(string)
	log.Println(i)

	i.Type = itemTypes[rand.Intn(len(itemTypes))]
	i.BaseValue = g.Items[itemTypes[rand.Intn(len(itemTypes))]]
	i.Material = materialChooser.Pick().(string)
	i.Quality = qualityChooser.Pick().(string)
	log.Println(i)

	i.Type = itemTypes[rand.Intn(len(itemTypes))]
	i.BaseValue = g.Items[itemTypes[rand.Intn(len(itemTypes))]]
	i.Material = materialChooser.Pick().(string)
	i.Quality = qualityChooser.Pick().(string)
	log.Println(i)

	i.Type = itemTypes[rand.Intn(len(itemTypes))]
	i.BaseValue = g.Items[itemTypes[rand.Intn(len(itemTypes))]]
	i.Material = materialChooser.Pick().(string)
	i.Quality = qualityChooser.Pick().(string)
	log.Println(i)
}
