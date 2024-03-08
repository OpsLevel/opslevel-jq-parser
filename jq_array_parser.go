package opslevel_jq_parser

type JQArrayParser []JQFieldParser

func NewJQArrayParser(expressions []string) JQArrayParser {
	programs := make([]JQFieldParser, len(expressions))
	for i, expression := range expressions {
		programs[i] = NewJQFieldParser(expression)
	}
	return programs
}

func (p JQArrayParser) Run(data string) ([]string, error) {
	output := make([]string, 0, len(p))
	for _, program := range p {
		response, err := program.Run(data)
		if err != nil {
			continue
		}
		// TODO: explain why we are checking "" or "null" here.
		if response == "" || response == "null" {
			continue
		}
		output = append(output, response)
	}
	return output, nil
}

func (p JQArrayParser) RunDeduplicated(data string) ([]string, error) {
	output := make(map[string]bool)
	for _, program := range p {
		response, err := program.Run(data)
		if err != nil {
			continue
		}
		// TODO: explain why we are checking "" or "null" here.
		if response == "" || response == "null" {
			continue
		}
		output[response] = true
	}
	return Keys(output), nil
}

// Keys returns the keys of the map m.
// The keys will be in an indeterminate order.
func Keys[M ~map[K]V, K comparable, V any](m M) []K {
	r := make([]K, 0, len(m))
	for k := range m {
		r = append(r, k)
	}
	return r
}
