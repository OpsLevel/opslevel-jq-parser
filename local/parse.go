package main

import (
	_ "embed"
	"fmt"

	opslevel_jq_parser "github.com/opslevel/opslevel-jq-parser/v2024"
)

//go:embed config.yaml
var configYamlFile string

//go:embed deployment.json
var deploymentJson string

func main() {
	config := getServiceRegistrationConfig()
	parser := opslevel_jq_parser.NewJQServiceParser(*config)
	service, err := parser.Run(deploymentJson)
	if err != nil {
		panic(err)
	}
	fmt.Println(service)
}

func getServiceRegistrationConfig() *opslevel_jq_parser.ServiceRegistrationConfig {
	svcRegConfig, err := opslevel_jq_parser.NewServiceRegistrationConfig(configYamlFile)
	if err != nil {
		panic(err)
	}
	return svcRegConfig
}
