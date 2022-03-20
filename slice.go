package fun

// All returns true if all elements return true for given predicate
func All[T any](s []T, fn func(T) bool) bool {
	for _, e := range s {
		if !fn(e) {
			return false
		}
	}
	return true
}

// Any returns true if at least one element returns true for given predicate
func Any[T any](s []T, fn func(T) bool) bool {
	for _, e := range s {
		if fn(e) {
			return true
		}
	}
	return false
}

// Associate returns a map containing key-value pairs returned by the given
// function applied to the elements of the given slice
func Associate[T, V any, K comparable](s []T, fn func(T) (K, V)) map[K]V {
	ret := make(map[K]V)
	for _, e := range s {
		k, v := fn(e)
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

// DropWhile returns a slice containing all elements except the first elements
// that satisfy the given predicate
func DropWhile[T any](s []T, fn func(T) bool) []T {
	if len(s) == 0 {
		return s
	}
	i := 0
	for ; i < len(s); i++ {
		if !fn(s[i]) {
			break
		}
	}
	return s[i:]
}

// Filter returns the slice obtained after retaining only those elements
// in the given slice for which the given function returns true
func Filter[T any](s []T, fn func(T) bool) []T {
	ret := make([]T, 0)
	for _, e := range s {
		if fn(e) {
			ret = append(ret, e)
		}
	}
	return ret
}

// FilterIndexed returns the slice obtained after retaining only those elements
// in the given slice for which the given function returns true. Predicate
// receives the value as well as its index in the slice.
func FilterIndexed[T any](s []T, fn func(int, T) bool) []T {
	ret := make([]T, 0)
	for i, e := range s {
		if fn(i, e) {
			ret = append(ret, e)
		}
	}
	return ret
}

// Fold accumulates values starting with given initial value and applying
// given function to current accumulator and each element.
func Fold[T, R any](s []T, initial R, fn func(R, T) R) R {
	acc := initial
	for _, e := range s {
		acc = fn(acc, e)
	}
	return acc
}

// FoldIndexed accumulates values starting with given initial value and applying
// given function to current accumulator and each element. Function also
// receives index of current element.
func FoldIndexed[T, R any](s []T, initial R, fn func(int, R, T) R) R {
	acc := initial
	for i, e := range s {
		acc = fn(i, acc, e)
	}
	return acc
}

// GroupBy returns a map where each key maps to slices of elements all having
// the same key as returned by given function
func GroupBy[T any, K comparable](s []T, fn func(T) K) map[K][]T {
	ret := make(map[K][]T)
	for _, e := range s {
		k := fn(e)
		group, ok := ret[k]
		if !ok {
			group = make([]T, 0)
		}
		group = append(group, e)
		ret[k] = group
	}
	return ret
}

// Map returns the slice obtained after applying the given function over every
// element in the given slice
func Map[T1, T2 any](s []T1, fn func(T1) T2) []T2 {
	ret := make([]T2, 0)
	for _, e := range s {
		ret = append(ret, fn(e))
	}
	return ret
}

// MapIndexed returns the slice obtained after applying the given function over every
// element in the given slice. The function also receives the index of each
// element in the slice.
func MapIndexed[T1, T2 any](s []T1, fn func(int, T1) T2) []T2 {
	ret := make([]T2, 0)
	for i, e := range s {
		ret = append(ret, fn(i, e))
	}
	return ret
}

// Partition returns two slices where the first slice contains elements for
// which the predicate returned true and the second slice contains elements for
// which it returned false.
func Partition[T any](s []T, fn func(T) bool) ([]T, []T) {
	trueList := make([]T, 0)
	falseList := make([]T, 0)
	for _, e := range s {
		if fn(e) {
			trueList = append(trueList, e)
		} else {
			falseList = append(falseList, e)
		}
	}
	return trueList, falseList
}

// Reduce accumulates the values starting with the first element and applying the
// operation from left to right to the current accumulator value and each element
// The input slice must have at least one element.
func Reduce[T any](s []T, fn func(T, T) T) T {
	if len(s) == 1 {
		return s[0]
	}
	return Fold(s[1:], s[0], fn)
}

// ReduceIndexed accumulates the values starting with the first element and applying the
// operation from left to right to the current accumulator value and each element
// The input slice must have at least one element. The function also receives
// the index of the element.
func ReduceIndexed[T any](s []T, fn func(int, T, T) T) T {
	if len(s) == 1 {
		return s[0]
	}
	acc := s[0]
	for i, e := range s[1:] {
		acc = fn(i+1, acc, e)
	}
	return acc
}

// Reverse reverses the elements of the list in place
func Reverse[T any](s []T) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

// Reversed returns a new list with the elements in reverse order
func Reversed[T any](s []T) []T {
	ret := make([]T, 0)
	for i := len(s) - 1; i >= 0; i-- {
		ret = append(ret, s[i])
	}
	return ret
}

// Take returns the slice obtained after taking the first n elements from the
// given slice.
// If n is greater than the length of the slice, returns the entire slice
func Take[T any](s []T, n int) []T {
	if len(s) <= n {
		return s
	}
	return s[:n]
}

// TakeLast returns the slice obtained after taking the last n elements from the
// given slice.
func TakeLast[T any](s []T, n int) []T {
	if len(s) <= n {
		return s
	}
	return s[len(s)-n:]
}
