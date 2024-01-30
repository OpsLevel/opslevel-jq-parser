package opslevel_jq_parser

import (
	"fmt"

	"github.com/flant/libjq-go"
	"github.com/flant/libjq-go/pkg/jq"
	"github.com/opslevel/opslevel-go/v2024"
)

type JQFieldParser struct {
	program *jq.JqProgram
}

func NewJQFieldParser(expression string) *JQFieldParser {
	if expression == "" {
		expression = "empty"
	}
	prg, err := libjq_go.Jq().Program(expression).Precompile()
	if err != nil {
		panic(fmt.Sprintf("unable to compile jq expression:  %s", expression))
	}
	return &JQFieldParser{
		program: prg,
	}
}

func (p *JQFieldParser) Run(data string) (opslevel.JsonString, error) {
	jqResult, err := p.program.RunRaw(data)
	return opslevel.JsonString(jqResult), err
}
