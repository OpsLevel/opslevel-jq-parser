package opslevel_jq_parser_test

import (
	_ "embed"
	"testing"

	opslevel_jq_parser "github.com/opslevel/opslevel-jq-parser/v2024"
	"github.com/rocktavious/autopilot/v2023"
)

//go:embed fixtures/deployment.json
var k8sResource string

//go:embed fixtures/config_simple.yaml
var SimpleConfig string

//go:embed fixtures/config_sample.yaml
var SampleConfig string

func TestJQServiceParserSimpleConfig(t *testing.T) {
	config, err := opslevel_jq_parser.NewServiceRegistrationConfig(SimpleConfig)
	autopilot.Ok(t, err)
	service := opslevel_jq_parser.RunWithConfig(config, k8sResource)

	// basic values
	autopilot.Equals(t, "", service.Description)
	autopilot.Equals(t, "", service.Framework)
	autopilot.Equals(t, "", service.Language)
	autopilot.Equals(t, "", service.Lifecycle)
	autopilot.Equals(t, "web", service.Name)
	autopilot.Equals(t, "self-hosted", service.Owner)
	autopilot.Equals(t, "", service.Product)
	autopilot.Equals(t, "", service.System)
	autopilot.Equals(t, "", service.Tier)

	// array of basic values
	autopilot.Equals(t, 1, len(service.Aliases))
	autopilot.Equals(t, "k8s:web-self-hosted", service.Aliases[0])
}

func TestJQServiceParserSampleConfig(t *testing.T) {
	config, err := opslevel_jq_parser.NewServiceRegistrationConfig(SampleConfig)
	autopilot.Ok(t, err)
	service := opslevel_jq_parser.RunWithConfig(config, k8sResource)

	// basic values (string, boolean, number)
	autopilot.Equals(t, "this is a description", service.Description)
	autopilot.Equals(t, "rails", service.Framework)
	autopilot.Equals(t, "ruby", service.Language)
	autopilot.Equals(t, "alpha", service.Lifecycle)
	autopilot.Equals(t, "web", service.Name)
	autopilot.Equals(t, "velero", service.Owner)
	autopilot.Equals(t, "jklabs", service.Product)
	autopilot.Equals(t, "monolith", service.System)
	autopilot.Equals(t, "tier_1", service.Tier)

	// array of basic values
	autopilot.Equals(t, 2, len(service.Aliases))
	autopilot.Equals(t, "k8s:web-self-hosted", service.Aliases[0])
	autopilot.Equals(t, "self-hosted-web", service.Aliases[1])

	// dictionary
	autopilot.Equals(t, 5, len(service.Properties))
	autopilot.Equals(t, "true", service.Properties["prop_bool"])
	autopilot.Equals(t, "{}", service.Properties["prop_empty_object"])
	autopilot.Equals(t, "", service.Properties["prop_empty_string"])
	autopilot.Equals(t, `{"message":"hello world","condition":true}`, service.Properties["prop_object"])
	autopilot.Equals(t, "hello world", service.Properties["prop_string"])

	// array of dictionaries
}

func BenchmarkJQParser_New(b *testing.B) {
	config, _ := opslevel_jq_parser.NewServiceRegistrationConfig(SimpleConfig)
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = opslevel_jq_parser.RunWithConfig(config, k8sResource)
	}
}
