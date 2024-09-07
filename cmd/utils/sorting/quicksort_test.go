package sorting_test

import (
	"math/rand"
	"testing"

	"github.com/caffeine-addictt/waku/cmd/utils/sorting"
	"github.com/stretchr/testify/assert"
)

// Comparator function for ascending order
func asc(a, b int) bool {
	return a < b
}

// Comparator function for descending order
func desc(a, b int) bool {
	return a > b
}

func TestQuicksort(t *testing.T) {
	tt := []struct {
		cmp      func(a, b int) bool
		name     string
		input    []int
		expected []int
	}{
		{
			cmp:      asc,
			name:     "Ascending order",
			input:    []int{3, 1, 4, 1, 5, 9, 2, 6, 5, 3, 5},
			expected: []int{1, 1, 2, 3, 3, 4, 5, 5, 5, 6, 9},
		},
		{
			cmp:      desc,
			name:     "Descending order",
			input:    []int{3, 1, 4, 1, 5, 9, 2, 6, 5, 3, 5},
			expected: []int{9, 6, 5, 5, 5, 4, 3, 3, 2, 1, 1},
		},
		{
			cmp:      asc,
			name:     "Already sorted",
			input:    []int{1, 2, 3, 4, 5},
			expected: []int{1, 2, 3, 4, 5},
		},
		{
			cmp:      asc,
			name:     "Reverse sorted",
			input:    []int{5, 4, 3, 2, 1},
			expected: []int{1, 2, 3, 4, 5},
		},
		{
			cmp:      asc,
			name:     "Single element",
			input:    []int{42},
			expected: []int{42},
		},
		{
			name:     "Empty slice",
			cmp:      asc,
			input:    []int{},
			expected: []int{},
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			sorting.Quicksort(tc.input, tc.cmp)
			assert.Equal(t, tc.expected, tc.input)
		})
	}
}

func TestQuicksortASC(t *testing.T) {
	tt := []struct {
		input    []int
		expected []int
	}{
		{[]int{3, 1, 4, 1, 5, 9, 2, 6, 5, 3, 5}, []int{1, 1, 2, 3, 3, 4, 5, 5, 5, 6, 9}},
		{[]int{5, 4, 3, 2, 1}, []int{1, 2, 3, 4, 5}},
		{[]int{1, 2, 3, 4, 5}, []int{1, 2, 3, 4, 5}},
		{[]int{}, []int{}},
		{[]int{1}, []int{1}},
	}

	for _, tc := range tt {
		sorting.QuicksortASC(tc.input)
		assert.Equal(t, tc.expected, tc.input)
	}
}

func TestQuicksortDESC(t *testing.T) {
	tt := []struct {
		input    []int
		expected []int
	}{
		{[]int{3, 1, 4, 1, 5, 9, 2, 6, 5, 3, 5}, []int{9, 6, 5, 5, 5, 4, 3, 3, 2, 1, 1}},
		{[]int{1, 2, 3, 4, 5}, []int{5, 4, 3, 2, 1}},
		{[]int{5, 4, 3, 2, 1}, []int{5, 4, 3, 2, 1}},
		{[]int{}, []int{}},
		{[]int{1}, []int{1}},
	}

	for _, tc := range tt {
		sorting.QuicksortDESC(tc.input)
		assert.Equal(t, tc.expected, tc.input)
	}
}

func BenchmarkQuicksortSmall(b *testing.B) {
	arr := make([]int, 100)
	for i := range arr {
		arr[i] = rand.Intn(1000)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		sorting.Quicksort(arr, func(a, b int) bool { return a < b })
	}
}

func BenchmarkQuicksortLarge(b *testing.B) {
	arr := make([]int, 10000)
	for i := range arr {
		arr[i] = rand.Intn(100000)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		sorting.Quicksort(arr, func(a, b int) bool { return a < b })
	}
}
