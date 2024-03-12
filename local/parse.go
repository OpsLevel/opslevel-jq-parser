package main

import (
	_ "embed"
	"fmt"
	"os"

	opslevel_jq_parser "github.com/opslevel/opslevel-jq-parser/v2024"
)

//go:embed config.yaml
var configYamlFile string

var deploymentJson string

func main() {
	deploymentJson, err := os.ReadFile("../testdata/deployment.json")
	if err != nil {
		panic(err)
	}

	config := getServiceRegistrationConfig()
	parser := opslevel_jq_parser.NewJQServiceParser(*config)
	service, err := parser.Run(string(deploymentJson))
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
