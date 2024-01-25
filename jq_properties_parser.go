package opslevel_jq_parser

import (
	"encoding/json"
	"strings"

	"github.com/rs/zerolog/log"

	"github.com/opslevel/opslevel-go/v2024"
)

type JQPropertiesParser struct {
	assigns []*JQFieldParser
}

func NewJQPropertiesParser(cfg PropertyRegistrationConfig) *JQPropertiesParser {
	assigns := make([]*JQFieldParser, len(cfg.Assign))
	for i, expression := range cfg.Assign {
		assigns[i] = NewJQFieldParser(expression)
	}
	return &JQPropertiesParser{
		assigns: assigns,
	}
}

func propertyAssignmentFromMap(definition string, input map[string]interface{}) (opslevel.PropertyInput, error) {
	var owner string
	var value opslevel.JsonString
	log.Debug().Msgf("property assignment '%s' input is %#v", definition, input)
	for k, v := range input {
		switch k {
		case "owner":
			owner = v.(string)
		case "value":
			jsonString, err := opslevel.NewJSONInput(v)
			if err != nil {
				return opslevel.PropertyInput{}, err
			}
			value = *jsonString
		}
	}
	if owner != "" && value != "" {
		return opslevel.PropertyInput{
			Owner:      *opslevel.NewIdentifier(owner),
			Definition: *opslevel.NewIdentifier(definition),
			Value:      value,
		}, nil
	}
	log.Warn().Msgf("got incomplete property assignment '%s'", definition)
	return opslevel.PropertyInput{}, nil
}

func (p *JQPropertiesParser) parse(programs []*JQFieldParser, data string) ([]opslevel.PropertyInput, error) {
	output := make([]opslevel.PropertyInput, 0, len(programs))
	for _, program := range programs {
		response, err := program.Run(data)
		if err != nil {
			log.Warn().Msg("properties parser: jq error")
			continue
		}
		if response == "" {
			log.Warn().Msg("properties parser: jq returned empty string")
			continue
		}

		if strings.HasPrefix(response, "[") && strings.HasSuffix(response, "]") {
			var properties []map[string]string
			if err := json.Unmarshal([]byte(response), &properties); err != nil {
				log.Warn().Msg("skipping a property assignment - error decoding at start")
				continue
			}
			for _, item := range properties {
				if len(item) != 1 {
					log.Warn().Msg("skipping a property assignment - bad format")
					continue
				}
				var def string
				for k := range item {
					def = k
				}
				var propertyBody map[string]interface{}
				if err := json.Unmarshal([]byte(item[def]), &propertyBody); err != nil {
					log.Warn().Msgf("got error decoding property assignment '%s'", def)
					continue
				}
				out, err := propertyAssignmentFromMap(def, propertyBody)
				if err != nil {
					log.Warn().Msgf("got error parsing property assignment '%s'", def)
					continue
				}
				output = append(output, out)
			}
		}
	}
	return output, nil
}

func (p *JQPropertiesParser) Run(data string) ([]opslevel.PropertyInput, error) {
	result, err := p.parse(p.assigns, data)
	if err != nil {
		return nil, err
	}
	return result, nil
}
