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
