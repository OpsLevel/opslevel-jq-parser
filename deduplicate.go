package opslevel_jq_parser

import (
	"fmt"

	"github.com/opslevel/opslevel-go/v2024"
)

func Deduplicated[T any](objects []T, keyFunc func(object T) string) []T {
	out := make([]T, 0)
	set := make(map[string]struct{})
	for _, obj := range objects {
		key := keyFunc(obj)
		if _, ok := set[key]; ok {
			continue
		}
		set[key] = struct{}{}
		out = append(out, obj)
	}
	return out
}

func DeduplicatedTools(objects []opslevel.ToolCreateInput) []opslevel.ToolCreateInput {
	return Deduplicated(objects, func(tool opslevel.ToolCreateInput) string {
		toolEnv := ""
		if tool.Environment != nil {
			toolEnv = *tool.Environment
		}
		return fmt.Sprintf("%s%s%s", tool.Category, tool.DisplayName, toolEnv)
	})
}
