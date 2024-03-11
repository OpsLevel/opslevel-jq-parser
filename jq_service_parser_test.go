package opslevel_jq_parser_test

import (
	_ "embed"
	"encoding/json"
	"fmt"
	"slices"
	"testing"

	opslevel_jq_parser "github.com/opslevel/opslevel-jq-parser/v2024"
	"github.com/rocktavious/autopilot/v2023"
)

//go:embed testdata/deployment.json
var k8sResource string

//go:embed testdata/simple_config.yaml
var simpleConfig string

//go:embed testdata/sample_config.yaml
var sampleConfig string

//go:embed testdata/simple_expectation.json
var simpleExpectation string

//go:embed testdata/sample_expectation.json
var sampleExpectation string

func TestJQServiceParser(t *testing.T) {
	type TestCase struct {
		name               string
		config             string
		expectedServiceReg string
	}
	testCases := []TestCase{
		{"using simple config", simpleConfig, simpleExpectation},
		{"using sample config", sampleConfig, sampleExpectation},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			config, err := opslevel_jq_parser.NewServiceRegistrationConfig(tc.config)
			autopilot.Ok(t, err)
			parser := opslevel_jq_parser.NewJQServiceParser(*config)
			service, err := parser.Run(k8sResource)
			autopilot.Ok(t, err)

			var expectedService opslevel_jq_parser.ServiceRegistration
			err = json.Unmarshal([]byte(tc.expectedServiceReg), &expectedService)
			autopilot.Ok(t, err)
			autopilot.Equals(t, expectedService, *service)
		})
	}
}

func BenchmarkJQParser_New(b *testing.B) {
	config, _ := opslevel_jq_parser.NewServiceRegistrationConfig(sampleConfig)
	parser := opslevel_jq_parser.NewJQServiceParser(*config)
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_, _ = parser.Run(k8sResource)
	}
}

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
	emptyList := make([]Beverage, 0)
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
