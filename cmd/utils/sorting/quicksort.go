package sorting

import (
	c "cmp"
	"sync"
)

// Threshold defines the threshold at which the algorithm switches
// to partition in parallel/concurrent mode
const Threshold = 200

// Quicksort sorts the given array using the quicksort algorithm
// in concurrent/parallel mode, in-place.
//
//	best | worst | avg | effective
//	time 0(nlog(n)) |  0(n^2) | 0(nlog(n))
//	space 0(log(n)) | 0(1) | 0(log(n))
//	effective space 0(nlog(n)) avg | 0(log(n) + sqrt(n)) - due to concurrency overhead
func Quicksort[T comparable](arr []T, cmp func(v, pivot T) bool) {
	var wg sync.WaitGroup
	wg.Add(1)
	parallelQuicksort(arr, 0, len(arr)-1, cmp, &wg)
	wg.Wait()
}

// QuicksortASC invokes Quicksort in ascending order
func QuicksortASC[T c.Ordered](arr []T) {
	Quicksort(arr, func(v, pivot T) bool {
		return v < pivot
	})
}

// QuicksortDESC invokes Quicksort in descending order
func QuicksortDESC[T c.Ordered](arr []T) {
	Quicksort(arr, func(v, pivot T) bool {
		return v > pivot
	})
}

func parallelQuicksort[T comparable](arr []T, left, right int, cmp func(v, pivot T) bool, wg *sync.WaitGroup) {
	if wg != nil {
		defer wg.Done()
	}

	if left >= right {
		return
	}

	pivotIndex := partition(arr, left, right, cmp)

	// Run concurrently if more than threshold
	if (right - left) > Threshold {
		wg.Add(1)
		go parallelQuicksort(arr, left, pivotIndex-1, cmp, wg)

		wg.Add(1)
		go parallelQuicksort(arr, pivotIndex+1, right, cmp, wg)
	} else {
		parallelQuicksort(arr, left, pivotIndex-1, cmp, nil)
		parallelQuicksort(arr, pivotIndex+1, right, cmp, nil)
	}
}

// partition rearranges elements based on the pivot
func partition[T comparable](arr []T, left, right int, cmp func(v, pivot T) bool) int {
	pivotIndex := left + (right-left)/2
	pivot := arr[pivotIndex]

	// Move pivot to the end
	arr[pivotIndex], arr[right] = arr[right], arr[pivotIndex]

	storeIndex := left
	for i := left; i < right; i++ {
		if cmp(arr[i], pivot) {
			arr[i], arr[storeIndex] = arr[storeIndex], arr[i]
			storeIndex++
		}
	}

	// Move pivot to its final place
	arr[storeIndex], arr[right] = arr[right], arr[storeIndex]
	return storeIndex
}
