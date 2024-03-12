package opslevel_jq_parser

import (
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

func (p JQDictParser) Run(data string) map[string]opslevel.JsonString {
	output := make(map[string]opslevel.JsonString)
	for key, expression := range p {
		jqRes := expression.Run(data)
		if jqRes == "" {
			continue
		}
		output[key] = opslevel.JsonString(jqRes)
	}
	return output
}
