package opslevel_jq_parser

import (
	"encoding/json"
	"github.com/rs/zerolog/log"
	"strings"
)

type JQArrayParser []JQFieldParser

func NewJQArrayParser(expressions []string) JQArrayParser {
	programs := make([]JQFieldParser, len(expressions))
	for i, expression := range expressions {
		programs[i] = NewJQFieldParser(expression)
	}
	return programs
}

func (p JQArrayParser) Run(data string) []string {
	set := make(map[string]bool)
	output := make([]string, 0)
	for _, program := range p {
		response := program.Run(data)
		if _, ok := set[response]; ok || response == "" {
			continue
		}

		if strings.HasPrefix(response, "[") && strings.HasSuffix(response, "]") {
			var items []map[string]string
			err := json.Unmarshal([]byte(response), &items)
			if err != nil {
				log.Debug().Err(err).Msg("error on unmarshal array")
			}

			for _, item := range items {
				var m []byte
				m, err := json.Marshal(item)
				if err != nil {
					log.Debug().Err(err).Msg("error on marshal object")
				}
				response = string(m)
				// TODO: what about empty objects? Is this situation even possible?
				if _, ok := set[response]; ok || response == "{}" {
					continue
				}
				set[response] = true
				output = append(output, response)
			}
			continue
		}

		set[response] = true
		output = append(output, response)
	}
	return output
}

func jsonArray(s string) bool {
	return strings.HasPrefix(s, "[") && strings.HasSuffix(s, "]")
}
