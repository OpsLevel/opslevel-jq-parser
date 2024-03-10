package opslevel_jq_parser

import (
	"encoding/json"
	"github.com/opslevel/opslevel-go/v2024"
)

type JQDictParser map[string]JQFieldParser

func NewJQDictParser(dict map[string]string) JQDictParser {
	output := make(map[string]JQFieldParser)
	for key, expression := range dict {
		output[key] = NewJQFieldParser(expression)
	}
	return output
}

func (p JQDictParser) Run(data string) map[string]opslevel.PropertyInput {
	output := make(map[string]opslevel.PropertyInput)
	for key, expression := range p {
		jqRes := expression.Run(data)
		if jqRes == "" {
			continue
		}
		var propertyInput opslevel.PropertyInput
		err := json.Unmarshal([]byte(jqRes), &propertyInput)
		if err != nil {
			continue
		}
		output[key] = propertyInput
	}
	return output
}
