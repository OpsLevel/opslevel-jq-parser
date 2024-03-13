package opslevel_jq_parser

import (
	"encoding/json"
	"strings"

	libjq_go "github.com/flant/libjq-go"

	"github.com/rs/zerolog/log"
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

func (p *JQArrayParser) Run(data string) ([]string, error) {
	output := make([]string, 0)
	for _, program := range p.programs {
		response, err := program.Run(data)
		if err != nil {
			log.Warn().Err(err).Msgf("jq execution error from expression: %s", program.program.Program)
			continue
		}
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
	return runJQUnique[string](output)
}

func runJQUnique[T any](inputArray []T) ([]T, error) {
	prg, err := libjq_go.Jq().Program("unique").Precompile()
	if err != nil {
		return nil, err
	}
	rawJSON, err := json.Marshal(inputArray)
	if err != nil {
		return nil, err
	}

	resultArray, err := prg.RunRaw(string(rawJSON))
	if err != nil {
		return nil, err
	}
	if resultArray == "null" {
		return nil, nil
	}

	var uniqueArray []T
	err = json.Unmarshal([]byte(resultArray), &uniqueArray)
	if err != nil {
		return nil, err
	}
	return uniqueArray, nil
}
