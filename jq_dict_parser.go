package opslevel_jq_parser

import (
	"fmt"
	"strings"

	libjq_go "github.com/flant/libjq-go"
	"github.com/opslevel/opslevel-go/v2024"
	"github.com/rs/zerolog/log"
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
		jqRes, err := expression.Run(data)
		if err != nil {
			log.Warn().Str("key", key).Err(err).Msg("error running jq expression")
			continue
		}
		if jqRes == "null" {
			// in the case that the expression returned nothing (happens in the case where the key was not found)
			// jq will return "null". This is not the same as empty string. So in that case, skip the item.
			continue
		}
		if strings.HasPrefix(jqRes, "{") && strings.HasSuffix(jqRes, "}") {
			// TODO: this can be placed inside the NewJSONInput function in opslevel-go
			// if the input given there is a string
			schema, err := opslevel.NewJSONSchema(jqRes)
			if err != nil {
				log.Warn().Str("key", key).Err(err).Msg("error decoding object")
				continue
			}
			schemaString := schema.AsString()
			output[key] = opslevel.JsonString(schemaString)
			continue
		}
		parsed, err := opslevel.NewJSONInput(jqRes)
		if err != nil {
			log.Warn().Str("key", key).Err(err).Msg("error decoding json")
			continue
		}
		output[key] = *parsed
	}
	return output, nil
}
