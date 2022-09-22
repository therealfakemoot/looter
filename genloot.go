package genloot

import (
	"fmt"
	"io"

	"github.com/BurntSushi/toml"
	wr "github.com/mroth/weightedrand"
)

type WeightedEntry struct {
	Mult   float64
	Weight int
}

type LootGenerator struct {
	Materials       map[string]WeightedEntry `toml:"material"`
	Quality         map[string]WeightedEntry `toml:"quality"`
	Items           map[string]float64       `toml:"items"`
	QualityChooser  *wr.Chooser
	MaterialChooser *wr.Chooser
	TypeChooser     *wr.Chooser
}

func NewLootGenerator(r io.Reader) (LootGenerator, error) {
	lg := LootGenerator{}
	err := lg.LoadReader(r)
	if err != nil {
		return lg, fmt.Errorf("error loading loot table from TOML: %w", err)
	}

	weightedMaterials := make([]wr.Choice, 0)
	for k, v := range lg.Materials {
		weightedMaterials = append(weightedMaterials, wr.Choice{
			Item:   k,
			Weight: uint(v.Weight),
		})
	}
	materialChooser, err := wr.NewChooser(weightedMaterials...)
	if err != nil {
		return lg, fmt.Errorf("error creating weighted Chooser for materials: %w", err)
	}

	lg.MaterialChooser = materialChooser

	weightedQualities := make([]wr.Choice, 0)
	for k, v := range lg.Quality {
		weightedQualities = append(weightedQualities, wr.Choice{
			Item:   k,
			Weight: uint(v.Weight),
		})
	}
	qualityChooser, err := wr.NewChooser(weightedQualities...)
	if err != nil {
		return lg, fmt.Errorf("error creating weighted Chooser for quality: %w", err)
	}
	lg.QualityChooser = qualityChooser

	weightedTypes := make([]wr.Choice, 0)
	for k := range lg.Items {
		weightedQualities = append(weightedQualities, wr.Choice{
			Item:   k,
			Weight: 1,
		})
	}
	typeChooser, err := wr.NewChooser(weightedTypes...)
	if err != nil {
		return lg, fmt.Errorf("error creating weighted Chooser for quality: %w", err)
	}
	lg.TypeChooser = typeChooser

	return lg, nil
}

func (lg *LootGenerator) LoadReader(r io.Reader) error {
	e := toml.NewDecoder(r)
	_, err := e.Decode(lg)
	return err
}

func (lg LootGenerator) Fill(target CashValue) []Item {
	items := make([]Item, 0)
	lootTotal := 0
	for lootTotal < target.UnitValue() {
		i := Item{}
		i.Table = lg
		i.Type = lg.TypeChooser.Pick().(string)
		i.BaseValue = lg.Items[i.Type]
		i.Material = lg.MaterialChooser.Pick().(string)
		i.Quality = lg.QualityChooser.Pick().(string)
	}

	return items
}
