package types_test

import (
	"strings"
	"testing"

	"github.com/caffeine-addictt/waku/internal/types"
	"github.com/goccy/go-json"
	"github.com/goccy/go-yaml"
	"github.com/stretchr/testify/assert"
)

func TestSetNewSet(t *testing.T) {
	set := types.NewSet(1, 2, 3)
	assert.Equal(t, 3, set.Len(), "expected length 3")
	assert.True(t, set.Contains(1), "expected set to contain 1")
	assert.True(t, set.Contains(2), "expected set to contain 2")
	assert.True(t, set.Contains(3), "expected set to contain 3")
}

func TestSetAdd(t *testing.T) {
	set := types.NewSet(1, 2, 3)
	set.Add(4)
	assert.Equal(t, 4, set.Len(), "expected length 4 after add")
	assert.True(t, set.Contains(4), "expected set to contain 4 after add")
}

func TestSetRemove(t *testing.T) {
	set := types.NewSet(1, 2, 3)
	set.Remove(2)
	assert.Equal(t, 2, set.Len(), "expected length 2 after remove")
	assert.False(t, set.Contains(2), "expected set not to contain 2 after remove")
}

func TestSetUnion(t *testing.T) {
	set := types.NewSet(1, 2, 3)
	otherSet := types.NewSet(3, 4, 5)
	unionSet := set.Union(otherSet)
	assert.Equal(t, 5, unionSet.Len(), "expected length 5 in union set")
	assert.True(t, unionSet.Contains(1), "expected union set to contain 1")
	assert.True(t, unionSet.Contains(2), "expected union set to contain 2")
	assert.True(t, unionSet.Contains(3), "expected union set to contain 3")
	assert.True(t, unionSet.Contains(4), "expected union set to contain 4")
	assert.True(t, unionSet.Contains(5), "expected union set to contain 5")
}

func TestSetIntersect(t *testing.T) {
	set := types.NewSet(1, 2, 3)
	otherSet := types.NewSet(2, 3, 4)
	intersectSet := set.Intersect(otherSet)
	assert.Equal(t, 2, intersectSet.Len(), "expected length 2 in intersect set")
	assert.True(t, intersectSet.Contains(2), "expected intersect set to contain 2")
	assert.True(t, intersectSet.Contains(3), "expected intersect set to contain 3")
}

func TestSetExclude(t *testing.T) {
	set := types.NewSet(1, 2, 3)
	otherSet := types.NewSet(2, 3)
	excludeSet := set.Exclude(otherSet)
	assert.Equal(t, 1, excludeSet.Len(), "expected length 1 in exclude set")
	assert.True(t, excludeSet.Contains(1), "expected exclude set to contain 1")
}

func TestSetToSlice(t *testing.T) {
	set := types.NewSet(1, 2, 3)
	assert.ElementsMatch(t, []int{1, 2, 3}, set.ToSlice(), "expected slice to match")
}

func TestSetMarshalJSON(t *testing.T) {
	set := types.NewSet(1, 2, 3)
	data, err := json.Marshal(set)
	assert.NoError(t, err, "unexpected error during Marshal")

	assert.ElementsMatch(t, []string{"1", "2", "3"}, strings.Split(strings.TrimSuffix(strings.TrimPrefix(string(data), "["), "]"), ","), "expected JSON output to match")
}

func TestSetUnmarshalJSON(t *testing.T) {
	var newSet types.Set[int]
	data := []byte(`[1,2,3]`)
	err := json.Unmarshal(data, &newSet)
	assert.NoError(t, err, "unexpected error during Unmarshal")
	assert.Equal(t, 3, newSet.Len(), "expected length 3 after unmarshal")
	assert.True(t, newSet.Contains(1), "expected set to contain 1 after unmarshal")
	assert.True(t, newSet.Contains(2), "expected set to contain 2 after unmarshal")
	assert.True(t, newSet.Contains(3), "expected set to contain 3 after unmarshal")
}

func TestSetMarshalYAML(t *testing.T) {
	set := types.NewSet(1, 2, 3)
	data, err := yaml.Marshal(set)
	assert.NoError(t, err, "unexpected error during Marshal")

	assert.ElementsMatch(t, []string{"1", "2", "3", ""}, strings.Split(strings.ReplaceAll(string(data), "- ", ""), "\n"), "expected YAML output to match")
}

func TestSetUnmarshalYAML(t *testing.T) {
	var newSet types.Set[int]
	data := []byte(`[1,2,3]`)
	err := yaml.Unmarshal(data, &newSet)
	assert.NoError(t, err, "unexpected error during Unmarshal")
	assert.Equal(t, 3, newSet.Len(), "expected length 3 after unmarshal")
	assert.True(t, newSet.Contains(1), "expected set to contain 1 after unmarshal")
	assert.True(t, newSet.Contains(2), "expected set to contain 2 after unmarshal")
	assert.True(t, newSet.Contains(3), "expected set to contain 3 after unmarshal")
}
