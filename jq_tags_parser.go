package opslevel_jq_parser

import (
	"encoding/json"
	"github.com/opslevel/opslevel-go/v2024"
	"github.com/opslevel/opslevel-jq-parser/v2024/common"
)

// tagObjectHandler will treat each key value pair in the object as an opslevel.TagInput
func tagObjectHandler(output *common.Set[opslevel.TagInput], rawJSON string) {
	var toMap map[string]string
	err := json.Unmarshal([]byte(rawJSON), &toMap)
	if err != nil {
		return
	}
	for k, v := range toMap {
		tag := opslevel.TagInput{Key: k, Value: v}
		output.Add(tag)
	}
}

func RunTags(p JQArrayParser, data string) []opslevel.TagInput {
	return run[opslevel.TagInput](p, data, tagObjectHandler, nil)
}
