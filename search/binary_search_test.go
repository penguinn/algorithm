package search

import (
	"testing"
)

func TestBinarySearch(t *testing.T) {
	type args struct {
		data  []int
		value int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test1",
			args: args{
				data:  []int{8, 11, 19, 23, 27, 33, 45, 55, 67, 98},
				value: 23,
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BinarySearch(tt.args.data, tt.args.value); got != tt.want {
				t.Errorf("BinarySearch() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBinarySearchGE(t *testing.T) {
	type args struct {
		data  []int
		value int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test1",
			args: args{
				data:  []int{1, 4, 7, 7, 8, 9, 10},
				value: 5,
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BinarySearchGE(tt.args.data, tt.args.value); got != tt.want {
				t.Errorf("BinarySearchGE() = %v, want %v", got, tt.want)
			}
		})
	}
}
