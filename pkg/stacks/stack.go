package stacks

import "fmt"

type Stack[T any] []T

func Of[T any](values []T) *Stack[T] {
	cp := make([]T, len(values))
	copy(cp, values)

	return (*Stack[T])(&cp)
}

func (s *Stack[T]) Len() int {
	return len(*s)
}

func (s *Stack[T]) Pop() T {
	item := (*s)[len(*s)-1]
	*s = append([]T(nil), (*s)[:len(*s)-1]...)
	return item
}

func (s *Stack[T]) PopN(n int) []T {
	items := make([]T, 0, n)
	for i := 0; i < n; i++ {
		items = append(items, s.Pop())
	}
	return items
}

func (s *Stack[T]) Push(items ...T) {
	*s = append(*s, items...)
}

func (s *Stack[T]) Unshift(items ...T) {
	*s = append(items, *s...)
}

func (s *Stack[T]) Append(items ...T) {
	*s = append(*s, items...)
}

func (s *Stack[T]) Shift() T {
	item := (*s)[0]
	*s = append([]T(nil), (*s)[1:]...)
	return item
}

func (s *Stack[T]) String() string {
	return fmt.Sprint(*s)
}
