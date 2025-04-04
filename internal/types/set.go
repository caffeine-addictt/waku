package types

import (
	"github.com/caffeine-addictt/waku/internal/config"
)

// A set implementation using map under the hood.
type Set[T comparable] map[T]struct{}

// NewSet creates a new set from a list of items
func NewSet[T comparable](items ...T) Set[T] {
	s := make(Set[T], len(items))
	for _, item := range items {
		s.Add(item)
	}
	return s
}

// Add an item to the set
func (s *Set[T]) Add(item T) {
	(*s)[item] = struct{}{}
}

// Remove an item from the set
func (s *Set[T]) Remove(item T) {
	delete(*s, item)
}

// Contains checks if an item is in the set
func (s *Set[T]) Contains(item T) bool {
	_, ok := (*s)[item]
	return ok
}

// Copy returns a new set containing all items in the set
func (s *Set[T]) Copy() Set[T] {
	n := make(Set[T], len(*s))
	for k := range *s {
		n[k] = struct{}{}
	}
	return n
}

// Count returns the number of items in the set
func (s *Set[T]) Len() int {
	return len(*s)
}

// ToSlice returns a slice of items in the set
func (s *Set[T]) ToSlice() []T {
	result := make([]T, 0, s.Len())
	for item := range *s {
		result = append(result, item)
	}
	return result
}

// Union returns a new set containing all items in both sets
func (s *Set[T]) Union(s2 Set[T]) Set[T] {
	s3 := NewSet[T]()
	for item := range *s {
		s3.Add(item)
	}
	for item := range s2 {
		s3.Add(item)
	}
	return s3
}

// Intersect returns a new set containing items in both sets
func (s *Set[T]) Intersect(s2 Set[T]) Set[T] {
	s3 := NewSet[T]()

	// Iterate over the smaller set
	small := (*s)
	big := s2
	if s2.Len() < s.Len() {
		small = s2
		big = (*s)
	}

	for item := range small {
		if big.Contains(item) {
			s3.Add(item)
		}
	}
	return s3
}

// Exclude returns a new set containing items in the first set but not in the second
func (s *Set[T]) Exclude(s2 Set[T]) Set[T] {
	s3 := NewSet[T]()
	for item := range *s {
		if !s2.Contains(item) {
			s3.Add(item)
		}
	}
	return s3
}

func (s *Set[T]) marshal(c config.ConfigType) ([]byte, error) { return c.Marshal(s.ToSlice()) }
func (s *Set[T]) unmarshal(c config.ConfigType, data []byte) error {
	var items []T
	if err := c.Unmarshal(data, &items); err != nil {
		return err
	}
	*s = NewSet(items...)
	return nil
}

func (s *Set[T]) UnmarshalJSON(data []byte) error { return s.unmarshal(config.JsonConfig{}, data) }
func (s Set[T]) MarshalJSON() ([]byte, error)     { return s.marshal(config.JsonConfig{}) }

func (s *Set[T]) UnmarshalYAML(data []byte) error { return s.unmarshal(config.YamlConfig{}, data) }
func (s Set[T]) MarshalYAML() ([]byte, error)     { return s.marshal(config.YamlConfig{}) }
