package opslevel_jq_parser

import (
	"fmt"
	libjq_go "github.com/flant/libjq-go"
	"github.com/flant/libjq-go/pkg/jq"
	"github.com/rs/zerolog/log"
)

type JQFieldParser struct {
	program *jq.JqProgram
}

func NewJQFieldParser(expression string) JQFieldParser {
	if expression == "" {
		expression = "empty"
	} else {
		expression = fmt.Sprintf("%s // empty", expression)
	}
	program, err := libjq_go.Jq().Program(expression).Precompile()
	if err != nil {
		log.Panic().Err(err).Msg("error from libjq-go")
	}
	return JQFieldParser{
		program: program,
	}
}

func (p JQFieldParser) Run(data string) string {
	result, err := p.program.RunRaw(data)
	if err != nil {
		log.Debug().Err(err).Msg("error from libjq-go")
		return ""
	}
	return result
}
