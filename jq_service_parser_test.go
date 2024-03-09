package opslevel_jq_parser_test

import (
	_ "embed"
	"encoding/json"
	"testing"

	opslevel_jq_parser "github.com/opslevel/opslevel-jq-parser/v2024"
	"github.com/rocktavious/autopilot/v2023"
)

//go:embed fixtures/deployment.json
var k8sResource string

//go:embed fixtures/config_simple.yaml
var configSimple string

//go:embed fixtures/config_sample.yaml
var configSample string

//go:embed fixtures/config_sample_expectation.json
var configSampleExpectation string

func TestJQServiceParserSimpleConfig(t *testing.T) {
	config, err := opslevel_jq_parser.NewServiceRegistrationConfig(configSimple)
	autopilot.Ok(t, err)
	service := opslevel_jq_parser.RunWithConfig(config, k8sResource)

	// values
	autopilot.Equals(t, "", service.Description)
	autopilot.Equals(t, "", service.Framework)
	autopilot.Equals(t, "", service.Language)
	autopilot.Equals(t, "", service.Lifecycle)
	autopilot.Equals(t, "web", service.Name)
	autopilot.Equals(t, "self-hosted", service.Owner)
	autopilot.Equals(t, "", service.Product)
	autopilot.Equals(t, "", service.System)
	autopilot.Equals(t, "", service.Tier)

	// array of values
	autopilot.Equals(t, 1, len(service.Aliases))
	autopilot.Equals(t, "k8s:web-self-hosted", service.Aliases[0])
}

func TestJQServiceParserSampleConfig(t *testing.T) {
	config, err := opslevel_jq_parser.NewServiceRegistrationConfig(configSample)
	autopilot.Ok(t, err)
	service := opslevel_jq_parser.RunWithConfig(config, k8sResource)

	var expectedService opslevel_jq_parser.ServiceRegistration
	err = json.Unmarshal([]byte(configSampleExpectation), &expectedService)
	autopilot.Ok(t, err)
	autopilot.Equals(t, expectedService, service)
}

func BenchmarkJQParser_New(b *testing.B) {
	config, _ := opslevel_jq_parser.NewServiceRegistrationConfig(configSimple)
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = opslevel_jq_parser.RunWithConfig(config, k8sResource)
	}
}
