package ints

func Max(values ...int) int {
	if len(values) == 0 {
		panic("values must be provided")
	}

	max := values[0]
	for _, v := range values[1:] {
		if v > max {
			max = v
		}
	}
	return max
}

func Min(values ...int) int {
	if len(values) == 0 {
		panic("values must be provided")
	}

	min := values[0]
	for _, v := range values[1:] {
		if v < min {
			min = v
		}
	}
	return min
}
