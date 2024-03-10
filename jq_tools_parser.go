package opslevel_jq_parser

import (
	"github.com/opslevel/opslevel-go/v2024"
)

func RunTools(p JQArrayParser, data string) []opslevel.ToolCreateInput {
	return run[opslevel.ToolCreateInput](p, data, defaultObjectHandler[opslevel.ToolCreateInput], nil)
}
