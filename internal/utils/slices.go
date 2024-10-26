package utils

// RemoveStructDuplicates 通用去重函数，接收一个元素数组和一个用于比较的函数
func RemoveStructDuplicates[T any, K comparable](elements []T, keyFunc func(T) K) []T {
	seen := make(map[K]struct{})
	var result []T

	for _, element := range elements {
		key := keyFunc(element)
		if _, exists := seen[key]; !exists {
			seen[key] = struct{}{}
			result = append(result, element)
		}
	}
	return result
}
