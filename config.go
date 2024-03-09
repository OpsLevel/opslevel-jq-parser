package opslevel_jq_parser

import (
	"github.com/opslevel/opslevel-go/v2024"
	"gopkg.in/yaml.v3"
)

type TagRegistrationConfig struct {
	Assign []string `json:"assign" yaml:"assign"` // JQ expressions that return a single string or a map[string]string
	Create []string `json:"create" yaml:"create"` // JQ expressions that return a single string or a map[string]string
}

// ServiceRegistrationConfig represents the jq expressions configuration that can turn json data into a ServiceRegistration
type ServiceRegistrationConfig struct {
	Aliases      []string              `json:"aliases" yaml:"aliases"` // JQ expressions that return a single string or a []string
	Description  string                `json:"description" yaml:"description"`
	Framework    string                `json:"framework" yaml:"framework"`
	Language     string                `json:"language" yaml:"language"`
	Lifecycle    string                `json:"lifecycle" yaml:"lifecycle"`
	Name         string                `json:"name" yaml:"name"`
	Owner        string                `json:"owner" yaml:"owner"`
	Product      string                `json:"product" yaml:"product"`
	Properties   map[string]string     `json:"properties" yaml:"properties"`
	Repositories []string              `json:"repositories" yaml:"repositories"` // JQ expressions that return a single string or []string or map[string]string or a []map[string]string
	System       string                `json:"system" yaml:"system"`
	Tags         TagRegistrationConfig `json:"tags" yaml:"tags"`
	Tier         string                `json:"tier" yaml:"tier"`
	Tools        []string              `json:"tools" yaml:"tools"` // JQ expressions that return a single map[string]string or a []map[string]string
}

// ServiceRegistration represents the parsed json data from a ServiceRegistrationConfig
type ServiceRegistration struct {
	Aliases      []string                                `json:",omitempty"`
	Description  string                                  `json:",omitempty"`
	Framework    string                                  `json:",omitempty"`
	Language     string                                  `json:",omitempty"`
	Lifecycle    string                                  `json:",omitempty"`
	Name         string                                  `json:",omitempty"`
	Owner        string                                  `json:",omitempty"`
	Product      string                                  `json:",omitempty"`
	Properties   map[string]string                       `json:",omitempty"`
	Repositories []opslevel.ServiceRepositoryCreateInput `json:",omitempty"` // This is a concrete class so fields are validated during `service preview`
	System       string                                  `json:",omitempty"`
	TagAssigns   []opslevel.TagInput                     `json:",omitempty"`
	TagCreates   []opslevel.TagInput                     `json:",omitempty"`
	Tier         string                                  `json:",omitempty"`
	Tools        []opslevel.ToolCreateInput              `json:",omitempty"` // This is a concrete class so fields are validated during `service preview`
}

func NewServiceRegistrationConfig(data string) (*ServiceRegistrationConfig, error) {
	var output ServiceRegistrationConfig
	err := yaml.Unmarshal([]byte(data), &output)
	return &output, err
}
