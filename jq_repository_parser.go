package opslevel_jq_parser

import (
	"encoding/json"

	"github.com/opslevel/opslevel-go/v2024"
	"github.com/opslevel/opslevel-jq-parser/v2024/common"
)

// RepositoryDTO is necessary to be able to set the identifier alias for the repository.
type RepositoryDTO struct {
	Name      string
	Directory string
	Repo      string
}

func (r RepositoryDTO) Convert() opslevel.ServiceRepositoryCreateInput {
	return opslevel.ServiceRepositoryCreateInput{
		Repository:    *opslevel.NewIdentifier(r.Repo),
		BaseDirectory: opslevel.RefOf(r.Directory),
		DisplayName:   opslevel.RefOf(r.Name),
	}
}

// repositoryObjectHandler will discard repositories that do not have the repository alias defined.
func repositoryObjectHandler(output *common.Set[RepositoryDTO], rawJSON string) {
	var repo RepositoryDTO
	err := json.Unmarshal([]byte(rawJSON), &repo)
	if err != nil || repo.Repo == "" {
		return
	}
	output.Add(repo)
}

// repositoryStringHandler will add a repository from a string by treating the string as a repository alias.
func repositoryStringHandler(output *common.Set[RepositoryDTO], rawJSON string) {
	repo := RepositoryDTO{Repo: rawJSON}
	output.Add(repo)
}

func RunRepositories(p JQArrayParser, data string) []opslevel.ServiceRepositoryCreateInput {
	dtos := run[RepositoryDTO](p, data, repositoryObjectHandler, repositoryStringHandler)
	repositories := make([]opslevel.ServiceRepositoryCreateInput, len(dtos))
	for i, dto := range dtos {
		repositories[i] = dto.Convert()
	}
	return repositories
}
