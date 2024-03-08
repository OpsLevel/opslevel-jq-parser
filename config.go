package opslevel_jq_parser

import (
	_ "embed"
	"gopkg.in/yaml.v3"
)

// ServiceRegistrationConfig represents the jq expressions configuration that can turn json data into a ServiceRegistration
type ServiceRegistrationConfig struct {
	Aliases     []string          `json:"aliases,omitempty" yaml:"aliases,omitempty"`
	Description string            `json:"description,omitempty" yaml:"description,omitempty"`
	Framework   string            `json:"framework,omitempty" yaml:"framework,omitempty"`
	Language    string            `json:"language,omitempty" yaml:"language,omitempty"`
	Lifecycle   string            `json:"lifecycle,omitempty" yaml:"lifecycle,omitempty"`
	Name        string            `json:"name,omitempty" yaml:"name,omitempty"`
	Owner       string            `json:"owner,omitempty" yaml:"owner,omitempty"`
	Product     string            `json:"product,omitempty" yaml:"product,omitempty"`
	Properties  map[string]string `json:"properties,omitempty" yaml:"properties,omitempty"`
	System      string            `json:"system,omitempty" yaml:"system,omitempty"`
	Tier        string            `json:"tier,omitempty" yaml:"tier,omitempty"`
}

// ServiceRegistration represents the parsed json data from a ServiceRegistrationConfig
type ServiceRegistration struct {
	Aliases     []string          `json:"aliases,omitempty"`
	Description string            `json:"description,omitempty"`
	Framework   string            `json:"framework,omitempty"`
	Language    string            `json:"language,omitempty"`
	Lifecycle   string            `json:"lifecycle,omitempty"`
	Name        string            `json:"name,omitempty"`
	Owner       string            `json:"owner,omitempty"`
	Product     string            `json:"product,omitempty"`
	Properties  map[string]string `json:"properties,omitempty"`
	System      string            `json:"system,omitempty"`
	Tier        string            `json:"tier,omitempty"`
}

func NewServiceRegistrationConfig(data string) (ServiceRegistrationConfig, error) {
	var output ServiceRegistrationConfig
	err := yaml.Unmarshal([]byte(data), &output)
	if err != nil {
		return ServiceRegistrationConfig{}, err
	}
	return output, nil
}
