package u

import (
	"maps"
	"slices"
)

type Set[T comparable] struct {
	elements map[T]struct{}
}

func NewSet[T comparable](vals ...T) *Set[T] {
	set := &Set[T]{elements: make(map[T]struct{})}
	if len(vals) > 0 {
		set.Add(vals...)
	}
	return set
}

func (s *Set[T]) Add(val ...T) {
	for _, v := range val {
		s.elements[v] = struct{}{}
	}
}

func (s *Set[T]) Delete(val T) {
	delete(s.elements, val)
}

func (s *Set[T]) Has(val T) bool {
	_, found := s.elements[val]
	return found
}

func (s *Set[T]) Size() int {
	return len(s.elements)
}

func (s *Set[T]) ToSlice() []T {
	return slices.Collect(maps.Keys(s.elements))
}

func (s *Set[T]) Merge(s1 *Set[T]) {
	maps.Copy(s.elements, s1.elements)
}

func (s *Set[T]) Clone() *Set[T] {
	return &Set[T]{elements: maps.Clone(s.elements)}
}

func (s *Set[T]) Union(s1 *Set[T]) *Set[T] {
	clone := s.Clone()
	clone.Merge(s1)
	return clone
}

func (s *Set[T]) Intersection(s1 *Set[T]) *Set[T] {
	set := NewSet[T]()
	for key := range s1.elements {
		if s.Has(key) {
			set.Add(key)
		}
	}
	return set
}

func (s *Set[T]) Difference(s1 *Set[T]) *Set[T] {
	set := NewSet[T]()
	for key := range s1.elements {
		if !s.Has(key) {
			set.Add(key)
		}
	}
	return set
}

func (s *Set[T]) ForEach(f func(key T)) {
	for key := range s.elements {
		f(key)
	}
}

func (s *Set[T]) Map(f func(key T) T) *Set[T] {
	set := NewSet[T]()
	for key := range s.elements {
		newKey := f(key)
		set.Add(newKey)
	}
	return set
}

func (s *Set[T]) Filter(f func(key T) bool) *Set[T] {
	set := NewSet[T]()
	for key := range s.elements {
		ok := f(key)
		if ok {
			set.Add(key)
		}
	}
	return set
}
