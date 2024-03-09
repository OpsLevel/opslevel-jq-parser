package opslevel_jq_parser

import (
	"encoding/json"
	"github.com/opslevel/opslevel-go/v2024"
	"github.com/opslevel/opslevel-jq-parser/v2024/orderedmap"
	"strings"
)

type JQTagsParser struct {
	creates []*JQFieldParser
	assigns []*JQFieldParser
}

func NewJQTagsParser(cfg TagRegistrationConfig) *JQTagsParser {
	creates := make([]*JQFieldParser, len(cfg.Create))
	for i, expression := range cfg.Create {
		creates[i] = NewJQFieldParser(expression)
	}
	assigns := make([]*JQFieldParser, len(cfg.Assign))
	for i, expression := range cfg.Assign {
		assigns[i] = NewJQFieldParser(expression)
	}
	return &JQTagsParser{
		creates: creates,
		assigns: assigns,
	}
}

// TODO: handle me
func IsObject(s string) bool {
	if strings.HasPrefix(s, "{") && strings.HasSuffix(s, "}") {
		return true
	}
	return false
}

// TODO: move me
func IsArray(s string) bool {
	if strings.HasPrefix(s, "[") && strings.HasSuffix(s, "]") {
		return true
	}
	return false
}

// TODO: move me
// TODO: comment case where this happens
// TODO: define interface?
func (p *JQTagsParser) handleObject(output *orderedmap.OrderedMap[opslevel.TagInput], toMap map[string]string) {
	for k, v := range toMap {
		tag := opslevel.TagInput{Key: k, Value: v}
		output.Add(tag.Key+tag.Value, tag)
	}
}

// TODO: return something compatible with tool registration?
// parse looks for JSON objects inside expression results and converts every key value pair into an opslevel.TagInput
func (p *JQTagsParser) parse(programs []*JQFieldParser, data string) []opslevel.TagInput {
	output := orderedmap.New[opslevel.TagInput]()
	for _, program := range programs {
		response, err := program.Run(data)
		if err != nil || response == "" {
			// TODO: log error
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
	return output.Values()
}

func (p *JQTagsParser) Run(data string) ([]opslevel.TagInput, []opslevel.TagInput, error) {
	return p.parse(p.creates, data), p.parse(p.assigns, data), nil
}
