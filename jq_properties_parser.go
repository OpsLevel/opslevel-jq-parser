package opslevel_jq_parser

import (
	"encoding/json"
	"fmt"
	"github.com/rs/zerolog/log"
	"strings"

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
	owner := ""
	var value opslevel.JsonString
	log.Debug().Str("func", "propertyAssignmentFromMap").Str("def", definition).Msgf("input %#v", input)
	for k, v := range input {
		switch k {
		case "owner":
			owner = v.(string)
		case "value":
			// TODO: safer type cast
			jsonString, err := opslevel.NewJSONInput(v)
			if err != nil {
				// TODO: log here
				return opslevel.PropertyInput{}, err
			}
			value = *jsonString
		}
	}
	if owner != "" && value != "" {
		return opslevel.PropertyInput{
			Owner:      *opslevel.NewIdentifier(owner),
			Definition: *opslevel.NewIdentifier(definition),
			Value:      value, // TODO: is this correct
		}, nil
	}
	log.Warn().Str("func", "propertyAssignmentFromMap").Str("definition", definition).Str("owner", owner).Interface("value", value).Msg("got incomplete property assignment")
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
				// TODO: log error
				continue
			}
			for _, item := range properties {
				if len(item) != 1 {
					panic("wrong")
				}
				var def string
				for k := range item {
					def = k
					fmt.Println("added def " + def)
				}
				var propertyBody map[string]interface{}
				if err := json.Unmarshal([]byte(item[def]), &propertyBody); err != nil {
					// TODO: log error
					fmt.Println(err)
					continue
				}
				out, err := propertyAssignmentFromMap(def, propertyBody)
				if err != nil {
					// TODO: log error
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
