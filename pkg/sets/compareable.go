package sets

import (
	"fmt"
	"strings"
)

type Set[T comparable] struct {
	s map[T]struct{}
}

func newSet[T comparable]() Set[T] {
	return Set[T]{make(map[T]struct{}, 0)}
}

func Of[T comparable](values []T) Set[T] {
	set := newSet[T]()
	for _, v := range values {
		set.s[v] = struct{}{}
	}
	return set
}

func (s Set[T]) Intersection(other Set[T]) Set[T] {
	// Optimization to base the iteration of the smallest set.
	// Probably highly unnecessary.
	a, b := s, other
	if b.Len() < a.Len() {
		a, b = b, a
	}

	intersection := newSet[T]()
	for v := range a.s {
		if _, ok := b.s[v]; ok {
			intersection.s[v] = struct{}{}
		}
	}

	return intersection
}

func (s Set[T]) Contains(other Set[T]) bool {
	if s.Len() < other.Len() {
		return false
	}

	return s.Intersection(other).Len() == other.Len()
}

func (s Set[T]) Overlaps(other Set[T]) bool {
	return s.Intersection(other).Len() > 0
}

func (s Set[T]) Len() int {
	return len(s.s)
}

func (s Set[T]) Values() []T {
	out := make([]T, 0, len(s.s))
	for v := range s.s {
		out = append(out, v)
	}

	return out
}

func (s Set[T]) String() string {
	sb := &strings.Builder{}

	sb.WriteString("[")
	for i, v := range s.Values() {
		if i > 0 {
			sb.WriteString(", ")
		}
		sb.WriteString(fmt.Sprint(v))
	}
	sb.WriteString("]")

	return sb.String()
}
