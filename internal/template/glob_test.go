package template_test

import (
	"fmt"
	"testing"

	"github.com/caffeine-addictt/waku/internal/template"
	"github.com/caffeine-addictt/waku/internal/types"
	"github.com/stretchr/testify/assert"
)

func TestResolveGlobs(t *testing.T) {
	tests := []struct {
		paths    types.Set[string]
		ignores  types.Set[string]
		expected types.Set[string]
		name     string
	}{
		{
			paths:    types.NewSet("path/to/file"),
			ignores:  types.NewSet("path/to/file"),
			expected: types.NewSet[string](),
			name:     "ignore single file",
		},
		{
			paths:    types.NewSet("path/to/file"),
			ignores:  types.NewSet("path/to/f*"),
			expected: types.NewSet[string](),
			name:     "ignore single file glob",
		},
		{
			paths:    types.NewSet("path/to/file"),
			ignores:  types.NewSet("path/to/*"),
			expected: types.NewSet[string](),
			name:     "ignore single dir level glob",
		},
		{
			paths:    types.NewSet("path/to/file"),
			ignores:  types.NewSet("path/**"),
			expected: types.NewSet[string](),
			name:     "ignore recursive dir level glob",
		},
		{
			paths:    types.NewSet("path/one", "path/two", "path/three"),
			ignores:  types.NewSet("*"),
			expected: types.NewSet[string](),
			name:     "ignore all files",
		},
		{
			paths:    types.NewSet("path/to/one", "path/to/two", "path/to/three"),
			ignores:  types.NewSet("path/to/*", "!path/to/one", "!path/to/two"),
			expected: types.NewSet("path/to/one", "path/to/two"),
			name:     "include negated files",
		},
		{
			paths:    types.NewSet("path/to/one"),
			ignores:  types.NewSet("!path/to/o*", "path/to/one"),
			expected: types.NewSet("path/to/one"),
			name:     "include negated file glob",
		},
		{
			paths:    types.NewSet("path/to/one", "path/to/two", "path/to/three"),
			ignores:  types.NewSet("!path/to/*", "path/**"),
			expected: types.NewSet("path/to/one", "path/to/two", "path/to/three"),
			name:     "include negated dir level glob",
		},
		{
			paths:    types.NewSet("path/to/one", "path/to/two", "path/to/three"),
			ignores:  types.NewSet("!path/**", "path/to/*"),
			expected: types.NewSet("path/to/one", "path/to/two", "path/to/three"),
			name:     "include negated recursive dir level glob",
		},
		{
			paths:    types.NewSet("path/to/one.+?^$()[]{}|\\"),
			ignores:  types.NewSet("path/to/one.+?^$()[]{}|\\"),
			expected: types.NewSet[string](),
			name:     "regex characters are escaped",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result := template.ResolveGlobs(tc.paths, tc.ignores)
			assert.ElementsMatch(t, tc.expected.ToSlice(), result.ToSlice())
		})
	}
}

func BenchmarkResolveIncludes(b *testing.B) {
	for _, tc := range []struct {
		paths   types.Set[string]
		ignores types.Set[string]
		name    string
	}{
		{
			name:    "Simple Exclude",
			paths:   types.NewSet("path/to/file1", "path/to/file2", "path/to/dir/file3"),
			ignores: types.NewSet("path/to/file1"),
		},
		{
			name:    "Single Level Glob Exclude",
			paths:   types.NewSet("path/to/file1", "path/to/file2", "path/to/dir/file3"),
			ignores: types.NewSet("path/to/dir/*"),
		},
		{
			name:    "Recursive Glob Exclude",
			paths:   types.NewSet("path/to/file1", "path/to/file2", "path/to/dir/file3", "path/to/dir/subdir/file4"),
			ignores: types.NewSet("path/to/dir/**"),
		},
		{
			name:    "Negation of Recursive Glob",
			paths:   types.NewSet("path/to/file1", "path/to/file2", "path/to/dir/file3", "path/to/dir/subdir/file4"),
			ignores: types.NewSet("path/to/dir/**", "!path/to/dir/subdir/file4"),
		},
		{
			name:    "Complex Pattern",
			paths:   types.NewSet("path/to/file1", "path/to/file2", "path/to/dir/file3", "path/to/dir/subdir/file4", "path/to/otherfile"),
			ignores: types.NewSet("path/to/dir/**", "!path/to/dir/subdir/file4", "path/to/otherfile"),
		},
		{
			name:    "Large Dataset",
			paths:   generateLargeDataset(),
			ignores: types.NewSet("path/to/dir/**"),
		},
	} {
		b.Run(tc.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				template.ResolveGlobs(tc.paths, tc.ignores)
			}
		})
	}
}

// Helper function to generate a large dataset for benchmarking
func generateLargeDataset() types.Set[string] {
	paths := make([]string, 0, 1000)
	for i := 0; i < 1000; i++ {
		paths = append(paths, fmt.Sprintf("path/to/dir/file%d", i))
	}
	return types.NewSet(paths...)
}
