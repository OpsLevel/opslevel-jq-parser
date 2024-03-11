package opslevel_jq_parser

import (
	"github.com/opslevel/opslevel-go/v2024"
	"gopkg.in/yaml.v3"
)

type TagRegistrationConfig struct {
	Assign []string `json:"assign" yaml:"assign"`
	Create []string `json:"create" yaml:"create"`
}

// ServiceRegistrationConfig represents the jq expressions configuration that can turn json data into a ServiceRegistration
type ServiceRegistrationConfig struct {
	Aliases      []string              `json:"aliases" yaml:"aliases"`
	Description  string                `json:"description" yaml:"description"`
	Framework    string                `json:"framework" yaml:"framework"`
	Language     string                `json:"language" yaml:"language"`
	Lifecycle    string                `json:"lifecycle" yaml:"lifecycle"`
	Name         string                `json:"name" yaml:"name"`
	Owner        string                `json:"owner" yaml:"owner"`
	Product      string                `json:"product" yaml:"product"`
	Properties   map[string]string     `json:"properties" yaml:"properties"`
	Repositories []string              `json:"repositories" yaml:"repositories"`
	System       string                `json:"system" yaml:"system"`
	Tags         TagRegistrationConfig `json:"tags" yaml:"tags"`
	Tier         string                `json:"tier" yaml:"tier"`
	Tools        []string              `json:"tools" yaml:"tools"`
}

// ServiceRegistration represents the parsed json data from a ServiceRegistrationConfig
type ServiceRegistration struct {
	Aliases      []string                                `json:"aliases,omitempty"`
	Description  string                                  `json:"description,omitempty"`
	Framework    string                                  `json:"framework,omitempty"`
	Language     string                                  `json:"language,omitempty"`
	Lifecycle    string                                  `json:"lifecycle,omitempty"`
	Name         string                                  `json:"name,omitempty"`
	Owner        string                                  `json:"owner,omitempty"`
	Product      string                                  `json:"product,omitempty"`
	Properties   map[string]string                       `json:"properties,omitempty"`
	Repositories []opslevel.ServiceRepositoryCreateInput `json:"repositories,omitempty"`
	System       string                                  `json:"system,omitempty"`
	TagAssigns   []opslevel.TagInput                     `json:"tagAssigns,omitempty"`
	TagCreates   []opslevel.TagInput                     `json:"tagCreates,omitempty"`
	Tier         string                                  `json:"tier,omitempty"`
	Tools        []opslevel.ToolCreateInput              `json:"tools,omitempty"`
}

func NewServiceRegistrationConfig(data string) (ServiceRegistrationConfig, error) {
	var output ServiceRegistrationConfig
	err := yaml.Unmarshal([]byte(data), &output)
	return output, err
}
