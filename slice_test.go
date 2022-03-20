package fun

import (
	"reflect"
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
