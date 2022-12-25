package maps

func Copy[K comparable, V any](m map[K]V) map[K]V {
	cp := make(map[K]V, len(m))
	for k, v := range m {
		cp[k] = v
	}
	return cp
}
