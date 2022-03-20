package fun

// Map returns the slice obtained after applying the given function over every
// element in the given slice
func Map[T1, T2 any](elems []T1, fn func(T1) T2) []T2 {
	ret := make([]T2, 0)
	for _, elem := range elems {
		ret = append(ret, fn(elem))
	}
	return ret
}

// Filter returns the slice obtained after retaining only those elements
// in the given slice for which the given function returns true
func Filter[T any](elems []T, fn func(T) bool) []T {
	ret := make([]T, 0)
	for _, elem := range elems {
		if fn(elem) {
			ret = append(ret, elem)
		}
	}
	return ret
}

// Take returns the slice obtained after taking the first n elements from the
// given slice.
// If n is greater than the length of the slice, return the entire slice
func Take[T any](elems []T, n int) []T {
	if len(elems) <= n {
		return elems
	}
	return elems[:n]
}

// All returns true if all elements return true for given predicate
func All[T any](elems []T, fn func(T) bool) bool {
	for _, elem := range elems {
		if !fn(elem) {
			return false
		}
	}
	return true
}

// Any returns true if at least one element returns true for given predicate
func Any[T any](elems []T, fn func(T) bool) bool {
	for _, elem := range elems {
		if fn(elem) {
			return true
		}
	}
	return false
}

// Associate returns a map containing key-value pairs returned by the given
// function applied to the elements of the given slice
func Associate[T, V any, K comparable](elems []T, fn func(T) (K, V)) map[K]V {
	ret := make(map[K]V)
	for _, elem := range elems {
		k, v := fn(elem)
		ret[k] = v
	}
	return ret
}
