package maps

func LookupOf[V any, K comparable](items []V, fn func(item V) K) map[K]V {
	lookup := make(map[K]V, len(items))
	for _, v := range items {
		lookup[fn(v)] = v
	}
	return lookup
}
