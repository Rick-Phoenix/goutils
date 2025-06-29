package u

import (
	"maps"
	"slices"
)

func CopyMaps[M ~map[K]V, K comparable, V any](dst M, sources ...M) M {
	out := make(M)
	for _, m := range sources {
		maps.Copy(out, m)
	}

	return out
}

func MapKeys[T comparable](m map[T]any) []T {
	iter := maps.Keys(m)
	keys := slices.Collect(iter)

	return keys
}

func MapValues[T comparable](m map[any]T) []T {
	iter := maps.Values(m)
	values := slices.Collect(iter)

	return values
}

type Entry[K, V any] struct {
	Key   K
	Value V
}

func MapEntries[K comparable, V any](m map[K]V) []Entry[K, V] {
	iter := maps.All(m)
	var entries []Entry[K, V]

	for k, v := range iter {
		entries = append(entries, Entry[K, V]{Key: k, Value: v})
	}

	return entries
}
