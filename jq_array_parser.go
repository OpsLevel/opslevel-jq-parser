package opslevel_jq_parser

type JQArrayParser []JQFieldParser

func NewJQArrayParser(expressions []string) JQArrayParser {
	programs := make([]JQFieldParser, len(expressions))
	for i, expression := range expressions {
		programs[i] = NewJQFieldParser(expression)
	}
	return programs
}

func (p JQArrayParser) Run(data string) []string {
	set := make(map[string]bool)
	output := make([]string, 0)
	for _, program := range p {
		response := program.Run(data)
		if _, ok := set[response]; ok || response == "" {
			continue
		}
		set[response] = true
		output = append(output, response)
	}
	return output
}
