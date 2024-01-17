package opslevel_jq_parser_test

import (
	"fmt"
	"slices"
	"testing"

	opslevel_jq_parser "github.com/opslevel/opslevel-jq-parser/v2024"
)

type Beverage struct {
	Name string
	Oz   int
}

func DeduplicatedBeverages(objects []Beverage) []Beverage {
	return opslevel_jq_parser.Deduplicated(objects, func(b Beverage) string {
		return fmt.Sprintf("%s%d", b.Name, b.Oz)
	})
}

func BeveragesEqual(b1 []Beverage, b2 []Beverage) bool {
	return slices.EqualFunc(b1, b2, func(b1, b2 Beverage) bool {
		return b1.Name == b2.Name && b1.Oz == b2.Oz
	})
}

func TestDeduplicated(t *testing.T) {
	emptyList := []Beverage{}
	emptyDedup := DeduplicatedBeverages(emptyList)
	if !BeveragesEqual(emptyList, emptyDedup) {
		t.Error("an empty list deduplicated should be equal to itself")
	}

	oneElem := []Beverage{
		{Name: "Energy Drink", Oz: 10},
	}
	oneElemDedup := DeduplicatedBeverages(oneElem)
	if !BeveragesEqual(oneElem, oneElemDedup) {
		t.Error("a single element list deduplicated should be equal to itself")
	}

	list := []Beverage{
		{Name: "Soda", Oz: 12},
		{Name: "Iced Tea", Oz: 12},
		{Name: "Soda", Oz: 12},
		{Name: "Soda", Oz: 12},
		{Name: "Iced Tea", Oz: 12},
		{Name: "Iced Tea", Oz: 24},
		{Name: "Soda", Oz: 24},
		{Name: "Energy Drink", Oz: 10},
	}
	listDedup := DeduplicatedBeverages(list)
	listDedupExp := []Beverage{
		{Name: "Soda", Oz: 12},
		{Name: "Iced Tea", Oz: 12},
		{Name: "Iced Tea", Oz: 24},
		{Name: "Soda", Oz: 24},
		{Name: "Energy Drink", Oz: 10},
	}
	if BeveragesEqual(list, listDedup) {
		t.Error("long list deduplicated should NOT be equal to itself")
	}
	if !BeveragesEqual(listDedup, listDedupExp) {
		t.Error("long list deduplicated should be equal to the expected list")
	}
}
