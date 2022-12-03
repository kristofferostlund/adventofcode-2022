package sets

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
	intersection := newSet[T]()
	for v := range s.s {
		if _, ok := other.s[v]; ok {
			intersection.s[v] = struct{}{}
		}
	}
	return intersection
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
