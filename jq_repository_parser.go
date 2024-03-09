package opslevel_jq_parser

import (
	"encoding/json"
	"strings"

	"github.com/opslevel/opslevel-go/v2024"
)

type RepositoryDTO struct {
	Name      string
	Directory string
	Repo      string
}

func (r *RepositoryDTO) Convert() opslevel.ServiceRepositoryCreateInput {
	return opslevel.ServiceRepositoryCreateInput{
		Repository:    *opslevel.NewIdentifier(r.Repo),
		BaseDirectory: opslevel.RefOf(r.Directory),
		DisplayName:   opslevel.RefOf(r.Name),
	}
}

type JQRepositoryParser struct {
	programs []*JQFieldParser
}

func NewJQRepositoryParser(expressions []string) *JQRepositoryParser {
	programs := make([]*JQFieldParser, len(expressions))
	for i, expression := range expressions {
		programs[i] = NewJQFieldParser(expression)
	}
	return &JQRepositoryParser{
		programs: programs,
	}
}

func (p *JQRepositoryParser) Run(data string) []opslevel.ServiceRepositoryCreateInput {
	output := make([]opslevel.ServiceRepositoryCreateInput, 0, len(p.programs))
	for _, program := range p.programs {
		response := program.Run(data)
		if response == "" {
			continue
		}
		if strings.HasPrefix(response, "[") && strings.HasSuffix(response, "]") {
			if response == "[]" {
				continue
			}
			var repos []RepositoryDTO
			if err := json.Unmarshal([]byte(response), &repos); err == nil {
				for _, repo := range repos {
					output = append(output, repo.Convert())
				}
			} else {
				// Try as []string
				var repoNames []string
				if err := json.Unmarshal([]byte(response), &repoNames); err == nil {
					for _, repoName := range repoNames {
						output = append(output, opslevel.ServiceRepositoryCreateInput{Repository: *opslevel.NewIdentifier(repoName)})
					}
				}
			}
		} else if strings.HasPrefix(response, "{") && strings.HasSuffix(response, "}") {
			var repo RepositoryDTO
			if err := json.Unmarshal([]byte(response), &repo); err != nil {
				continue
			}
			output = append(output, repo.Convert())
		} else {
			output = append(output, opslevel.ServiceRepositoryCreateInput{Repository: *opslevel.NewIdentifier(response)})
		}

	}
	return output
}
