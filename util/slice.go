package util

// Unique 按首次出现顺序去重，不修改原切片，元素按 Go 相等规则比较。
func Unique[T comparable](values []T) []T {
	seen := make(map[T]bool, len(values))
	result := make([]T, 0, len(values))
	for _, value := range values {
		if seen[value] {
			continue
		}
		seen[value] = true
		result = append(result, value)
	}
	return result
}
