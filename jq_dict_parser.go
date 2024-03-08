package opslevel_jq_parser

type JQDictParser map[string]JQFieldParser

func NewJQDictParser(dict map[string]string) map[string]JQFieldParser {
	output := make(map[string]JQFieldParser)
	for key, expression := range dict {
		output[key] = NewJQFieldParser(expression)
	}
	return output
}

func (p JQDictParser) Run(data string) (map[string]string, error) {
	output := make(map[string]string)
	for key, program := range p {
		response, err := program.Run(data)
		if err != nil {
			continue
		}
		// TODO: explain why we are not checking "".
		if response == "null" {
			continue
		}
		output[key] = response
	}
	return output, nil
}
