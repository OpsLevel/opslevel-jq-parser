package opslevel_jq_parser

import (
	"encoding/json"
	"strings"

	"github.com/rs/zerolog/log"

	"github.com/opslevel/opslevel-go/v2024"
)

type JQPropertiesParser struct {
	programs []*JQFieldParser
}

func NewJQPropertiesParser(expressions []string) *JQPropertiesParser {
	programs := make([]*JQFieldParser, len(expressions))
	for i, expression := range expressions {
		programs[i] = NewJQFieldParser(expression)
	}
	return &JQPropertiesParser{
		programs: programs,
	}
}

// parse returns a map so the definitions are already deduplicated.
func (p *JQPropertiesParser) parse(data string) (map[string]opslevel.JsonString, error) {
	output := make(map[string]opslevel.JsonString)
	for _, program := range p.programs {
		response, err := program.Run(data)
		if err != nil {
			log.Warn().Err(err).Msg("properties parser: jq error")
			continue
		}
		if response == "" {
			log.Warn().Msg("properties parser: jq returned empty string")
			continue
		}

		if !strings.HasPrefix(response, "[") && !strings.HasSuffix(response, "]") {
			log.Warn().Msg("properties parser: expected array")
			continue
		}
		var properties []map[string]interface{}
		if err := json.Unmarshal([]byte(response), &properties); err != nil {
			log.Warn().Err(err).Msg("properties parser: error decoding inside array")
			continue
		}
		for _, prop := range properties {
			// prop is map[my_object:{"message": "hello world", "boolean": true}]
			if len(prop) != 1 {
				log.Warn().Msg("properties parser: bad format")
				continue
			}
			var def string
			for key := range prop {
				def = key
			}
			value, err := opslevel.NewJSONInput(prop[def])
			if err != nil {
				log.Warn().Err(err).Msgf("properties parser: expected valid json string")
				continue
			}
			output[def] = *value
		}
	}
	return output, nil
}

func (p *JQPropertiesParser) Run(data string) (map[string]opslevel.JsonString, error) {
	result, err := p.parse(data)
	if err != nil {
		return nil, err
	}
	return result, nil
}
