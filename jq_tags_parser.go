package opslevel_jq_parser

import (
	"encoding/json"
)

type JQTagsParser JQArrayParser

func NewJQTagsParser(expressions []string) JQTagsParser {
	return JQTagsParser(NewJQArrayParser(expressions))
}

func (p JQTagsParser) Run(data string) []string {
	set := NewSet()
	for _, program := range p {
		response := program.Run(data)
		if _, ok := set[response]; ok || response == "" {
			continue
		}

		var toMap map[string]string
		err := json.Unmarshal([]byte(response), &toMap)
		if err != nil {
			continue
		}
		for k, v := range toMap {
			tag := map[string]string{"key": k, "value": v}
			marshaled, err := json.Marshal(tag)
			if err != nil {
				continue
			}
			set.Insert(string(marshaled))
		}
	}
	return set.Keys()
}
