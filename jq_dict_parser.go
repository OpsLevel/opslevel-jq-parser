package opslevel_jq_parser

import (
	"fmt"

	libjq_go "github.com/flant/libjq-go"
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

func (p JQDictParser) Run(data string) (map[string]string, error) {
	output := make(map[string]string)
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
		output[key] = jqRes
	}
	return output, nil
}
