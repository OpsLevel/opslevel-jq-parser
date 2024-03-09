package opslevel_jq_parser

import (
	"encoding/json"
	"strings"
)

type JQArrayParser struct {
	programs []*JQFieldParser
}

func NewJQArrayParser(expressions []string) *JQArrayParser {
	programs := make([]*JQFieldParser, len(expressions))
	for i, expression := range expressions {
		programs[i] = NewJQFieldParser(expression)
	}
	return &JQArrayParser{
		programs: programs,
	}
}

func (p *JQArrayParser) Run(data string) []string {
	output := make([]string, 0, len(p.programs))
	for _, program := range p.programs {
		response := program.Run(data)
		if response == "" {
			continue
		}
		if strings.HasPrefix(response, "[") && strings.HasSuffix(response, "]") {
			var aliases []string
			if err := json.Unmarshal([]byte(response), &aliases); err == nil {
				for _, alias := range aliases {
					if alias == "" {
						continue
					}
					output = append(output, alias)
				}
			}
		} else {
			output = append(output, response)
		}
	}
	return output
}
