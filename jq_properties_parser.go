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

func parsePropertiesInner(output map[string]opslevel.JsonString, prop map[string]interface{}) {
	// prop is map[my_object:{"message": "hello world", "boolean": true}]
	if len(prop) != 1 {
		log.Warn().Msg("properties parser: bad format")
		return
	}
	var def string
	for key := range prop {
		def = key
	}
	value, err := opslevel.NewJSONInput(prop[def])
	if err != nil {
		log.Warn().Err(err).Msgf("properties parser: expected valid json string")
		return
	}
	output[def] = *value
}

func parsePropertiesArray(output map[string]opslevel.JsonString, response string) {
	var properties []map[string]interface{}
	if err := json.Unmarshal([]byte(response), &properties); err != nil {
		log.Warn().Err(err).Msg("properties parser: error decoding inside array")
		return
	}
	for _, prop := range properties {
		parsePropertiesInner(output, prop)
	}
}

func parsePropertiesObject(output map[string]opslevel.JsonString, response string) {
	var prop map[string]interface{}
	if err := json.Unmarshal([]byte(response), &prop); err != nil {
		log.Warn().Err(err).Msg("properties parser: error decoding object")
		return
	}
	parsePropertiesInner(output, prop)
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

		if strings.HasPrefix(response, "[") && strings.HasSuffix(response, "]") {
			parsePropertiesArray(output, response)
		} else if strings.HasPrefix(response, "{") && strings.HasSuffix(response, "}") {
			parsePropertiesObject(output, response)
		} else {
			log.Warn().Msg("properties parser: expected array or object")
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
