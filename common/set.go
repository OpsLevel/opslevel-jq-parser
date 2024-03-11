package common

import "encoding/json"

// Set enables deduplication of objects
type Set[T any] struct {
	data map[string]T
}

func NewSet[T any]() *Set[T] {
	return &Set[T]{
		data: make(map[string]T),
	}
}

// Add performs deduplication of objects in the Set by treating the json marshaled value as a unique key
func (set Set[T]) Add(object T) {
	marshaled, err := json.Marshal(&object)
	if err != nil {
		return
	}
	set.data[string(marshaled)] = object
}

func (set Set[T]) Values() []T {
	result := make([]T, len(set.data))
	i := 0
	for _, v := range set.data {
		result[i] = v
		i++
	}
	return result
}
