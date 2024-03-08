package opslevel_jq_parser

import (
	"github.com/flant/libjq-go/pkg/jq"
	"github.com/rs/zerolog/log"

	"github.com/flant/libjq-go"
)

type JQFieldParser struct {
	program *jq.JqProgram
}

func NewJQFieldParser(expression string) JQFieldParser {
	// TODO: why is this here? I think this will happen with the System config bug.
	if expression == "" {
		expression = "empty"
	}
	// TODO: nil check on program?
	program, err := libjq_go.Jq().Program(expression).Precompile()
	if err != nil {
		log.Panic().Err(err).Str("expression", expression).Msg("error compiling jq expression")
	}
	return JQFieldParser{
		program: program,
	}
}

func (p JQFieldParser) Run(data string) (string, error) {
	// TODO: explain why we are not checking "" or "null" here.
	return p.program.RunRaw(data)
}
