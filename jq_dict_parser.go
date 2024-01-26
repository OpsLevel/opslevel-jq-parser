package opslevel_jq_parser

import (
	"fmt"

	libjq_go "github.com/flant/libjq-go"
	"github.com/opslevel/opslevel-go/v2024"
)

type JQDictParser map[string]JQFieldParser

func NewJQDictParser(dict map[string]string) map[string]JQFieldParser {
	output := make(map[string]JQFieldParser)
	if dict == nil {
		return output
	}
	for key, expression := range dict {
		prg, err := libjq_go.Jq().Program(expression).Precompile()
		if err != nil {
			panic(fmt.Sprintf("unable to compile jq dict: %s", dict))
		}
		output[key] = JQFieldParser{
			program: prg,
		}
	}
	return output
}

func (p JQDictParser) Run(data string) (map[string]opslevel.JsonString, error) {
	output := make(map[string]opslevel.JsonString)
	for key, expression := range p {
		x, err := expression.Run(data)
		if err != nil {
			return nil, err
		}
		fmt.Printf("x is '%s' %T\n", x, x)
		if x == "null" {
			continue
		}
		parsed, err := opslevel.NewJSONInput(x)
		if err != nil {
			return nil, err
		}
		output[key] = *parsed
	}
	return output, nil
}
