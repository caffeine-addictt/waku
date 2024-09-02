package searching_test

import (
	"testing"

	"github.com/caffeine-addictt/template/cmd/utils/searching"
	"github.com/stretchr/testify/assert"
)

func TestBinarySearch(t *testing.T) {
	tt := []struct {
		cmp    func(int, int) bool
		arr    []int
		target int
		want   int
	}{
		{arr: []int{1, 2, 3, 4, 5}, target: 3, cmp: func(a, b int) bool { return a < b }, want: 2},
		{arr: []int{1, 2, 3, 4, 5}, target: 1, cmp: func(a, b int) bool { return a < b }, want: 0},
		{arr: []int{1, 2, 3, 4, 5}, target: 5, cmp: func(a, b int) bool { return a < b }, want: 4},
		{arr: []int{1, 2, 3, 4, 5}, target: 6, cmp: func(a, b int) bool { return a < b }, want: -1},
		{arr: []int{1, 2, 3, 4, 5}, target: 0, cmp: func(a, b int) bool { return a < b }, want: -1},
		{arr: []int{}, target: 1, cmp: func(a, b int) bool { return a < b }, want: -1},
	}

	for _, tc := range tt {
		got := searching.BinarySearch(tc.arr, tc.target, tc.cmp)
		assert.Equal(t, tc.want, got)
	}
}

func TestBinarySearchAuto(t *testing.T) {
	tt := []struct {
		arr    []int
		target int
		want   int
	}{
		{arr: []int{1, 2, 3, 4, 5}, target: 3, want: 2},
		{arr: []int{1, 2, 3, 4, 5}, target: 1, want: 0},
		{arr: []int{1, 2, 3, 4, 5}, target: 5, want: 4},
		{arr: []int{1, 2, 3, 4, 5}, target: 6, want: -1},
		{arr: []int{1, 2, 3, 4, 5}, target: 0, want: -1},
		{arr: []int{}, target: 1, want: -1},
	}

	for _, tc := range tt {
		got := searching.BinarySearchAuto(tc.arr, tc.target)
		assert.Equal(t, tc.want, got)
	}
}
