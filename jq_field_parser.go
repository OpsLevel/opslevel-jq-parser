package opslevel_jq_parser

import (
	"fmt"

	"github.com/flant/libjq-go"
	"github.com/flant/libjq-go/pkg/jq"
)

type JQFieldParser struct {
	program *jq.JqProgram
}

func NewJQFieldParser(expression string) *JQFieldParser {
	prg, err := libjq_go.Jq().Program(expression).Precompile()
	if err != nil {
		panic(fmt.Sprintf("unable to compile jq expression:  %s", expression))
	}
	return &JQFieldParser{
		program: prg,
	}
}

func (p *JQFieldParser) Run(data string) (string, error) {
	parsedData, err := p.program.RunRaw(data)
	if err != nil {
		return "", err
	}
	if parsedData == "null" {
		return "", nil
	}
	return parsedData, nil
}
