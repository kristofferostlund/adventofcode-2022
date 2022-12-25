package slices

func Copy[V any](items []V) []V {
	cp := make([]V, len(items))
	copy(cp, items)
	return cp
}
