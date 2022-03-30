# fun
[![GoDoc](https://godoc.org/github.com/luraim/fun?status.svg)](https://godoc.org/github.com/luraim/fun)

### Simple generic utility functions to reduce golang boilerplate
#### Inspired by Kotlin and Rust collection functions

## List of functions
 - [All](#all)
 - [Any](#any)
 - [Associate](#associate)
 - [Chunked](#chunked)
 - [Distinct](#distinct)
 - [DistinctBy](#distinctby)
 - [Drop](#drop)
 - [DropLast](#droplast)
 - [DropWhile](#dropwhile)
 - [DropLastWhile](#droplastwhile)
 - [Filter](#filter)
 - [FilterIndexed](#filterindexed)
 - [FilterMap](#filtermap)
 - [Fold](#fold)
 - [FoldIndexed](#foldindexed)
 - [GroupBy](#groupby)
 - [Map](#map)
 - [MapIndexed](#mapindexed)
 - [Partition](#partition)
 - [Reduce](#reduce)
 - [ReduceIndexed](#reduceindexed)
 - [Reverse](#reverse)
 - [Reversed](#reversed)
 - [Take](#take)
 - [TakeLast](#takelast)
 - [TakeWhile](#takewhile)
 - [TakeLastWhile](#takelastwhile)
 - [Unzip](#unzip)
 - [Windowed](#windowed)
 - [Zip](#zip)

### All
- Returns true if all elements return true for given predicate
```go
All([]int{1, 2, 3, 4, 5}, func(i int)bool {return i < 7})
// true

All([]int{1, 2, 3, 4, 5}, func(i int)bool {return i % 2 == 0})
// false

```

### Any
- Returns true if at least one element returns true for given predicate
```go
Any([]int{1, 2, 3}, func(i int)bool {return i%2==0})
// true

Any([]int{1, 2, 3}, func(i int)bool {return i > 7})
// false
```

### Associate
- Returns a map containing key-value pairs returned by the given function applied to the elements of the given slice
```go
Associate([]int{1, 2, 3, 4}, func(i int) (string, int) {
    return fmt.Sprintf("M%d", i), i * 10
})
// {"M1": 10, "M2": 20, "M3": 30, "M4": 40}
```

### Chunked
- Splits the slice into a slice of slices, each not exceeding given chunk size
- The last slice might have fewer elements than the given chunk size
```go
Chunked([]int{1, 2, 3, 4, 5, 6, 7, 8, 9}, 2)
// [[1, 2], [3, 4], [5, 6], [7, 8], [9]]
```

### Distinct
- Returns a slice containing only distinct elements from the given slice
```go
Distinct([]int{1, 1, 2, 3, 3, 4, 4, 4, 4, 5, 5, 5})
// [1, 2, 3, 4, 5]
```

### DistinctBy
- Returns a slice containing only distinct elements from the given slice as distinguished by the given selector function
```go
DistinctBy([]string{"a", "A", "b", "B", "c", "C"},func(s string) string {
	return strings.ToLower(s)
})
// ["a", "b", "c"]
```

### Drop
- Returns a slice containing all elements except the first n.
```go
// letters = ['a'..'z']
Drop(letters, 23)
// ['x', 'y', 'z']
```

### DropLast
- Returns a slice containing all elements except the last n.
```go
// letters = ['a'..'z']
DropLast(letters, 23)
// ['a', 'b', 'c']
```

### DropWhile
- Returns a slice containing all elements except the first elements that satisfy the given predicate.
```go
// letters = ['a'..'z']
DropWhile(letters, func(r rune) bool { return r < 'x' })
// ['x', 'y', 'z']
```

### DropLastWhile
- Returns a slice containing all elements except the last elements that satisfy the given predicate.
```go
// letters = ['a'..'z']
DropLastWhile(letters, func(r rune) bool { return r > 'c' })
// ['a', 'b', 'c']
```

### Filter
- Returns the slice obtained after retaining only those elements in the given slice for which the given function returns true
```go
Filter([]int{1, 2, 3, 4, 5, 6, 7, 8}, func(i int)bool {return i%2==0})
// [2, 4, 6, 8]
```

### FilterIndexed
- Returns the slice obtained after retaining only those elements in the given slice for which the given function returns true
- Predicate function receives the value as well as its index in the slice.
```go
FilterIndexed([]int{0, 1, 2, 3, 4, 8, 6}, func(index int, v int) bool {
	return index == v
})
// [0, 1, 2, 3, 4, 6]
```

### FilterMap
- FilterMap returns the slice obtained after both filtering and mapping using the given function.
- The function should return two values - the result of the mapping operation and whether the element should be included or dropped.
- This is faster than doing separate filter and map operations, since it avoids extra allocations and slice traversals.
- Inspired by std::iter::filter_map in Rust
```go
FilterMap([]int{1, 2, 3, 4, 5},
    func(i int) (int, bool) {
        if i%2 != 0 {
            return i, false // drop odd numbers
        }
        return i * i, true // square even numbers
    })
// [4, 16]
```

### Fold
- Accumulates values starting with given initial value and applying given function to current accumulator and each element.
```go
Fold([]int{1, 2, 3, 4, 5}, func(acc, v int) int { return acc + v })
// 15
```

### FoldIndexed
- Accumulates values starting with given initial value and applying given function to current accumulator and each element.
- Function also receives index of current element.
```go
FoldIndexed([]int{1, 2, 3, 4, 5}, func(index, acc, v int) int {
	return acc + index*v
})
// 40
```

### GroupBy
- Returns a map where each key maps to slices of elements all having the same key as returned by the given function
```go
GroupBy([]string{"a", "abc", "ab", "def", "abcd"}, func(s string) int {
	return len(s)
})
// {1: ["a"], 2: ["ab"], 3: ["abc", "def"], 4: ["abcd"]},
```

### Map
- Returns the slice obtained after applying the given function over every element in the given slice
```go
Map([]int{1, 2, 3, 4, 5}, func(i int) int { return i * i })
// [1, 4, 9, 16, 25]
```

### MapIndexed
- Returns the slice obtained after applying the given function over every element in the given slice
- The function also receives the index of each element in the slice.
```go
MapIndexed([]int{1, 2, 3, 4, 5}, func(index, i int) int { return index * i })
// [0, 2, 6, 12, 20]
```

### Partition
- Returns two slices where the first slice contains elements for which the predicate returned true and the second slice contains elements for which it returned false.
```go
type person struct {
    name string
    age  int
}

tom := &person{"Tom", 18}
andy := &person{"Andy", 32}
sarah := &person{"Sarah", 22}

Partition([]*person{tom, andy, sarah}, func(p *person) bool { return p.age < 30 })
// [tom, sarah], [andy]
```

### Reduce
- Accumulates the values starting with the first element and applying the operation from left to right to the current accumulator value and each element.
- The input slice must have at least one element.
```go
Reduce([]int{1, 2, 3, 4, 5}, func(acc, v int) int { return acc + v })
// 15
```

### ReduceIndexed
- Accumulates the values starting with the first element and applying the operation from left to right to the current accumulator value and each element.
- The input slice must have at least one element.
- The function also receives the index of each element.
```go
ReduceIndexed([]string{"a", "b", "c", "d"}, func(index int, acc, v string) string {
    return fmt.Sprintf("%s%s%d", acc, v, index)
})
// "ab1c2d3"
```

### Reverse
- Reverses the elements of the list in place.
```go
// s = [1, 2, 3, 4, 5, 6, 7]
Reverse(s)
// s = [7, 6, 5, 4, 3, 2, 1]
```

### Reversed
- Returns a new list with the elements in reverse order.
```go
// s = [1, 2, 3, 4, 5, 6, 7]
r := Reversed(s)
// r = [7, 6, 5, 4, 3, 2, 1]
// s = [1, 2, 3, 4, 5, 6, 7]
```

### Take
- Returns the slice obtained after taking the first n elements from the given slice.
```go
// letters = ['a'..'z']
Take(letters, 2)
// ['a', 'b']
```

### TakeLast
- Returns the slice obtained after taking the last n elements from the given slice.
```go
// letters = ['a'..'z']
TakeLast(letters, 2)
// ['y', 'z']
```

### TakeWhile
- Returns a slice containing the first elements satisfying the given predicate
```go
// letters = ['a'..'z']
TakeWhile(letters,  func(s rune) bool { return s < 'f' })
// ['a', 'b', 'c', 'd', 'e']
```

### TakeLastWhile
- Returns a slice containing the last elements satisfying the given predicate
```go
// letters = ['a'..'z']
TakeLastWhile(letters, func(s rune) bool { return s > 'w' })
// ['x', 'y', 'z']
```

### Unzip
- Returns two slices, where:
- the first slice is built from the first values of each pair from the input slice
- the second slice is built from the second values of each pair
```go
Unzip([]*Pair[string, int]{{"a", 1}, {"b", 2}, {"c", 3}})
// ["a", "b", "c"], [1, 2, 3]
```

### Windowed
- Returns a slice of sliding windows, each of the given size, and with the given step
- Several last slices may have fewer elements than the given size
```go
Windowed([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 5, 1)
// [
//     [1, 2, 3, 4, 5],
//     [2, 3, 4, 5, 6],
//     [3, 4, 5, 6, 7],
//     [4, 5, 6, 7, 8],
//     [5, 6, 7, 8, 9],
//     [6, 7, 8, 9, 10],
//     [7, 8, 9, 10],
//     [8, 9, 10],
//     [9, 10],
//     [10]
// ]

Windowed([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 5, 3)
// [
//     [1, 2, 3, 4, 5],
//     [4, 5, 6, 7, 8],
//     [7, 8, 9, 10],
//     [10]
// ]
```

### Zip
- Returns a slice of pairs from the elements of both slices with the same index
- The returned slice has the length of the shortest input slice
```go
Zip([]string{"a", "b", "c", "d"}, []int{1, 2, 3})
// []*Pair[string, int]{{"a", 1}, {"b", 2}, {"c", 3}}
```
