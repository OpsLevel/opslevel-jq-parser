package opslevel_jq_parser

import (
	"github.com/opslevel/opslevel-go/v2024"
	"github.com/opslevel/opslevel-jq-parser/v2024/common"
)

// RepositoryDTO TODO: do we need a data transfer object here?
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

// repositoryStringHandler creates a RepositoryDTO by setting the identifier
func repositoryStringHandler(output *common.Set[RepositoryDTO], rawJSON string) {
	dto := RepositoryDTO{Name: rawJSON}
	output.Add(dto)
}

// TODO: bug in repositories....
// should not be outputting 2 repositories with the same alias. that doesn't make sense.
// in the test case, github.com:hashicorp/vault is duplicated.
// should change the UID to be this.
// TODO: look for other places this might happen.

func RunRepositories(p JQArrayParser, data string) []opslevel.ServiceRepositoryCreateInput {
	dtos := run[RepositoryDTO](p, data, defaultObjectHandler[RepositoryDTO], repositoryStringHandler)
	repositories := make([]opslevel.ServiceRepositoryCreateInput, len(dtos))
	for i, dto := range dtos {
		repositories[i] = dto.Convert()
	}
	return repositories
}
