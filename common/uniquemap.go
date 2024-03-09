package common

// UniqueMap is a simple wrapper around map[string]T
type UniqueMap[T any] map[string]T

// Add will return false and do nothing if the key is already set.
func (uMap UniqueMap[T]) Add(key string, value T) bool {
	if uMap.Contains(key) {
		return false
	}
	uMap[key] = value
	return true
}

func (uMap UniqueMap[T]) Contains(key string) bool {
	_, ok := uMap[key]
	return ok
}

func (uMap UniqueMap[T]) Keys() []string {
	result := make([]string, len(uMap))
	i := 0
	for k := range uMap {
		result[i] = k
		i++
	}
	return result
}

func (uMap UniqueMap[T]) Values() []T {
	result := make([]T, len(uMap))
	i := 0
	for _, v := range uMap {
		result[i] = v
		i++
	}
	return result
}
