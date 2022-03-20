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
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Take(tt.args.elems, tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Take() = %v, want %v", got, tt.want)
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
	type args struct {
		s       []int
		initial int
		fn      func(int, int) int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"summation by fold",
			args{
				[]int{1, 2, 3, 4, 5},
				0,
				func(acc, v int) int { return acc + v },
			},
			15,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Fold(tt.args.s, tt.args.initial, tt.args.fn); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Fold() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFoldIndexed(t *testing.T) {
	type args struct {
		s       []int
		initial int
		fn      func(int, int, int) int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"fold indexed",
			args{
				[]int{1, 2, 3, 4, 5},
				0,
				func(index, acc, v int) int { return acc + index*v },
			},
			40,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FoldIndexed(tt.args.s, tt.args.initial, tt.args.fn); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FoldIndexed() = %v, want %v", got, tt.want)
			}
		})
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

func TestReduce(t *testing.T) {
	type args struct {
		s  []int
		fn func(int, int) int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"reduce",
			args{
				[]int{1, 2, 3, 4, 5},
				func(acc, v int) int { return acc + v },
			},
			15,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Reduce(tt.args.s, tt.args.fn); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Reduce() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestReduceIndexed(t *testing.T) {
	type args struct {
		s  []string
		fn func(int, string, string) string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"reduce indexed",
			args{
				[]string{"a", "b", "c", "d"},
				func(index int, acc, v string) string {
					return fmt.Sprintf("%s%s%d", acc, v, index)
				},
			},

			"ab1c2d3",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ReduceIndexed(tt.args.s, tt.args.fn); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ReduceIndexed() = %v, want %v", got, tt.want)
			}
		})
	}
}
