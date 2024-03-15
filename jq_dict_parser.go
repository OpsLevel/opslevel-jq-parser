package opslevel_jq_parser

import (
	"fmt"

	"github.com/opslevel/opslevel-go/v2024"

	libjq_go "github.com/flant/libjq-go"
	"github.com/rs/zerolog/log"
)

type JQDictParser map[string]JQFieldParser

func NewJQDictParser(dict map[string]string) map[string]JQFieldParser {
	output := make(map[string]JQFieldParser)
	if dict == nil {
		return output
	}
	jq := libjq_go.Jq()
	for key, expression := range dict {
		prg, err := jq.Program(expression).Precompile()
		if err != nil {
			panic(fmt.Sprintf("unable to compile jq dict: %s", dict))
		}
		output[key] = JQFieldParser{
			program: prg,
		}
	}
	return output
}

func (p JQDictParser) Run(data string) ([]opslevel.PropertyInput, error) {
	output := make([]opslevel.PropertyInput, 0)
	for key, expression := range p {
		jqRes, err := expression.Run(data)
		if err != nil {
			log.Warn().Str("key", key).Err(err).Msg("error running jq expression")
			continue
		}
		if jqRes == "" {
			continue
		}
		propertyInput := opslevel.PropertyInput{
			Definition: *opslevel.NewIdentifier(key),
			Value:      opslevel.JsonString(jqRes),
		}
		output = append(output, propertyInput)
	}
	return output, nil
}
