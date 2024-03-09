package opslevel_jq_parser

import (
	"encoding/json"
	"fmt"
	"github.com/mitchellh/mapstructure"
	"github.com/opslevel/opslevel-go/v2024"
	"github.com/opslevel/opslevel-jq-parser/v2024/common"
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

func (p *JQToolsParser) handleObject(output common.UniqueMap[opslevel.ToolCreateInput], toMap map[string]string) {
	var tool opslevel.ToolCreateInput
	err := mapstructure.Decode(toMap, &tool)
	if err != nil {
		fmt.Println(err)
		return
	}
	output.Add(fmt.Sprintf("%s%s%v", tool.Category, tool.DisplayName, tool.Environment), tool)
}

func (p *JQToolsParser) Run(data string) ([]opslevel.ToolCreateInput, error) {
	output := make(common.UniqueMap[opslevel.ToolCreateInput])
	for _, program := range p.programs {
		response, err := program.Run(data)
		if err != nil || response == "" {
			log.Warn().Msgf("unable to parse alias from expression: %s", program.program.Program)
			continue
		}

		if common.Object(response) {
			var toMap map[string]string
			err = json.Unmarshal([]byte(response), &toMap)
			if err != nil {
				continue
			}
			p.handleObject(output, toMap)
		} else if common.Array(response) {
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
