package opslevel_jq_parser

import (
	"fmt"

	"github.com/flant/libjq-go"
	"github.com/flant/libjq-go/pkg/jq"
)

func appendEmptyExpr(expression string) string {
	if expression == "" {
		return "empty"
	}
	return expression + " // empty"
}

type JQFieldParser struct {
	program *jq.JqProgram
}

func NewJQFieldParser(expression string) *JQFieldParser {
	expression = appendEmptyExpr(expression)
	prg, err := libjq_go.Jq().Program(expression).Precompile()
	if err != nil {
		panic(fmt.Sprintf("unable to compile jq expression:  %s", expression))
	}
	return &JQFieldParser{
		program: prg,
	}
}

func (p *JQFieldParser) Run(data string) string {
	output, _ := p.program.RunRaw(data)
	return output
}
