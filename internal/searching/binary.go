package searching

import c "cmp"

// BinarySearch searches for the target in the array using binary search
func BinarySearch[T comparable](arr []T, target T, cmp func(T, T) bool) int {
	left, right := 0, len(arr)-1

	for left <= right {
		mid := left + (right-left)/2

		if arr[mid] == target {
			return mid
		}

		if cmp(arr[mid], target) {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return -1 // target not found
}

// BinarySearchAuto is a shorthand for BinarySearch that uses the default comparator
func BinarySearchAuto[T c.Ordered](arr []T, target T) int {
	return BinarySearch(arr, target, func(a, b T) bool {
		return a < b
	})
}
