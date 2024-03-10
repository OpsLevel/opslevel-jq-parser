package opslevel_jq_parser

import (
	"encoding/json"
	"github.com/opslevel/opslevel-jq-parser/v2024/common"
)

type JQArrayParser []JQFieldParser

func NewJQArrayParser(expressions []string) JQArrayParser {
	programs := make([]JQFieldParser, len(expressions))
	for i, expression := range expressions {
		programs[i] = NewJQFieldParser(expression)
	}
	return programs
}

type objectHandler[T any] func(*common.Set[T], string)

type stringHandler[T any] func(*common.Set[T], string)

func defaultObjectHandler[T any](output *common.Set[T], rawJSON string) {
	var object T
	err := json.Unmarshal([]byte(rawJSON), &object)
	if err != nil {
		return
	}
	output.Add(object)
}

func defaultStringHandler(output *common.Set[string], rawJSON string) {
	output.Add(rawJSON)
}

// parse will handle type T in different formats (object, array, string) using objectHandler and stringHandler
func parse[T any](output *common.Set[T], rawJSON string, objectHandler objectHandler[T], stringHandler stringHandler[T]) {
	if common.Object(rawJSON) {
		if objectHandler != nil {
			objectHandler(output, rawJSON)
		}
		return
	}
	if common.Array(rawJSON) {
		var array []any
		err := json.Unmarshal([]byte(rawJSON), &array)
		if err != nil {
			return
		}
		for _, item := range array {
			marshaled, err := json.Marshal(item)
			if err != nil {
				continue
			}
			parse(output, string(marshaled), objectHandler, stringHandler)
		}
		return
	}
	if stringHandler != nil {
		stringHandler(output, rawJSON)
	}
}

func run[T any](p JQArrayParser, data string, objectHandler objectHandler[T], stringHandler stringHandler[T]) []T {
	output := common.NewSet[T]()
	for _, program := range p {
		jqRes := program.ParseValue(data)
		if jqRes == "" {
			continue
		}
		parse[T](output, jqRes, objectHandler, stringHandler)
	}
	return output.Values()
}
