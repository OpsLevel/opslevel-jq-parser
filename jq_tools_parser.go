package opslevel_jq_parser

import (
	"encoding/json"
	"fmt"
	"github.com/mitchellh/mapstructure"
	"github.com/opslevel/opslevel-go/v2024"
	"github.com/opslevel/opslevel-jq-parser/v2024/orderedmap"
	"github.com/rs/zerolog/log"
)

type JQToolsParser struct {
	programs []*JQFieldParser
}

func NewJQToolsParser(expressions []string) *JQToolsParser {
	programs := make([]*JQFieldParser, len(expressions))
	for i, expression := range expressions {
		programs[i] = NewJQFieldParser(expression)
	}
	return &JQToolsParser{
		programs: programs,
	}
}

// TODO: test me
// TODO: move me
func MapHasKeys[T any](m map[string]T, keys ...string) bool {
	for _, k := range keys {
		if _, ok := m[k]; !ok {
			return false
		}
	}
	return true
}

// TODO: move me
// TODO: comment case where this happens
// TODO: define interface?
func (p *JQToolsParser) handleObject(output *orderedmap.OrderedMap[opslevel.ToolCreateInput], toMap map[string]string) {
	var tool opslevel.ToolCreateInput
	err := mapstructure.Decode(toMap, &tool)
	if err != nil {
		fmt.Println(err)
		return
	}
	output.Add(fmt.Sprintf("%s%s%v", tool.Category, tool.DisplayName, tool.Environment), tool)
}

// TODO: this does not need an error...
func (p *JQToolsParser) Run(data string) ([]opslevel.ToolCreateInput, error) {
	output := orderedmap.New[opslevel.ToolCreateInput]()
	for _, program := range p.programs {
		response, err := program.Run(data)
		if err != nil || response == "" {
			log.Warn().Msgf("unable to parse alias from expression: %s", program.program.Program)
			continue
		}

		if IsObject(response) {
			var toMap map[string]string
			err = json.Unmarshal([]byte(response), &toMap)
			if err != nil {
				continue
			}
			p.handleObject(output, toMap)
		} else if IsArray(response) {
			var toSlice []map[string]string
			err = json.Unmarshal([]byte(response), &toSlice)
			if err != nil {
				continue
			}
			for _, item := range toSlice {
				p.handleObject(output, item)
			}
		}
	}
	return output.Values(), nil
}
