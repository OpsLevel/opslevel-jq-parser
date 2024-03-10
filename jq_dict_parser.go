package opslevel_jq_parser

type JQDictParser map[string]JQFieldParser

func NewJQDictParser(dict map[string]string) map[string]JQFieldParser {
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
		if jqRes == "null" {
			// in the case that the expression returned nothing (happens in the case where the key was not found)
			// jq will return "null". This is not the same as empty string. So in that case, skip the item.
			continue
		}
		output[key] = jqRes
	}
	return output
}
