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

// Chunked splits the slice into a slice of slices, each not exceeding given size
// The last slice might have fewer elements than the given size
func Chunked[T any](s []T, chunkSize int) [][]T {
	ret := make([][]T, 0)
	sz := len(s)
	var sub []T
	for i := 0; i < sz; i++ {
		if i%chunkSize == 0 {
			if len(sub) > 0 {
				ret = append(ret, sub)
			}
			sub = make([]T, 0)
		}
		sub = append(sub, s[i])
	}
	if len(sub) > 0 {
		ret = append(ret, sub)
	}
	return ret
}

// Distinct returns a slice containing only distinct elements from the given slice
// Elements will retain their original order.
func Distinct[T comparable](s []T) []T {
	m := make(map[T]bool)
	ret := make([]T, 0)
	for _, e := range s {
		_, ok := m[e]
		if ok {
			continue
		}
		m[e] = true
		ret = append(ret, e)
	}
	return ret
}

// DistinctBy returns a slice containing only distinct elements from the
// given slice as distinguished by the given selector function
// Elements will retain their original order.
func DistinctBy[T any, K comparable](s []T, fn func(T) K) []T {
	m := make(map[K]bool)
	ret := make([]T, 0)
	for _, e := range s {
		k := fn(e)
		_, ok := m[k]
		if ok {
			continue
		}
		m[k] = true
		ret = append(ret, e)
	}
	return ret
}

// Drop returns a slice containing all elements except the first n
func Drop[T any](s []T, n int) []T {
	if n >= len(s) {
		return make([]T, 0)
	}
	return s[n:]
}

// DropLast returns a slice containing all elements except the last n
func DropLast[T any](s []T, n int) []T {
	if n >= len(s) {
		return make([]T, 0)
	}
	return s[:len(s)-n]
}

// DropLastWhile returns a slice containing all elements except the last elements
// that satisfy the given predicate
func DropLastWhile[T any](s []T, fn func(T) bool) []T {
	if len(s) == 0 {
		return s
	}
	i := len(s) - 1
	for ; i >= 0; i-- {
		if !fn(s[i]) {
			break
		}
	}
	return s[:i+1]
}
