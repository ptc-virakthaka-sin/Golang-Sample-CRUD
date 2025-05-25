package common

func RemoveComparable[T comparable](slice []T, element T) []T {
	var result []T
	for _, item := range slice {
		if item != element {
			result = append(result, item)
		}
	}
	return result
}
