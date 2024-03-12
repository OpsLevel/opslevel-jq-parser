package opslevel_jq_parser

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/opslevel/opslevel-go/v2024"
	"github.com/rs/zerolog/log"
)

type RepositoryDTO struct {
	Name      string
	Directory string
	Repo      string
}

func (r *RepositoryDTO) Convert() (opslevel.ServiceRepositoryCreateInput, error) {
	var repo opslevel.ServiceRepositoryCreateInput
	if r.Repo == "" {
		return repo, fmt.Errorf("cannot create ServiceRepositoryCreateInput without repository alias")
	}
	return opslevel.ServiceRepositoryCreateInput{
		Repository:    *opslevel.NewIdentifier(r.Repo),
		BaseDirectory: opslevel.RefOf(r.Directory),
		DisplayName:   opslevel.RefOf(r.Name),
	}, nil
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

func (p *JQRepositoryParser) Run(data string) ([]opslevel.ServiceRepositoryCreateInput, error) {
	output := make([]opslevel.ServiceRepositoryCreateInput, 0, len(p.programs))
	for _, program := range p.programs {
		response, err := program.Run(data)
		// log.Warn().Msgf("expression: %s\nresponse: %s", program.program.Program, response)
		if err != nil {
			log.Warn().Msgf("unable to parse alias from expression: %s", program.program.Program)
			return nil, err
		}
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
					repoInput, err := repo.Convert()
					if err != nil {
						continue
					}
					output = append(output, repoInput)
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
				// TODO: log error
				log.Warn().Err(err).Msgf("unable to marshal repo expression: %s\n%s", program.program.Program, response)
				continue
			}
			repoInput, err := repo.Convert()
			if err != nil {
				continue
			}
			output = append(output, repoInput)
		} else {
			output = append(output, opslevel.ServiceRepositoryCreateInput{Repository: *opslevel.NewIdentifier(response)})
		}
	}
	return runJQUnique[opslevel.ServiceRepositoryCreateInput](output)
}
