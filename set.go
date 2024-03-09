package opslevel_jq_parser

type Set map[string]bool

func NewSet() Set {
	return make(map[string]bool)
}

func (set Set) Insert(val string) {
	set[val] = true
}

func (set Set) Contains(val string) bool {
	_, ok := set[val]
	return ok
}

func (set Set) Keys() []string {
	r := make([]string, 0, len(set))
	for k := range set {
		r = append(r, k)
	}
	return r
}
