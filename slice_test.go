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

func TestAppendToGroup(t *testing.T) {
	grp := make(map[string][]int)
	AppendToGroup(grp, "a", 1)
	AppendToGroup(grp, "b", 2)
	AppendToGroup(grp, "a", 10)
	AppendToGroup(grp, "b", 20)
	AppendToGroup(grp, "a", 100)
	AppendToGroup(grp, "b", 200)

	want := map[string][]int{
		"a": {1, 10, 100},
		"b": {2, 20, 200},
	}
	if !reflect.DeepEqual(grp, want) {
		t.Errorf("AppendToGroup() = %v, want %v", grp, want)
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

func TestChunkedBy(t *testing.T) {
	input := []int{10, 20, 30, 40, 31, 31, 33, 34, 21, 22, 23, 24, 11, 12, 13, 14}
	output := ChunkedBy(input, func(prev, next int) bool {
		return prev < next
	})
	expected := [][]int{
		{10, 20, 30, 40},
		{31},
		{31, 33, 34},
		{21, 22, 23, 24},
		{11, 12, 13, 14},
	}
	if !reflect.DeepEqual(output, expected) {
		t.Errorf("ChunkedBy() = %v, want %v", output, expected)
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

func TestGroupByWithOriginalTypesForKeyAndValue(t *testing.T) {

	input := []string{"a", "abc", "ab", "def", "abcd"}
	want := map[int][]string{
		1: {"a"},
		2: {"ab"},
		3: {"abc", "def"},
		4: {"abcd"},
	}
	got := GroupBy(input, func(str string) (int, string) {
		return len(str), str
	})
	if !reflect.DeepEqual(got, want) {
		t.Errorf("GroupBy() = %v, want %v", got, want)
	}
}

type wrapped struct {
	value string
}

func (w wrapped) String() string {
	return fmt.Sprintf("Wrapped:'%s'", w.value)
}

func TestGroupByWithNewTypesForKeyAndValue(t *testing.T) {

	input := []string{"a", "abc", "ab", "def", "abcd"}
	want := map[float64][]*wrapped{
		10.0: {&wrapped{"a"}},
		20.0: {&wrapped{"ab"}},
		30.0: {&wrapped{"abc"}, &wrapped{"def"}},
		40.0: {&wrapped{"abcd"}},
	}
	got := GroupBy(input, func(str string) (float64, *wrapped) {
		return float64(len(str)) * 10.0, &wrapped{str}
	})
	for k, vs := range want {
		avs, ok := got[k]
		if !ok {
			t.Errorf("expected key '%v' not found", k)
			return
		}
		if len(vs) != len(avs) {
			t.Errorf("expected %d elements for key:'%v'. got %d",
				len(vs), k, len(avs))
			return
		}
		for i := 0; i < len(vs); i++ {
			av := avs[i]
			v := vs[i]
			if av.value != v.value {
				t.Errorf("expected value: %s, got %s", v, av)
			}
		}

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

func TestFoldItems(t *testing.T) {
	type args struct {
		m       map[int]int
		initial map[string]string
		fn      func(map[string]string, int, int) map[string]string
	}
	tests := []struct {
		name string
		args args
		want map[string]string
	}{
		{
			"test fold over map items",
			args{
				map[int]int{1: 10, 2: 20, 3: 30},
				make(map[string]string),
				func(acc map[string]string, k, v int) map[string]string {
					acc[fmt.Sprintf("entry_%d", k)] = fmt.Sprintf("%d->%d", k, v)
					return acc
				},
			},
			map[string]string{
				"entry_1": "1->10",
				"entry_2": "2->20",
				"entry_3": "3->30",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FoldItems(tt.args.m, tt.args.initial, tt.args.fn); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FoldItems() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTransformMap(t *testing.T) {
	type args struct {
		m  map[string][]int
		fn func(k string, v []int) (string, []int, bool)
	}
	tests := []struct {
		name string
		args args
		want map[string][]int
	}{
		{
			"filter entries",
			args{
				map[string][]int{
					"a": {1, 2, 3, 4},
					"b": {1, 2},
					"c": {1, 2, 3},
				},
				func(k string, v []int) (string, []int, bool) {
					if len(v) < 3 {
						return k, v, false
					}
					return k, v, true
				},
			},
			map[string][]int{
				"a": {1, 2, 3, 4},
				"c": {1, 2, 3},
			},
		},
		{
			"map entries",
			args{
				map[string][]int{
					"a": {1, 2, 3, 4},
					"b": {5, 6},
				},
				func(k string, v []int) (string, []int, bool) {
					newK := strings.ToUpper(k)
					newV := Map(v, func(i int) int { return i * 10 })
					return newK, newV, true
				},
			},
			map[string][]int{
				"A": {10, 20, 30, 40},
				"B": {50, 60},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := TransformMap(tt.args.m, tt.args.fn); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TransformMap() = %v, want %v", got, tt.want)
			}
		})
	}
}
