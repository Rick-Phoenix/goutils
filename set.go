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

func (s *Set[T]) Equal(s1 *Set[T]) bool {
	return maps.Equal(s.elements, s1.elements)
}

func (s *Set[T]) DeleteFunc(filter func(T) bool) {
	for v := range s.elements {
		if filter(v) {
			delete(s.elements, v)
		}
	}
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

type StringSet struct {
	elements map[string]struct{}
}

func NewStringSet(vals ...string) *StringSet {
	set := &StringSet{elements: make(map[string]struct{})}
	if len(vals) > 0 {
		set.Add(vals...)
	}
	return set
}

func (s StringSet) Add(val ...string) {
	for _, v := range val {
		s.elements[v] = struct{}{}
	}
}

func (s StringSet) Delete(val string) {
	delete(s.elements, val)
}

func (s StringSet) Equal(s1 StringSet) bool {
	return maps.Equal(s.elements, s1.elements)
}

func (s StringSet) DeleteFunc(filter func(string) bool) {
	for v := range s.elements {
		if filter(v) {
			delete(s.elements, v)
		}
	}
}

func (s StringSet) Has(val string) bool {
	_, found := s.elements[val]
	return found
}

func (s StringSet) Size() int {
	return len(s.elements)
}

func (s StringSet) ToSlice() []string {
	return slices.Collect(maps.Keys(s.elements))
}

func (s StringSet) Merge(s1 StringSet) {
	maps.Copy(s.elements, s1.elements)
}

func (s StringSet) Clone() *StringSet {
	return &StringSet{elements: maps.Clone(s.elements)}
}

func (s StringSet) Union(s1 StringSet) *StringSet {
	clone := s.Clone()
	clone.Merge(s1)
	return clone
}

func (s StringSet) Intersection(s1 StringSet) *StringSet {
	set := NewStringSet()
	for key := range s1.elements {
		if s.Has(key) {
			set.Add(key)
		}
	}
	return set
}

func (s StringSet) Difference(s1 StringSet) *StringSet {
	set := NewStringSet()
	for key := range s1.elements {
		if !s.Has(key) {
			set.Add(key)
		}
	}
	return set
}

func (s StringSet) ForEach(f func(key string)) {
	for key := range s.elements {
		f(key)
	}
}

func (s StringSet) Map(f func(key string) string) *StringSet {
	set := NewStringSet()
	for key := range s.elements {
		newKey := f(key)
		set.Add(newKey)
	}
	return set
}

func (s StringSet) Filter(f func(key string) bool) *StringSet {
	set := NewStringSet()
	for key := range s.elements {
		ok := f(key)
		if ok {
			set.Add(key)
		}
	}
	return set
}

type IntSet struct {
	elements map[int]struct{}
}

func NewIntSet(vals ...int) *IntSet {
	set := &IntSet{elements: make(map[int]struct{})}
	if len(vals) > 0 {
		set.Add(vals...)
	}
	return set
}

func (s IntSet) Add(val ...int) {
	for _, v := range val {
		s.elements[v] = struct{}{}
	}
}

func (s IntSet) Delete(val int) {
	delete(s.elements, val)
}

func (s IntSet) Equal(s1 IntSet) bool {
	return maps.Equal(s.elements, s1.elements)
}

func (s IntSet) DeleteFunc(filter func(int) bool) {
	for v := range s.elements {
		if filter(v) {
			delete(s.elements, v)
		}
	}
}

func (s IntSet) Has(val int) bool {
	_, found := s.elements[val]
	return found
}

func (s IntSet) Size() int {
	return len(s.elements)
}

func (s IntSet) ToSlice() []int {
	return slices.Collect(maps.Keys(s.elements))
}

func (s IntSet) Merge(s1 IntSet) {
	maps.Copy(s.elements, s1.elements)
}

func (s IntSet) Clone() *IntSet {
	return &IntSet{elements: maps.Clone(s.elements)}
}

func (s IntSet) Union(s1 IntSet) *IntSet {
	clone := s.Clone()
	clone.Merge(s1)
	return clone
}

func (s IntSet) Intersection(s1 IntSet) *IntSet {
	set := NewIntSet()
	for key := range s1.elements {
		if s.Has(key) {
			set.Add(key)
		}
	}
	return set
}

func (s IntSet) Difference(s1 IntSet) *IntSet {
	set := NewIntSet()
	for key := range s1.elements {
		if !s.Has(key) {
			set.Add(key)
		}
	}
	return set
}

func (s IntSet) ForEach(f func(key int)) {
	for key := range s.elements {
		f(key)
	}
}

func (s IntSet) Map(f func(key int) int) *IntSet {
	set := NewIntSet()
	for key := range s.elements {
		newKey := f(key)
		set.Add(newKey)
	}
	return set
}

func (s IntSet) Filter(f func(key int) bool) *IntSet {
	set := NewIntSet()
	for key := range s.elements {
		ok := f(key)
		if ok {
			set.Add(key)
		}
	}
	return set
}
