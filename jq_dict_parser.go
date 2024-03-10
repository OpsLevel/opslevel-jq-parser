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
	for key, expression := range p {
		jqRes := expression.Run(data)
		if jqRes == "" {
			continue
		}
		output[key] = jqRes
	}
	return output
}
