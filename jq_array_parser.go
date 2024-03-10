package opslevel_jq_parser

import (
	"encoding/json"
	"github.com/opslevel/opslevel-jq-parser/v2024/common"
)

type JQArrayParser struct {
	programs []JQFieldParser
}

func NewJQArrayParser(expressions []string) JQArrayParser {
	programs := make([]JQFieldParser, len(expressions))
	for i, expression := range expressions {
		programs[i] = NewJQFieldParser(expression)
	}
	return JQArrayParser{
		programs: programs,
	}
}

func (p JQArrayParser) Run(data string) []string {
	output := make(common.UniqueMap[bool])
	for _, program := range p.programs {
		response := program.Run(data)
		if response == "" {
			continue
		}

		if common.Array(response) {
			var elements []string
			err := json.Unmarshal([]byte(response), &elements)
			if err != nil {
				continue
			}
			for _, elem := range elements {
				if elem == "" {
					continue
				}
				output.Add(response, true)
			}
			continue
		}
		output.Add(response, true)
	}
	return output.Keys()
}
