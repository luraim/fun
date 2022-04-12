package fun

import (
	"fmt"
	"reflect"
	"strings"
	"testing"
)

func TestAll(t *testing.T) {
	type args struct {
		elems []int
		fn    func(int) bool
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"positive case",
			args{
				[]int{1, 2, 3, 4, 5},
				func(i int) bool { return i < 10 },
			},
			true,
		},
		{"negative case",
			args{
				[]int{1, 2, 3, 4, 5},
				func(i int) bool { return i%2 == 0 },
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := All(tt.args.elems, tt.args.fn); got != tt.want {
				t.Errorf("All() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAny(t *testing.T) {
	type args struct {
		elems []int
		fn    func(int) bool
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"positive case",
			args{
				[]int{1, 2, 3, 4, 5, 6},
				func(i int) bool { return i%2 == 0 },
			},
			true,
		},
		{"negative case",
			args{
				[]int{1, 2, 3, 4, 5, 6},
				func(i int) bool { return i > 7 },
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Any(tt.args.elems, tt.args.fn); got != tt.want {
				t.Errorf("Any() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFilter(t *testing.T) {
	type args struct {
		elems []int
		fn    func(int) bool
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"test filtering",
			args{
				[]int{1, 2, 3, 4, 5, 6, 7, 8},
				func(i int) bool { return i%2 == 0 },
			},
			[]int{2, 4, 6, 8},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Filter(tt.args.elems, tt.args.fn); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Filter() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAssociate(t *testing.T) {
	type args struct {
		elems []int
		fn    func(int) (string, int)
	}
	tests := []struct {
		name string
		args args
		want map[string]int
	}{
		{"associate strings with ints",
			args{
				[]int{1, 2, 3, 4},
				func(i int) (string, int) {
					return fmt.Sprintf("M%d", i), i * 10
				},
			},
			map[string]int{"M1": 10, "M2": 20, "M3": 30, "M4": 40},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Associate(tt.args.elems, tt.args.fn); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Associate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestChunked(t *testing.T) {
	type args struct {
		s []int
		n int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{"exact multiple",
			args{
				[]int{1, 2, 3, 4, 5, 6, 7, 8, 9},
				3,
			},
			[][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}},
		},
		{"extra elements",
			args{
				[]int{1, 2, 3, 4, 5, 6, 7, 8, 9},
				2,
			},
			[][]int{{1, 2}, {3, 4}, {5, 6}, {7, 8}, {9}},
		},
		{"not enough elements",
			args{
				[]int{1, 2, 3},
				5,
			},
			[][]int{{1, 2, 3}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Chunked(tt.args.s, tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Chunked() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDistinct(t *testing.T) {
	type args struct {
		s []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"test distinct",
			args{
				[]int{1, 1, 2, 3, 3, 4, 4, 4, 4, 5, 5, 5},
			},
			[]int{1, 2, 3, 4, 5},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Distinct(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Distinct() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDistinctBy(t *testing.T) {
	type args struct {
		s  []string
		fn func(string) string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{{"test distinctBy",
		args{[]string{"a", "A", "b", "B", "c", "C"},
			func(s string) string { return strings.ToLower(s) },
		},
		[]string{"a", "b", "c"},
	}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DistinctBy(tt.args.s, tt.args.fn); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DistinctBy() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDrop(t *testing.T) {
	type args struct {
		s []int
		n int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"drop less than slice length",
			args{
				[]int{1, 2, 3, 4, 5, 6, 7, 8, 9},
				4,
			},
			[]int{5, 6, 7, 8, 9},
		},
		{"drop more than slice length",
			args{
				[]int{1, 2, 3, 4, 5, 6, 7, 8, 9},
				12,
			},
			[]int{},
		},
		{"drop slice length",
			args{
				[]int{1, 2, 3, 4, 5, 6, 7, 8, 9},
				9,
			},
			[]int{},
		},
		{"drop all but last",
			args{
				[]int{1, 2, 3, 4, 5, 6, 7, 8, 9},
				8,
			},
			[]int{9},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Drop(tt.args.s, tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Drop() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDropLast(t *testing.T) {
	type args struct {
		s []int
		n int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"drop less than slice length",
			args{
				[]int{1, 2, 3, 4, 5, 6, 7, 8, 9},
				4,
			},
			[]int{1, 2, 3, 4, 5},
		},
		{"drop more than slice length",
			args{
				[]int{1, 2, 3, 4, 5, 6, 7, 8, 9},
				12,
			},
			[]int{},
		},
		{"drop slice length",
			args{
				[]int{1, 2, 3, 4, 5, 6, 7, 8, 9},
				9,
			},
			[]int{},
		},
		{"drop all but last",
			args{
				[]int{1, 2, 3, 4, 5, 6, 7, 8, 9},
				8,
			},
			[]int{1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DropLast(tt.args.s, tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DropLast() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDropLastWhile(t *testing.T) {
	type args struct {
		s  []int
		fn func(int) bool
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			"test dropLastWhile",
			args{
				[]int{1, 2, 3, 4, 5, 6},
				func(i int) bool { return i > 3 },
			},
			[]int{1, 2, 3},
		},
		{
			"emtpy list",
			args{
				[]int{},
				func(i int) bool { return i > 3 },
			},
			[]int{},
		},
		{
			"drop all",
			args{
				[]int{1, 2, 3, 4, 5, 6},
				func(i int) bool { return i > 0 },
			},
			[]int{},
		},
		{
			"drop all but one",
			args{
				[]int{1, 2, 3, 4, 5, 6},
				func(i int) bool { return i > 1 },
			},
			[]int{1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DropLastWhile(tt.args.s, tt.args.fn); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DropLastWhile() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDropWhile(t *testing.T) {
	type args struct {
		s  []int
		fn func(int) bool
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"less than slice len",
			args{
				[]int{1, 2, 3, 4, 5, 6, 7, 8},
				func(i int) bool { return i < 5 },
			},
			[]int{5, 6, 7, 8},
		},
		{"empty slice",
			args{
				[]int{},
				func(i int) bool { return i < 5 },
			},
			[]int{},
		},

		{"drop all",
			args{
				[]int{1, 2, 3, 4, 5, 6, 7, 8},
				func(i int) bool { return i < 9 },
			},
			[]int{},
		},
		{"drop all but one",
			args{
				[]int{1, 2, 3, 4, 5, 6, 7, 8},
				func(i int) bool { return i < 8 },
			},
			[]int{8},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DropWhile(tt.args.s, tt.args.fn); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DropWhile() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDropVariants(t *testing.T) {
	letters := alphabet()

	got := Drop(letters, 23)
	expected := []rune{'x', 'y', 'z'}
	if !reflect.DeepEqual(got, expected) {
		t.Errorf("Drop() = %v want %v", got, expected)
	}

	got = DropLast(letters, 23)
	expected = []rune{'a', 'b', 'c'}
	if !reflect.DeepEqual(got, expected) {
		t.Errorf("DropLast() = %v want %v", got, expected)
	}

	got = DropWhile(letters, func(r rune) bool { return r < 'x' })
	expected = []rune{'x', 'y', 'z'}
	if !reflect.DeepEqual(got, expected) {
		t.Errorf("DropWhile() = %v want %v", got, expected)
	}

	got = DropLastWhile(letters, func(r rune) bool { return r > 'c' })
	expected = []rune{'a', 'b', 'c'}
	if !reflect.DeepEqual(got, expected) {
		t.Errorf("DropLastWhile() = %v want %v", got, expected)
	}
}

func TestFilterIndexed(t *testing.T) {
	type args struct {
		s  []int
		fn func(int, int) bool
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"filter indexed",
			args{
				[]int{0, 1, 2, 3, 4, 8, 6},
				func(index int, v int) bool { return index == v },
			},
			[]int{0, 1, 2, 3, 4, 6},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FilterIndexed(tt.args.s, tt.args.fn); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FilterIndexed() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFold(t *testing.T) {
	s := []int{1, 2, 3, 4, 5}
	fn := func(acc *int, v int) { *acc += v }
	sum := 0
	Fold(s, &sum, fn)
	expected := 15
	if sum != expected {
		t.Errorf("Fold() = %d want %d", sum, expected)
	}
}

func TestFoldIndexed(t *testing.T) {
	s := []int{1, 2, 3, 4, 5}
	fn := func(index int, acc *int, v int) { *acc += index * v }
	sum := 0
	FoldIndexed(s, &sum, fn)
	expected := 40
	if sum != expected {
		t.Errorf("FoldIndexed() = %d want %d", sum, expected)
	}
}

func TestFoldItems(t *testing.T) {
	m := map[int]int{1: 10, 2: 20, 3: 30}
	res := make(map[string]string)
	fn := func(acc map[string]string, k, v int) {
		acc[fmt.Sprintf("entry_%d", k)] = fmt.Sprintf("%d->%d", k, v)
	}
	FoldItems(m, res, fn)
	expected := map[string]string{
		"entry_1": "1->10",
		"entry_2": "2->20",
		"entry_3": "3->30",
	}
	if !reflect.DeepEqual(expected, res) {
		t.Errorf("FoldItems() = %v want %v", res, expected)
	}
}
func TestGroupBy(t *testing.T) {
	type args struct {
		s  []string
		fn func(string) int
	}
	tests := []struct {
		name string
		args args
		want map[int][]string
	}{
		{"group by",
			args{
				[]string{"a", "abc", "ab", "def", "abcd"},
				func(str string) int { return len(str) },
			},
			map[int][]string{
				1: []string{"a"},
				2: []string{"ab"},
				3: []string{"abc", "def"},
				4: []string{"abcd"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GroupBy(tt.args.s, tt.args.fn); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GroupBy() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMap(t *testing.T) {
	type args struct {
		elems []int
		fn    func(int) int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"test mapping",
			args{
				[]int{1, 2, 3, 4, 5},
				func(i int) int { return i * i },
			},
			[]int{1, 4, 9, 16, 25},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Map(tt.args.elems, tt.args.fn); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Map() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMapIndexed(t *testing.T) {
	type args struct {
		s  []int
		fn func(int, int) int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"map indexed",
			args{
				[]int{1, 2, 3, 4, 5},
				func(index, i int) int { return index * i },
			},
			[]int{0, 2, 6, 12, 20},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MapIndexed(tt.args.s, tt.args.fn); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MapIndexed() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPartition(t *testing.T) {
	type person struct {
		name string
		age  int
	}
	type args struct {
		s  []*person
		fn func(*person) bool
	}
	tom := &person{"Tom", 18}
	andy := &person{"Andy", 32}
	sarah := &person{"Sarah", 22}
	tests := []struct {
		name  string
		args  args
		want  []*person
		want1 []*person
	}{
		{"partition",
			args{
				[]*person{tom, andy, sarah},
				func(p *person) bool { return p.age < 30 },
			},
			[]*person{tom, sarah},
			[]*person{andy},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := Partition(tt.args.s, tt.args.fn)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Partition() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("Partition() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestReverse(t *testing.T) {
	s := []int{1, 2, 3, 4, 5, 6, 7}
	Reverse(s)
	want := []int{7, 6, 5, 4, 3, 2, 1}
	if !reflect.DeepEqual(s, want) {
		t.Errorf("Reversed() = %v, want %v", s, want)
	}
}

func TestReversed(t *testing.T) {
	type args struct {
		s []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"reversed",
			args{
				[]int{1, 2, 3, 4, 5},
			},
			[]int{5, 4, 3, 2, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Reversed(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Reversed() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTake(t *testing.T) {
	type args struct {
		elems []int
		n     int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"take less than slice length",
			args{
				[]int{1, 2, 3, 4, 5, 6, 7, 8},
				4,
			},
			[]int{1, 2, 3, 4},
		},
		{"take more than slice length",
			args{
				[]int{1, 2, 3, 4, 5, 6, 7, 8},
				10,
			},
			[]int{1, 2, 3, 4, 5, 6, 7, 8},
		},
		{"take slice length",
			args{
				[]int{1, 2, 3, 4, 5, 6, 7, 8},
				8,
			},
			[]int{1, 2, 3, 4, 5, 6, 7, 8},
		},
		{"empty list",
			args{
				[]int{},
				8,
			},
			[]int{},
		},
		{"take 0 elements",
			args{
				[]int{1, 2, 3, 4, 5, 6, 7, 8},
				0,
			},
			[]int{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Take(tt.args.elems, tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Take() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTakeLast(t *testing.T) {
	type args struct {
		s []int
		n int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"less than slice len",
			args{
				[]int{1, 2, 3, 4, 5, 6, 7, 8},
				2,
			},
			[]int{7, 8},
		},
		{"more than slice len",
			args{
				[]int{1, 2, 3, 4, 5, 6, 7, 8},
				9,
			},
			[]int{1, 2, 3, 4, 5, 6, 7, 8},
		},
		{"exactly slice len",
			args{
				[]int{1, 2, 3, 4, 5, 6, 7, 8},
				8,
			},
			[]int{1, 2, 3, 4, 5, 6, 7, 8},
		},
		{"empty list",
			args{
				[]int{},
				8,
			},
			[]int{},
		},
		{"take 0 elements",
			args{
				[]int{1, 2, 3, 4, 5, 6, 7, 8},
				0,
			},
			[]int{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := TakeLast(tt.args.s, tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TakeLast() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTake2(t *testing.T) {
	got := Take(alphabet(), 2)
	expected := []rune{'a', 'b'}
	if !reflect.DeepEqual(got, expected) {
		t.Errorf("TakeLast() = %v want %v", got, expected)
	}
}

func TestTakeLast2(t *testing.T) {
	got := TakeLast(alphabet(), 2)
	expected := []rune{'y', 'z'}
	if !reflect.DeepEqual(got, expected) {
		t.Errorf("TakeLast() = %v want %v", got, expected)
	}
}

func TestTakeLastWhile(t *testing.T) {
	got := TakeLastWhile(alphabet(), func(s rune) bool { return s > 'w' })
	expected := []rune{'x', 'y', 'z'}
	if !reflect.DeepEqual(got, expected) {
		t.Errorf("TakeLastWhile() = %v want %v", got, expected)
	}
}

func TestTakeWhile(t *testing.T) {
	got := TakeWhile(alphabet(), func(s rune) bool { return s < 'f' })
	expected := []rune{'a', 'b', 'c', 'd', 'e'}
	if !reflect.DeepEqual(got, expected) {
		t.Errorf("TakeWhile() = %v want %v", got, expected)
	}
}

func TestWindowed(t *testing.T) {
	type args struct {
		s    []int
		size int
		step int
	}
	input := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{"size = 5, step = 1",
			args{input, 5, 1},
			[][]int{
				{1, 2, 3, 4, 5},
				{2, 3, 4, 5, 6},
				{3, 4, 5, 6, 7},
				{4, 5, 6, 7, 8},
				{5, 6, 7, 8, 9},
				{6, 7, 8, 9, 10},
				{7, 8, 9, 10},
				{8, 9, 10},
				{9, 10},
				{10},
			},
		},
		{"size = 5, step = 3",
			args{input, 5, 3},
			[][]int{
				{1, 2, 3, 4, 5},
				{4, 5, 6, 7, 8},
				{7, 8, 9, 10},
				{10},
			},
		},
		{"size = 3, step = 4",
			args{input, 3, 4},
			[][]int{
				{1, 2, 3},
				{5, 6, 7},
				{9, 10},
			},
		},

		{"slice smaller than size",
			args{[]int{1, 2, 3}, 4, 1},
			[][]int{
				{1, 2, 3},
				{2, 3},
				{3},
			},
		},
		{"slice smaller than size and step",
			args{[]int{1, 2, 3}, 4, 4},
			[][]int{
				{1, 2, 3},
			},
		},
		{"slice larger than size and smaller than step",
			args{[]int{1, 2, 3}, 2, 4},
			[][]int{
				{1, 2},
			},
		},
		{"empty slice",
			args{[]int{}, 4, 4},
			[][]int{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Windowed(tt.args.s, tt.args.size, tt.args.step); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Windowed() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestZip(t *testing.T) {
	s1 := []string{"a", "b", "c", "d"}
	s2 := []int{1, 2, 3}
	got := Zip(s1, s2)
	want := []*Pair[string, int]{
		{"a", 1},
		{"b", 2},
		{"c", 3},
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Zip() = %v, want %v", got, want)
	}
}

func TestUnZip(t *testing.T) {
	ps := []*Pair[string, int]{
		{"a", 1},
		{"b", 2},
		{"c", 3},
	}
	want1 := []string{"a", "b", "c"}
	want2 := []int{1, 2, 3}
	got1, got2 := Unzip(ps)
	if !reflect.DeepEqual(got1, want1) {
		t.Errorf("Zip() first list = %v, want %v", got1, want1)
	}
	if !reflect.DeepEqual(got2, want2) {
		t.Errorf("Zip() first list = %v, want %v", got2, want2)
	}
}

func TestFilterMap(t *testing.T) {
	s := []int{1, 2, 3, 4, 5}
	got := FilterMap(s,
		func(i int) (int, bool) {
			if i%2 != 0 {
				return i, false // drop odd numbers
			}
			return i * i, true // square even numbers
		},
	)
	want := []int{4, 16}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("FilterMap() = %v, want %v", got, want)
	}
}

func alphabet() []rune {
	ret := make([]rune, 0)
	for r := 'a'; r <= 'z'; r++ {
		ret = append(ret, r)
	}
	return ret
}

func TestGetOrInsert(t *testing.T) {

	m := map[int]int{1: 10, 2: 20}
	fn := func(k int) int { return k * 10 }

	// case where key is present
	expected := 20
	got := GetOrInsert(m, 2, fn)
	if got != expected {
		t.Errorf("GetOrInsert() = %d expected = %d", got, expected)
	}

	// case where key is not present, but populated by invoking the function
	expected = 30
	got = GetOrInsert(m, 3, fn)
	if got != expected {
		t.Errorf("GetOrInsert() = %d expected = %d", got, expected)
	}

	// check that the new value was stored in the map as well
	expected = 30
	got, ok := m[3]
	if !ok {
		t.Errorf("GetOrInsert did not insert new value in map!!")
	}
	if got != expected {
		t.Errorf("value in map = %d expected = %d", got, expected)
	}

}
