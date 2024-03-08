package opslevel_jq_parser

type JQDictParser map[string]JQFieldParser

func NewJQDictParser(dict map[string]string) JQDictParser {
	output := make(map[string]JQFieldParser)
	for key, expression := range dict {
		output[key] = NewJQFieldParser(expression)
	}
	return output
}

func (p JQDictParser) Run(data string) map[string]string {
	output := make(map[string]string)
	for key, program := range p {
		response := program.Run(data)
		if response == "" {
			continue
		}
		output[key] = response
	}
	return output
}
