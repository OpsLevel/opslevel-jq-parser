package opslevel_jq_parser_test

import (
	"cmp"
	_ "embed"
	"encoding/json"
	"slices"
	"testing"

	opslevel "github.com/opslevel/opslevel-go/v2024"
	opslevel_jq_parser "github.com/opslevel/opslevel-jq-parser/v2024"
	"github.com/rocktavious/autopilot/v2023"
)

//go:embed testdata/deployment.json
var k8sResource string

//go:embed testdata/simple_config.yaml
var simpleConfig string

//go:embed testdata/sample_config.yaml
var sampleConfig string

//go:embed testdata/duplicate_config.yaml
var duplicateConfig string

//go:embed testdata/simple_expectation.json
var simpleExpectation string

//go:embed testdata/sample_expectation.json
var sampleExpectation string

//go:embed testdata/duplicate_expectation.json
var duplicateExpectation string

func compAny[T any](a, b T) int {
	x, _ := json.Marshal(a)
	y, _ := json.Marshal(b)
	return cmp.Compare(string(x), string(y))
}

func sortSlices(service *opslevel_jq_parser.ServiceRegistration) {
	slices.Sort(service.Aliases)
	slices.SortFunc(service.Properties, compAny[opslevel.PropertyInput])
	slices.SortFunc(service.Repositories, compAny[opslevel.ServiceRepositoryCreateInput])
	slices.SortFunc(service.TagAssigns, compAny[opslevel.TagInput])
	slices.SortFunc(service.TagCreates, compAny[opslevel.TagInput])
	slices.SortFunc(service.Tools, compAny[opslevel.ToolCreateInput])
}

func TestJQServiceParser(t *testing.T) {
	type TestCase struct {
		name               string
		config             string
		expectedServiceReg string
	}
	testCases := []TestCase{
		{"using simple config", simpleConfig, simpleExpectation},
		{"using sample config", sampleConfig, sampleExpectation},
		{"using duplicate config", duplicateConfig, duplicateExpectation},
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

			// order of the slices does not matter - JSON marshal will output struct keys in order defined in the struct
			// so before comparing, sort the slices
			sortSlices(service)
			sortSlices(&expectedService)
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
