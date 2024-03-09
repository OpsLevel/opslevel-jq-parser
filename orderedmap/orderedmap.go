package orderedmap

// OrderedMap maintains the order of insertion of its elements. Is write-only, does not support deletion.
type OrderedMap[T any] struct {
	hashMap map[string]T
	keys    []string
}

func New[T any]() *OrderedMap[T] {
	return &OrderedMap[T]{
		hashMap: make(map[string]T),
		keys:    make([]string, 0),
	}
}

// Add will do nothing and return false if the key is already set.
func (thisMap *OrderedMap[T]) Add(key string, value T) bool {
	if thisMap.Contains(key) {
		return false
	}
	thisMap.hashMap[key] = value
	thisMap.keys = append(thisMap.keys, key)
	return true
}

func (thisMap *OrderedMap[T]) Contains(key string) bool {
	_, ok := thisMap.hashMap[key]
	return ok
}

// Keys returns the list of keys in order of insertion
func (thisMap *OrderedMap[T]) Keys() []string {
	return thisMap.keys
}

// Values returns the list of values in order of insertion
func (thisMap *OrderedMap[T]) Values() []T {
	result := make([]T, len(thisMap.keys))
	for i, k := range thisMap.keys {
		result[i] = thisMap.hashMap[k]
	}
	return result
}

func (thisMap *OrderedMap[T]) Len() int {
	return len(thisMap.keys)
}
