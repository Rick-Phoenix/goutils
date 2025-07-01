package u

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"slices"
)

func Dedupe[T comparable](s []T) []T {
	seen := make(map[T]struct{})
	var uniqueItems []T

	for _, v := range s {
		if _, exists := seen[v]; !exists {
			seen[v] = struct{}{}
			uniqueItems = append(uniqueItems, v)
		}
	}

	return uniqueItems
}

func DedupeNonComp[T any](s []T) []T {
	seen := make(map[string]struct{})
	uniqueItems := []T{}

	for _, item := range s {
		var buf bytes.Buffer
		encoder := gob.NewEncoder(&buf)

		err := encoder.Encode(item)
		if err != nil {
			fmt.Printf("Error encoding item %v with gob: %v\n", item, err)
			continue
		}

		fingerprint := buf.String()

		if _, exists := seen[fingerprint]; !exists {
			seen[fingerprint] = struct{}{}
			uniqueItems = append(uniqueItems, item)
		}
	}

	return uniqueItems
}

func FilterAndDedupe[T comparable](target []T, filter func(T) bool) []T {
	seen := make(map[T]struct{})
	out := []T{}

	for _, i := range target {
		if _, alreadySeen := seen[i]; alreadySeen {
			continue
		}

		seen[i] = struct{}{}

		if filter(i) {
			out = append(out, i)
		}
	}

	return out
}

func FindItem[T any](s []T, filter func(i T) bool) *T {
	idx := slices.IndexFunc(s, filter)
	var item *T

	if idx != -1 {
		item = &s[idx]
	}

	return item
}

func ToPtrSlice[T any](s []T) []*T {
	out := make([]*T, len(s))
	for i, v := range s {
		out[i] = &v
	}

	return out
}

func ToValSlice[T any](s []*T) []T {
	out := make([]T, len(s))
	for i, v := range s {
		out[i] = *v
	}

	return out
}

func JoinSlice[T any](s []T, separator string) string {
	out := ""

	for i, n := range s {
		out += fmt.Sprintf("%+v", n)
		if i != len(s)-1 {
			out += separator
		}
	}

	return out
}

func SliceIntersects[T comparable](s1 []T, s2 []T) bool {
	set := NewSet(s1...)
	set.Add(s2...)
	return set.Size() > 0
}
